package utils

import (
	"os"
	"path"
)

func GoToDir(mainName string) (string, int){
	dir, _ := os.Getwd()
	temp:=dir
	
	names, _, _, _, _:= ReadDir(temp)
	found := 1
	for !Contains(names, mainName){
		temp = path.Dir(temp)
		names, _, _, _, _ = ReadDir(temp)
		if(temp=="/"){
			found = 0
			break
		}
	}
	if found==1{
		return path.Join(temp, mainName), found
	}else{
		return path.Join(dir, mainName), found
	}
}