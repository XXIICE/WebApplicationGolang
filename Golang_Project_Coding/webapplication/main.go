// package main

// import (
// 	// Import the gorilla/mux library we just installed
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	// Declare a new router
// 	r := mux.NewRouter()

// 	// This is where the router is useful, it allows us to declare methods that
// 	// this path will be valid for
// 	r.HandleFunc("/hello", handler).Methods("GET")

// 	// We can then pass our router (after declaring all our routes) to this method
// 	// (where previously, we were leaving the second argument as nil)
// 	http.ListenAndServe(":8080", r)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// /hello in path
func TestRouter(t *testing.T) {
	// Instantiate the router using the constructor function that
	// we defined previously
	r := newRouter()

	// Create a new server using the "httptest" libraries `NewServer` method
	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	// The mock server we created runs a server and exposes its location in the
	// URL attribute
	// We make a GET request to the "hello" route we defined in the router
	resp, err := http.Get(mockServer.URL + "/hello")

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// convert the bytes to a string
	respString := string(b)
	expected := "Hello World!"

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms, that the
	// route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

// 404 test route
// func TestRouterForNonExistentRoute(t *testing.T) {
// 	r := newRouter()
// 	mockServer := httptest.NewServer(r)
// 	// Most of the code is similar. The only difference is that now we make a
// 	//request to a route we know we didn't define, like the `POST /hello` route.
// 	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// We want our status to be 405 (method not allowed)
// 	if resp.StatusCode != http.StatusMethodNotAllowed {
// 		t.Errorf("Status should be 405, got %d", resp.StatusCode)
// 	}

// 	// The code to test the body is also mostly the same, except this time, we
// 	// expect an empty body
// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	respString := string(b)
// 	expected := ""

// 	if respString != expected {
// 		t.Errorf("Response should be %s, got %s", expected, respString)
// 	}

// }
