package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// main struct for holding profile data
type profile struct {
	Id   int
	Name string

	Age uint8
	//favorite programming language
	Lang string
	//favorite operating system
	OS string
	//favorite editor
	Editor string
	//last time you took a shower
	LastShower string
	//code stippet
	Code string
}

// JSON version of profile
type profileJSON struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        uint8  `json:"age"`
	Lang       string `json:"lang"`
	OS         string `json:"os"`
	Editor     string `json:"editor"`
	LastShower string `json:"lastShower"`
	Code       string `json:"code"`
}

// some random data for the profiles
var profiles = []profile{
	{Id: 0, Name: "", Age: 0, Lang: "", OS: "", Editor: "", LastShower: "", Code: ""},
	{Id: 1, Name: "Bob", Age: 20, Lang: "Go", OS: "Windows", Editor: "VS Code", LastShower: "yesterday", Code: "fmt.Println(\"Hello World\")"},
	{Id: 2, Name: "Alice", Age: 21, Lang: "Python", OS: "Linux", Editor: "Vim", LastShower: "what is a shower", Code: "print(\"Hello World\")"},
	{Id: 3, Name: "John", Age: 22, Lang: "C++", OS: "MacOS", Editor: "Xcode", LastShower: "last week", Code: "cout << \"Hello World\" << endl;"},
}

// save profiles to a file
// FIXME
func saveProfile([]profile) {
	//save profiles to a file, cause databases are hard

	//first we need to create a file
	file, err := os.Create("profiles.txt")
	if err != nil {
		log.Fatal(err)
	}
	//then we need to write to the file
	//print as comma separated values
	for _, p := range profiles {
		fmt.Fprintf(file, "%s,%d,%s,%s,%s,%s,%s \n", p.Name, p.Age, p.Lang, p.OS, p.Editor, p.LastShower, p.Code)
	}
	if err != nil {
		log.Fatal(err)
	}
	//then we need to close the file
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

// load profiles from a file
func loadProfiles() {
	//load profiles from a file
	//first we need to open the file
	file, err := os.Open("profiles.json")
	if err != nil {
		log.Fatal(err)
	}
	//then we need to read the file
	//read as comma separated values
	for {
		var p profile
		_, err := fmt.Fscanf(file, "%s,%d,%s,%s,%s,%s,%s \n", &p.Name, &p.Age, &p.Lang, &p.OS, &p.Editor, &p.LastShower, &p.Code)
		if err != nil {
			break
		}
		profiles = append(profiles, p)
	}
	if err != nil {
		log.Fatal(err)
	}
	//then we need to close the file
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

// converts profile to JSON format
func cringeToJSON(temp profile) profileJSON {
	var tempJSON profileJSON
	tempJSON.Name = temp.Name
	tempJSON.Age = temp.Age
	tempJSON.Lang = temp.Lang
	tempJSON.OS = temp.OS
	tempJSON.Editor = temp.Editor
	tempJSON.LastShower = temp.LastShower
	tempJSON.Code = temp.Code
	return tempJSON
}

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
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//default route
	myRouter.HandleFunc("/", homePage)
	//route to get a profile
	myRouter.HandleFunc("/profile/{id}", returnProfile)
	//route to create a profile
	myRouter.HandleFunc("/signup", createProfile).Methods("POST")
	//route to get a random profile
	myRouter.HandleFunc("/getBitches", getBitches)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	// random data for profile array
	handleRequests()
}
