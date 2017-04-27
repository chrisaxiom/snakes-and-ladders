package v1multi

import (
	"math"
	"sync"
)

// my first (perhaps niave) attempt

// problem given was always 100 tile board
// but I tested to about 100000

// for now, a brute force approach
// a goroutine for each roll (yikes)
// we can do 2 things to trim the solution space
// 1. report 1st found path, stop computation on all gorutines that already have a longer path
// 2. keep history of which nodes we've visited in a path so that we know to prevent cycles
// 3. perhaps expanding on that, we need to keep track of how many rolls to a particular
// square, so that if we know we can reach that square in less rolls than your current # rolls, then abandon your search

// yes, I perhaps should use a channel,
type rollMap struct {
	sync.Mutex
	// 1-100 indexed
	rolls map[int32]int
}

var rollmap = rollMap{
	rolls: make(map[int32]int),
}

func resetMap() {
	// initialize roll map state
	for i := 0; i < int(maxPosition); i++ {
		rollmap.rolls[int32(i+1)] = math.MaxInt32
	}
}

func getAnswer() int {
	return rollmap.rolls[maxPosition]
}

// returns true if you should continue
func (rm *rollMap) update(position int32, rolls int) bool {
	rm.Lock()
	defer rm.Unlock()
	if val, exists := rm.rolls[position]; exists {
		// continue if your rolls are less than what has been seen before
		if val > rolls {
			rm.rolls[position] = rolls
			return true
		}
		return false
	}
	rm.rolls[position] = rolls
	return true
}

const (
	// scales to about 10000
	// memory becomes an issue after that
	maxPosition int32 = 100
)

//func pathFinder(workchan chan<- func(), curPos int32, numRolls int, snakes, ladders map[int32]int32) {
func pathFinder(wg *sync.WaitGroup, curPos int32, numRolls int, snakes, ladders map[int32]int32) {
	//fmt.Printf("found myself at %d after rolling %d times\n", curPos, numRolls)
	// base cases
	// check for snakes
	if newv, exists := snakes[curPos]; exists {
		// call ourselves with new position (no new rolls are needed)
		pathFinder(wg, newv, numRolls, snakes, ladders)
		return
	}
	// check for ladders
	if newv, exists := ladders[curPos]; exists {
		//fmt.Printf("taking ladder %d to %d\n", curPos, newv)
		// call ourselves with new position (no new rolls are needed)
		pathFinder(wg, newv, numRolls, snakes, ladders)
		return
	}
	// check to see if we are already are past the best
	// check to see if our current spot has already seen a better solution
	if !rollmap.update(curPos, numRolls) {
		return
	}
	// check to see if we have a solution
	for i := 1; i < 7; i++ {
		newvalue := curPos + int32(i)
		if newvalue <= maxPosition {
			wg.Add(1)
			//fmt.Printf("rolling from %d to %d\n", curPos, newvalue)
			go func(pos int32, rolls int) {
				defer wg.Done()
				pathFinder(wg, pos, rolls, snakes, ladders)
			}(newvalue, numRolls+1)
		}
	}
}

// FindMinimumRolls returns the minimum number of rolls needed
// to win the game
// returns int max if no solution
func FindMinimumRolls(snakes, ladders map[int32]int32) int {

	resetMap()

	// TODO: tune the number of go-routines ?
	// current on my mac it gets to almost 300% CPU
	// but not sure how much of that is busy waiting
	// TODO: do singlethreaded to see how much faster/slower it is

	var wg sync.WaitGroup
	// send goroutines to the workChan
	for i := 1; i < 7; i++ {
		wg.Add(1)
		go func(val int32) {
			defer wg.Done()
			pathFinder(&wg, val, 1, snakes, ladders)
		}(int32(i))
	}

	// wait until these are all finished
	wg.Wait()
	// return the value at the map position 100
	return getAnswer()
}
