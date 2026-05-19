package main

import (
  "fmt"
  "example.com/school/info" //import func and others from info.go
  "bufio" //write and read data by buffered way
  "os"
  "strings"
) 

func main(){
  name , major , score  := GetInfoData()
  studentInfo,err := info.New(name , major , score) //New function in info.go get values to set data to fields of struct
  if err != nil{
    fmt.Println(err)
	return 
  }

  studentInfo.Display() //Use Display func from info.go to show our struct values 
  err = studentInfo.Save() //Use Save func from info.go to store data and convert to json file 
  if err !=nil {
	fmt.Println("Saving the information failed.")
	return
  }
  fmt.Println("\nSaving the information succeeded.")
}

func getStudentInput(prompt string) (string) { //the function we use to get data from student
   
   fmt.Print(prompt)
   
   reader := bufio.NewReader(os.Stdin) //read the input that user type
   text , err := reader.ReadString('\n') // need string from user until user enters Enter
   if err != nil{
	return ""
   }
   //Delete \n , \r from end of the text
   text=strings.TrimSuffix(text,"\n") 
   text=strings.TrimSuffix(text,"\r")
   
   return text
}

func GetInfoData() (string,string,string) {
//the function that set the variables and get the value from user
  
  name:= getStudentInput("Student's Name :")
  major := getStudentInput("Student's Major :")
  score := getStudentInput("Student's Score :")

  return name,major,score
}