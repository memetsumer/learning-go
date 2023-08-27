package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	id int
	name string 
}

func cn() {
	now := time.Now()
	userID := 10
	respch := make(chan string, 3)
	wg := &sync.WaitGroup{}

	go fetchUserData(userID, respch, wg)
	go fetchUserRecommendations(userID, respch, wg)
	go fetchUserLikes(userID, respch, wg)
	wg.Add(3)
	wg.Wait()

	close(respch)

	// fmt.Println(userData)
	// fmt.Println(userRecs)
	// fmt.Println(userLikes)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	respch <- "user data"

	wg.Done()
}

func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	respch <- "user recommendations"

	wg.Done()
}

func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	respch <- "user likes"

	wg.Done()
}
