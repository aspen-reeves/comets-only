package api

import (
	"fmt"
	"log"
	"os"
)

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

// more dummy data
var profileData = []profileJSON{
	{Id: 0, Name: "", Age: 0, Lang: "", OS: "", Editor: "", LastShower: "", Code: ""},
	{Id: 1, Name: "Bob", Age: 20, Lang: "Go", OS: "Windows", Editor: "VS Code", LastShower: "yesterday", Code: "fmt.Println(\"Hello World\")"},
	{Id: 2, Name: "Alice", Age: 21, Lang: "Python", OS: "Linux", Editor: "Vim", LastShower: "what is a shower", Code: "print(\"Hello World\")"},
	{Id: 3, Name: "John", Age: 22, Lang: "C++", OS: "MacOS", Editor: "Xcode", LastShower: "last week", Code: "cout << \"Hello World\" << endl;"},
}

// save profiles to a file in JSON format
func saveProfile([]profileJSON) {
	//save as whole JSON object
	//first we need to create a file
	file, err := os.Create("profiles.json")
	if err != nil {
		log.Fatal(err)
	}
	//then we need to write to the file
	//just print the whole JSON object
	for _, p := range profileData {
		fmt.Fprintln(file, p)
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
	//read as JSON object
	for _, p := range profileData {
		fmt.Fprintln(file, p)
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
