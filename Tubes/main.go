package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	Depth         int     `json:"depth"`
	RecursiveTime float64 `json:"recursive_time"`
	IterativeTime float64 `json:"iterative_time"`
}

// Rekursif (simulasi pemecahan segitiga)
func sierpinskiRecursive(depth int) {
	if depth == 0 {
		return
	}
	sierpinskiRecursive(depth - 1)
	sierpinskiRecursive(depth - 1)
	sierpinskiRecursive(depth - 1)
}

// Iteratif (simulasi jumlah segitiga)
func sierpinskiIterative(depth int) {
	total := 1
	for i := 0; i < depth; i++ {
		total *= 3
	}

	for i := 0; i < total; i++ {
		_ = i * i
	}
}

// Route halaman utama
func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// Route pengujian algoritma
func test(w http.ResponseWriter, r *http.Request) {
	depthStr := r.URL.Query().Get("depth")
	depth, err := strconv.Atoi(depthStr)
	if err != nil || depth < 0 {
		depth = 5
	}

	const repeat = 10000

	start := time.Now()
	for i := 0; i < repeat; i++ {
		sierpinskiRecursive(depth)
	}
	recursiveTime := time.Since(start).Seconds()

	start = time.Now()
	for i := 0; i < repeat; i++ {
		sierpinskiIterative(depth)
	}
	iterativeTime := time.Since(start).Seconds()

	result := Result{
		Depth:         depth,
		RecursiveTime: recursiveTime,
		IterativeTime: iterativeTime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/test", test)

	println("Server berjalan di http://127.0.0.1:5000")
	http.ListenAndServe(":5000", nil)
}
