package utils

import (
	"os"
	"path"
)

func GoToDir(mainName string) (string, bool){
	dir, _ := os.Getwd()
	temp:=dir
	
	names, _, _, _, _:= ReadDir(temp)
	found := true
	for Contains(names, mainName)==-1{
		temp = path.Dir(temp)
		names, _, _, _, _ = ReadDir(temp)
		if(temp=="/"){
			found = false
			break
		}
	}
	if found{
		return path.Join(temp, mainName), found
	}else{
		return path.Join(dir, mainName), found
	}
}