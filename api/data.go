package api

import (
	"fmt"
	"log"
	"os"
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
