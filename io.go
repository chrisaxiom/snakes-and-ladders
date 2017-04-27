package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type pair []int32

type snakes []pair
type ladders []pair

// Data is a struct for mapping the json data
type Data struct {
	Ladders ladders `json:"ladders"`
	Snakes  snakes  `json:"snakes"`
}

// loadData outputs the snakes and ladders and an error if there is one
func loadData(f string) (map[int32]int32, map[int32]int32, error) {
	snakes := make(map[int32]int32)
	ladders := make(map[int32]int32)
	file, err := os.Open(f)
	if err != nil {
		return nil, nil, err
	}
	dec := json.NewDecoder(file)
	var d Data
	err = dec.Decode(&d)
	if err != nil {
		return nil, nil, err
	}
	// convert to maps for faster lookups
	for _, snake := range d.Snakes {
		// sanity check to make sure there are not two ladders from
		// a single spot
		if len(snake) < 2 {
			return nil, nil, fmt.Errorf("Invlalid snake slice length: %d", len(snake))
		}
		start := snake[0]
		if _, exists := snakes[start]; exists {
			return nil, nil, fmt.Errorf("%d shouldn't exist in the snake map already", start)
		}
		snakes[start] = snake[1]
	}

	for _, ladder := range d.Ladders {
		// sanity check to make sure there are not two ladders from
		// a single spot
		if len(ladder) < 2 {
			return nil, nil, fmt.Errorf("Invlalid ladder slice length: %d", len(ladder))
		}
		start := ladder[0]
		if _, exists := ladders[start]; exists {
			return nil, nil, fmt.Errorf("%d shouldn't exist in the ladder map already", start)
		}
		ladders[start] = ladder[1]
	}
	return snakes, ladders, nil
}
