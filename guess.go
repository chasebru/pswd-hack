package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TryPassword(guess int32, hash string) bool {

	fguess, err := FormatGuess(guess)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(fguess))
	if err == nil {
		return true
	}

	return false
}

func FormatGuess(index int32) (string, error) {
	if index < 0 || index > 9999 {
		return "", fmt.Errorf("Value out of Range")
	}
	return fmt.Sprintf("%04d", index), nil
}

func GuessDigit(hash string) (int32, int64) {
	start := time.Now()

	const maxPassword int32 = 9999
	workers := runtime.GOMAXPROCS(0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make(chan int32, 1)

	var wg sync.WaitGroup
	wg.Add(workers)

	for worker := 0; worker < workers; worker++ {
		go func(startPassword int32) {
			defer wg.Done()

			for password := startPassword; password <= maxPassword; password += int32(workers) {
				if ctx.Err() != nil {
					return
				}

				if TryPassword(password, hash) {
					select {
					case results <- password:
						cancel()
					default:
					}
					return
				}
			}
		}(int32(worker))
	}

	wg.Wait()
	close(results)

	var password int32 = -1
	if result, ok := <-results; ok {
		password = result
	}

	return password, int64(time.Since(start).Seconds())
}
