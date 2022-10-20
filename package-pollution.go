package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// START OMIT

const url = "https://api.example.com"

type apiResponse struct {
	ID      string
	Options []string
}

func main() {
	result := doThing(r)
	fmt.Println(result)
}

func doThing() string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	var r apiResponse
	_ = json.NewDecoder(resp.Body).Decode(&r)
	return r.ID
}

// END OMIT

// START FIXED OMIT

func main() {
	result := doThing(r)
	fmt.Println(result)
}

func doThing() string {
	// does make constants harder to edit since they aren't
	// at the top of files but it does make the scope of the data
	// more clear
	const url = "https://api.example.com"

	// response struct is now local to the function
	// hard to see any downside here
	type apiResponse struct {
		ID      string
		Options []string
	}

	resp, _ := http.Get("api_addr")
	defer resp.Body.Close()
	var r apiResponse
	_ = json.NewDecoder(resp.Body).Decode(&r)
	return r.ID
}

// END FIXED OMIT
