package v1single

import "math"

// my first (perhaps niave) attempt is to simply build a graph of the problem,
// then do a "shortest path" search from beginning to end node

// always 100 tile board

// brute force approach
// a goroutine for each roll (exponential ^6)
// we can do 2 things to trim the solution space
// 1. report 1st found path, stop computation on all gorutines that already have a longer path
// 2. keep history of which nodes we've visited in a path so that we know to prevent cycles
// 3. perhaps expanding on that, we need to keep track of how many rolls to a particular
// square, so that if we know we can reach that square in less rolls than your current # rolls, then abandon your search

// yes, I perhaps should use a channel,
type rollMap struct {
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
	// memory, cpu becomes an issue after that
	maxPosition int32 = 100
)

func pathFinder(curPos int32, numRolls int, snakes, ladders map[int32]int32) {
	// base cases
	// check for snakes
	if newv, exists := snakes[curPos]; exists {
		// call ourselves with new position (no new rolls are needed)
		pathFinder(newv, numRolls, snakes, ladders)
		return
	}
	// check for ladders
	if newv, exists := ladders[curPos]; exists {
		// call ourselves with new position (no new rolls are needed)
		pathFinder(newv, numRolls, snakes, ladders)
		return
	}
	// check to see if we are already are past the best
	if !rollmap.update(curPos, numRolls) {
		return
	}
	// check to see if we have a solution
	for i := 1; i < 7; i++ {
		newvalue := curPos + int32(i)
		if newvalue <= maxPosition {
			pathFinder(newvalue, numRolls+1, snakes, ladders)
		}
	}
}

// FindMinimumRolls returns the minimum number of rolls needed
// to win the game
// returns int max if no solution
func FindMinimumRolls(snakes, ladders map[int32]int32) int {

	resetMap()

	for i := 1; i < 7; i++ {
		pathFinder(int32(i), 1, snakes, ladders)
	}

	// return the value at the map position 100
	return getAnswer()
}
