package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)


func fetchWeatherCity(city string, ch chan<-string, wg *sync.WaitGroup) interface{} {
    var data struct {
        Main struct {
            Temp float64 `json:"temp"`
        } `json:"main"`
    }

    defer wg.Done()

    apiKey := os.Getenv("API_KEY") 

    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error fetching weather for %s: %s\n", city, err)
        return data
    }

    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
        return data
    }

    ch<-fmt.Sprintf("This is the %s", city)

    return data
}

func RunWeather() {
    startNow := time.Now()

    cities := []string{"Toronto", "London", "Ankara", "Istanbul", "New York"}

    ch := make(chan string)
    var wg sync.WaitGroup
    
    for _, city := range(cities) {
        wg.Add(1)
        go fetchWeatherCity(city, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for result := range ch {
        fmt.Println(result)
    }

    fmt.Println("This operation took: ", time.Since(startNow))

}
