package main

import (
    "fmt"
    "math/rand"
    "sort"
    "sync"
    "time"
)

const MAX_CONCURRENCY int = 365

func getStonks() {

}

func main() {
    min := 93
    max := 100

    answers := []string{
        "green",
        "red",
    }

    var wg sync.WaitGroup

    date := time.Now()

    results := []string{}
    resultsChan := make(chan string, MAX_CONCURRENCY)

    for i := 0; i < MAX_CONCURRENCY; i++ {
        wg.Add(1)

        go func(iteration int) {
            defer func() {
                wg.Done()

            }()
            rand.Seed(time.Now().UTC().UnixNano())

            result := fmt.Sprintf("%v will be a %v day (Accuracy: %v%%)",
                date.AddDate(0, 0, iteration).Format("2006-01-02"),
                answers[rand.Intn(len(answers))],
                rand.Intn(max-min)+min,
            )

            resultsChan <- result
        }(i)
    }

    go func() {
        for j := range resultsChan {
            results = append(results, j)
        }
    }()

    wg.Wait()

    sort.Slice(results, func(i, j int) bool {
        return results[i] < results[j]
    })

    for _, i := range results {
        fmt.Println(i)
    }

}
