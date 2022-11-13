package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// landing page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// returns a profile with a given id
func returnProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	idTemp, _ := strconv.Atoi(key)
	if idTemp > len(profiles) {
		fmt.Fprintf(w, "Profile not found")
		return
	}
	tempJSON := cringeToJSON(profiles[idTemp])
	fmt.Println("Endpoint Hit: returnProfile")
	fmt.Println(tempJSON)
	fmt.Println(profiles[idTemp])
	json.NewEncoder(w).Encode(tempJSON)
}

// APT format
// {"name":"","age":,"lang":"","os":"","editor":""}
// {"name":"joe mama","age":69,"lang":"js","os":"Linux","editor":"Vim", "lastShower":"2020-10-10T10:10:10Z", "code":"console.log(\"Hello World\")"}

// create a profile from a POST request
func createProfile(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var temp profile
	json.Unmarshal(reqBody, &temp)
	profiles = append(profiles, temp)
	fmt.Println(temp)
	saveProfile(profiles)
	fmt.Println("Endpoint Hit: createProfile")
	json.NewEncoder(w).Encode(temp)
}

// create api call for random person
func getBitches(w http.ResponseWriter, r *http.Request) {
	// return a random person
	fmt.Println("Endpoint Hit: getBitches")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	json.NewEncoder(w).Encode(profiles[r1.Intn(len(profiles))])

}

// function for handling requests
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//default route
	myRouter.HandleFunc("/", homePage)
	//route to get a profile
	myRouter.HandleFunc("/profile/{id}", returnProfile)
	//route to create a profile
	myRouter.HandleFunc("/signup", createProfile).Methods("POST")
	//route to get a random profile
	myRouter.HandleFunc("/getbitches", getBitches)
	//cors nonsense
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(myRouter)

	log.Fatal(http.ListenAndServe(":10000", handler))
}
