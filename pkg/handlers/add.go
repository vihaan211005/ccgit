package handlers

import (
	"ccgit/pkg/utils"
	"fmt"
	"os"
	"path/filepath"
)

var paths []string

func visit(path string, f os.FileInfo, err error) error {
	
	if !f.IsDir(){
		paths = append(paths, path)
	}
	
	return nil
} 

func AddHandler(args []string, dir string){
	origin := filepath.Dir(dir)

	if(len(args)==2){
		fmt.Println("Kuch dalne ko to de")
	}

	inIndex := utils.ReadIndex(dir)
	for i := 0; i < len(inIndex); i++ {
		inIndex[i].Path = filepath.Join(origin, inIndex[i].Path)
	}
	for i := 2;i<(len(args));i++{
		toAdd,_ := filepath.Abs(args[i])
		
		if(!filepath.HasPrefix(toAdd, origin)){
			fmt.Println(toAdd + " not in " + origin)
			continue
		}

		for i := 0; i < len(inIndex); i++ {
			if(filepath.HasPrefix(inIndex[i].Path, toAdd)){
				inIndex = append(inIndex[:i], inIndex[i+1:]...)
			}
		}

		if _, err := os.Stat(toAdd); !os.IsNotExist(err) {
			filepath.Walk(toAdd, visit)
		}
	}
	fmt.Println(paths, inIndex)
	for i := 0; i < len(paths); i++ {
		var new utils.IndexData
		new.Path = paths[i]
		data_bte,_ := os.ReadFile(paths[i])
		new.Id = utils.MakeObject(data_bte, "blob", dir)
		inIndex = append(inIndex, new)
	}
	for i := 0; i < len(inIndex); i++ {
		inIndex[i].Path,_ = filepath.Rel(origin, inIndex[i].Path)
	}
	utils.WriteIndex(dir, inIndex)
}