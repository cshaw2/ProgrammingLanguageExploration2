package main

import "fmt"

func main() {
   var collegeMap map[string]string
 
   collegeMap = make(map[string]string)
   
   
   collegeMap["Coastal Carolina University"] = "Conway"
   collegeMap["College of Charleston"] = "Charleston"
   collegeMap["Charleston Southern University"] = "Charleston"
   collegeMap["The Citadel"] = "Charleston"
   collegeMap["Clemson University"] = "Clemson"
   collegeMap["Francis Marion University"] = "Florence"
   collegeMap["Furman University"] = "Greenville"
   collegeMap["University of South Carolina"] = "Columbia"
   
   
   for school := range collegeMap {
      fmt.Println(school,"is a 4-year university located in",collegeMap[school])
   }
   
  
   location, ok := collegeMap["Nova Southeastern University"]
   
  
   if(ok){
      fmt.Println("This 4-year university is located in", location)  
   } else {
      fmt.Println("This university is not located in South Carolina") 
   }
}