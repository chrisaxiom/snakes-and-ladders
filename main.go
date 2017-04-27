package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chrisaxiom/snakes-and-ladders/v1multi"
)

func main() {

	now := time.Now()

	if len(os.Args) < 2 {
		log.Fatal("must provide path")
	}

	filePath := os.Args[1]

	var timeSec float64

	f := func(file string) error {
		snakes, ladders, err := loadData(file)
		if err != nil {
			return err
		}
		now := time.Now()

		//fa := v1single.FindMinimumRolls(snakes, ladders) // finished in 0.011812 seconds, 0.008289 algo time
		fa := v1multi.FindMinimumRolls(snakes, ladders) // finished in 0.009739 seconds, 0.005233 algo time
		timeSec += time.Since(now).Seconds()
		fmt.Printf("firstAttempt: %s == %d\n", file, fa)
		return nil
	}

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		// don't do directories, even if they name a directory with an extension
		if info.IsDir() {
			return nil
		}

		// only look for txt files
		if filepath.Ext(path) == ".json" {
			//if strings.HasSuffix(path, "snakes_easy_test_1.json") {
			if err := f(path); err != nil {
				return err
			}
			//}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("finished in %f seconds, %f algo time\n", time.Since(now).Seconds(), timeSec)
}
