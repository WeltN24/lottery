package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rcrowley/go-metrics/exp"

	"github.com/rcrowley/go-metrics"

	"net/http"
)

const (
	max   = 49
	count = 5
)

var (
	timer = metrics.NewTimer()
	h     = metrics.NewHistogram(metrics.NewUniformSample(max))
)

func init() {
	metrics.Register("numbers", h)
	metrics.Register("timer", timer)
	exp.Exp(metrics.DefaultRegistry)
}

func main() {
	http.HandleFunc("/numbers", numbers)
	http.HandleFunc("/health", health)
	log.Println("staring server ...")
	http.ListenAndServe(":8080", nil)
}

func health(rw http.ResponseWriter, req *http.Request) {
	if n := rand.Intn(50); n == 0 {
		http.Error(rw, "down", http.StatusInternalServerError)
	}
	fmt.Fprint(rw, "up")
}

func numbers(rw http.ResponseWriter, req *http.Request) {
	defer timer.UpdateSince(time.Now())
	numbers := make([]int, 0, count)
	used := make([]bool, max)
	for i := 0; i < cap(numbers); i++ {
		for {
			if n := rand.Intn(max); !used[n] {
				numbers = append(numbers, n+1)
				h.Update(int64(n + 1))
				break
			}
		}
	}
	err := json.NewEncoder(rw).Encode(numbers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
