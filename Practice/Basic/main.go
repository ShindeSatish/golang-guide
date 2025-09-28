package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Closure example 1: Basic adder
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// Closure example 2: Event handlers with context
	saveHandler := createButtonHandler("user123", "save")
	deleteHandler := createButtonHandler("user123", "delete")

	saveHandler()
	deleteHandler()

	// Closure example 3: Middleware with logging
	// Create a logger instance (this was missing!)
	myLogger := log.New(os.Stdout, "[HTTP] ", log.LstdFlags)
	
	// Create the middleware
	logMiddleware := withLogging(myLogger)

	//Create another middleware
	authMiddleware := withAuth(myLogger)
	
	// Create a sample handler
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	
	// Wrap the handler with logging middleware
	wrappedHandler := logMiddleware(authMiddleware(helloHandler))
	
	// Demonstrate the middleware (simulate a request)
	fmt.Println("\n--- Demonstrating HTTP Middleware ---")
	fmt.Println("In a real server, this would log actual HTTP requests")
	fmt.Println("For demo purposes, we're showing how the closure captures the logger:")
	
	// Note: To actually test this, you'd need to start an HTTP server
	http.Handle("/", wrappedHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
	_ = wrappedHandler // Prevent unused variable warning
}

func createButtonHandler(userID string, action string) func() {
	return func() {
		fmt.Printf("User %s performed %s\n", userID, action)
		// Handle the specific action for this user
	}
}

func withLogging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("Request: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}

func withAuth(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("Authenticated: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}