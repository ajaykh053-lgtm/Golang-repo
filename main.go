package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// handlerHome handles requests to the root path "/"
func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Simple Go Server!\n")
	fmt.Fprintf(w, "Try visiting /hello or /time\n")
}

// handlerHello handles requests to "/hello" and greets the user
func handlerHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s! Welcome to our Go server.\n", name)
}

// handlerTime handles requests to "/time" and returns the current server time
func handlerTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "The current server time is: %s\n", currentTime)
}

func main() {
	// Define the port the server will listen on
	port := ":8080"

	// Create a new ServeMux to route requests to handlers
	mux := http.NewServeMux()

	// Register our handlers
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/hello", handlerHello)
	mux.HandleFunc("/time", handlerTime)

	fmt.Printf("Starting server on http://localhost%s...\n", port)
	fmt.Println("Press Ctrl+C to stop the server")

	// Start the server using the mux
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
