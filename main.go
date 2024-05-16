package main

import (
	"ccgit/makeobject"
	"ccgit/readdir"
	"ccgit/readgandmrae"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"time"
)

var Birth time.Time = time.Date(2005, time.October, 21, 4, 30, 0, 0, time.Local)
var MainName string = ".ccgit"
var Ignore string = ".gandmrae"
var dir string

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func initialize(args []string){
	if len(args) == 3{
		dir = path.Join(path.Dir(dir), args[2], MainName)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		log.Fatal("Already initialized at " + dir)
	}

	head_ref := path.Join("refs", "heads", "master")

	os.MkdirAll(path.Join(dir, "objects"), os.ModePerm)
	os.MkdirAll(path.Join(dir, path.Dir(head_ref)), os.ModePerm)

	os.WriteFile(path.Join(dir,"HEAD"), []byte("ref: "+ head_ref + "\n"), os.ModePerm)

	fmt.Println("Initialized empty Git repository in " + path.Join(dir))
}

func commit_dir(origin string, cur string, remove []string) []byte{
	names, perms, datas, isdirs,n := readdir.ReadDir(path.Join(origin, cur))
	var tree_data []byte
	for i := range n{
		if(!contains(remove, path.Join(cur, names[i]))){
			fmt.Println(names[i], perms[i], string(datas[i]), isdirs[i])
			var id string
			if(isdirs[i]){
				data := commit_dir(origin, path.Join(cur, names[i]), remove)
				id = makeobject.MakeObject(data, "tree", dir)
			}else{
				id = makeobject.MakeObject(datas[i], "blob", path.Join(origin, MainName))
			}
			id_bytes, _ := hex.DecodeString(id)
			tree_data = slices.Concat(tree_data, []byte("100" + perms[i]+ " " + names[i] + "\000"), id_bytes)
		}
	}
	return tree_data
}

func commit(){
	if _, err := os.Stat(".ccgit"); os.IsNotExist(err) {
		log.Fatal("Not initialized")
	}//Not initialized

	remove := readgandmrae.ReadGandMrae(path.Join(path.Dir(dir), Ignore))
	remove = append(remove, MainName)
	fmt.Println(remove)
	
	tree_data := commit_dir(path.Dir(dir), "", remove)
	tree_id := makeobject.MakeObject(tree_data, "tree", dir)
	
	duration := time.Since(Birth)
	seconds := strconv.Itoa(int(duration.Seconds()))

	var commit_data []byte
	commit_data = slices.Concat(commit_data, []byte("tree" + " " + tree_id + "\n" + "author " + os.Getenv("AUTHOR_NAME") + " " + os.Getenv("AUTHOR_EMAIL") + " " + seconds))
	
	commit_id := makeobject.MakeObject(commit_data, "commit", dir)
	
	os.WriteFile(path.Join(dir,"refs", "heads", "master"), []byte(commit_id), os.ModePerm)
	fmt.Println(commit_id)
}

func main(){
	dir, _ = os.Getwd()
	dir = path.Join(dir, MainName)

	args := os.Args
	if len(args)<2{
		log.Fatal("No arguements provided")
	}
	switch args[1]{
	case "chalu":
		initialize(args)
	case "cumit":
		commit()
	default:
		log.Fatal("Invalid arguement")
	}
}