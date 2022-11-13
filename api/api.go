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
	if idTemp > len(profileData) {
		fmt.Fprintf(w, "Profile not found")
		return
	}
	//tempJSON := cringeToJSON(profiles[idTemp])
	fmt.Println("Endpoint Hit: returnProfile")
	fmt.Println(profileData[idTemp])
	json.NewEncoder(w).Encode(profileData[idTemp])
}

// APT format
// {"name":"","age":,"lang":"","os":"","editor":""}
// {"name":"joe mama","age":69,"lang":"js","os":"Linux","editor":"Vim", "lastShower":"2020-10-10T10:10:10Z", "code":"console.log(\"Hello World\")"}

// create a profile from a POST request
func createProfile(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	// unmarshal this into a new Profile struct
	var temp profileJSON
	json.Unmarshal(reqBody, &temp)
	// append this to our profiles array.
	temp.Id = len(profileData)
	profileData = append(profileData, temp)
	fmt.Println(temp)
	saveProfile(profileData)
	// return the newly created profile
	fmt.Println("Endpoint Hit: createProfile")
	json.NewEncoder(w).Encode(temp)
}

// function for login and signup
func userAuth(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	// unmarshal this into a new Profile struct
	var temp accountJSON
	json.Unmarshal(reqBody, &temp)

	// check if the user exists
	for _, v := range accountData {
		if v.Username == temp.Username {
			if v.Password == temp.Password {
				fmt.Println("Endpoint Hit: userAuth")
				json.NewEncoder(w).Encode(v.Id)
				return
			} else {
				fmt.Println("Endpoint Hit: userAuth")
				json.NewEncoder(w).Encode("wrong password")
				return
			}
		}
	}
	// if the user does not exist, create a new account
	temp.Id = len(accountData)
	accountData = append(accountData, temp)
	fmt.Println("Endpoint Hit: newUserAuth")
	json.NewEncoder(w).Encode(temp.Id)
	json.NewEncoder(w).Encode(false)

}

// function for handling matches
func gotBitches(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var token int
	json.Unmarshal(reqBody, &token)
	// this will have id1 and id2, and a boolean for if they matched
	for _, v := range matchData {
		if v.Id1 == token || v.Id2 == token {
			if v.IsMutual {
				fmt.Println("Endpoint Hit: theyGotBitches")
				json.NewEncoder(w).Encode(v)
			}
		}
	}
	fmt.Println("Endpoint Hit: noBitches")
}

// handleMatch will handle the matching of two people
func handleMatch(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var token matchJSON
	json.Unmarshal(reqBody, &token)
	// this will have id1 and id2, and a boolean for if they matched, but we will ignore the boolean
	for i := 0; i < len(matchData); i++ {
		if (matchData[i].Id1 == token.Id1 && matchData[i].Id2 == token.Id2) || (matchData[i].Id1 == token.Id2 && matchData[i].Id2 == token.Id1) {
			matchData[i].IsMutual = true
			fmt.Println("Endpoint Hit: handleMatchBitches")
			return
		}
	}
	token.IsMutual = false
	matchData = append(matchData, token)
	fmt.Println("Endpoint Hit: handleMatchNoBitches")
}

// create api call for random person
func getBitches(w http.ResponseWriter, r *http.Request) {
	// return a random person
	fmt.Println("Endpoint Hit: getBitches")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	json.NewEncoder(w).Encode(profileData[r1.Intn(len(profileData))])

}

// cors middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// function for handling requests
func HandleRequests() {
	myRouter := mux.NewRouter()
	myRouter.Use(corsMiddleware)
	//default route
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/auth", userAuth).Methods("POST")
	//route to get a profile
	myRouter.HandleFunc("/profile/{id}", returnProfile)
	//route to create a profile
	myRouter.HandleFunc("/signup", createProfile).Methods("POST")
	myRouter.HandleFunc("/checkmatches", gotBitches).Methods("POST")
	myRouter.HandleFunc("/match", handleMatch).Methods("POST")
	//route to get a random profile
	myRouter.HandleFunc("/getbitches", getBitches)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
