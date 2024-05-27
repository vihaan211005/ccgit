package handlers

import (
	"fmt"
	"os"
	"path"
)

func InitHandler(args []string, dir string, MainName string){
	if len(args) == 3{
		dir = path.Join(path.Dir(dir), args[2], MainName)
	}

	head_ref := path.Join("refs", "heads", "master")

	//making objects and ref folder
	os.MkdirAll(path.Join(dir, "objects"), os.ModePerm)
	os.MkdirAll(path.Join(dir, path.Dir(head_ref)), os.ModePerm)

	//writing HEAD file
	os.WriteFile(path.Join(dir,"HEAD"), []byte("ref: "+ head_ref + "\n"), os.ModePerm)

	fmt.Println("Initialized empty Git repository in " + path.Join(dir))
}