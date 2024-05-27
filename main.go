package main

import (
	"ccgit/pkg/handlers"
	"ccgit/pkg/utils"
	"log"
	"os"
	"time"
)

var Birth time.Time = time.Date(2005, time.October, 21, 4, 30, 0, 0, time.Local)
var MainName string = ".ccgit"
var Ignore string = ".gandmrae"

func main(){
	dir, found := utils.GoToDir(MainName)
	
	args := os.Args
	if len(args)<2{
		log.Fatal("No arguements provided")
	}
	switch args[1]{
	case "chalu":
		if (found==0){
			handlers.InitHandler(args, dir, MainName)
		}else{
			log.Fatal("Already Initialised at "+dir)
		}
	case "cumit":
		if(found==1){
			handlers.CommitHandler(dir, Ignore, MainName, Birth)
		}else{
			log.Fatal("Not initialised(Initialise by chalu)")
		}
	default:
		log.Fatal("Invalid arguement")
	}
}