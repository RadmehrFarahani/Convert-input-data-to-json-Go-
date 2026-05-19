package info

import (
   "encoding/json" //using json to encode data
   "errors"
   "fmt"
   "os"
   "strings"
)

type Info struct { //create struct to store data
	Name string   `json:"Student_Name" `
	Major string  `json:"Student_Major" `
	Score string  `json:"Student_Score" `
}

func New (name,major,score string) (Info,error) { //set constructor function for info struct
  
   if name == "" || major == "" || score == "" {  //validate if user doesn't insert the value 
     return Info{} , errors.New("Invalid Input")
   }

  return Info {
	Name : name,
	Major : major,
	Score : score ,
  },nil //nil mean's we don't have errors
}

func (info Info) Display() { //set Display function to show the values for output
  fmt.Printf("%v study %v major with %v score ",info.Name,info.Major,info.Score)
}

func (info Info) Save() error { // use this function to saving data
	fileName :=strings.ReplaceAll(info.Name, " " , "_") //replace empty space with _
	fileName = strings.ToLower(fileName) + ".json" // make all of words to lowercase

	json ,err := json.Marshal(info) //convert struct to json
	if err != nil {
		return err
	}

    return os.WriteFile(fileName,json,0644)
}