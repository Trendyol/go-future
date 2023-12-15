package main

import (
	"fmt"
	"github.com/Trendyol/go-future/future"
	"log"
	"math/rand"
	"time"
)

func main() {
	defer TimeTrack(time.Now(), "MultipleRequest")
	ids := []string{"A", "B", "C", "D", "E"}
	futGroup := future.Group[string]{}
	for i := range ids {
		id := ids[i]
		futGroup.Go(func() (string, error) {
			return APICall(id)
		})
	}
	log.Println("Waiting for future result...")
	results, err := futGroup.Get()
	if err != nil {
		log.Println("An error occurred...")
		return
	}
	for i := range results {
		log.Printf("Result: %s", results[i])
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %d ms", name, elapsed.Milliseconds())
}

func APICall(request string) (string, error) {
	start := time.Now()
	r := rand.NewSource(start.UnixNano())
	waitMillis := rand.New(r).Intn(1000)
	time.Sleep(time.Duration(waitMillis) * time.Millisecond)
	return fmt.Sprintf("Reqeuest %s / response time %d ms", request, time.Since(start).Milliseconds()), nil
}
