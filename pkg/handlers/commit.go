package handlers

import (
	"ccgit/pkg/utils"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"time"
)

func commit_dir(origin string, cur string, remove []string, MainName string) []byte{
	names, perms, datas, isdirs,n := utils.ReadDir(path.Join(origin, cur))
	var tree_data []byte
	for i := range n{
		if(!utils.Contains(remove, path.Join(cur, names[i]))){
			fmt.Println(names[i], perms[i], string(datas[i]), isdirs[i])
			var id string
			if(isdirs[i]){
				perms[i] = "40000"
				data := commit_dir(origin, path.Join(cur, names[i]), remove, MainName)
				id = utils.MakeObject(data, "tree", path.Join(origin, MainName))
			}else{
				perms[i] = "100"+perms[i]
				id = utils.MakeObject(datas[i], "blob", path.Join(origin, MainName))
			}
			id_bytes, _ := hex.DecodeString(id)
			tree_data = slices.Concat(tree_data, []byte(perms[i]+ " " + names[i] + "\000"), id_bytes)
		}
	}
	return tree_data
}

func CommitHandler(dir string, Ignore string, MainName string, Birth time.Time){
	remove := utils.ReadGandMrae(path.Join(path.Dir(dir), Ignore))
	remove = append(remove, MainName)
	
	tree_data := commit_dir(path.Dir(dir), "", remove, MainName)
	tree_id := utils.MakeObject(tree_data, "tree", dir)

	root_commit := false
	if _, err := os.Stat(path.Join(dir,"refs","heads","master")); os.IsNotExist(err) {
		root_commit = true;
	}
	
	duration := time.Since(Birth)
	seconds := strconv.Itoa(int(duration.Seconds()))

	var commit_data []byte
	parent_commit := ""
	if(!root_commit){
		t, _ := os.ReadFile(path.Join(dir,"refs","heads","master"))
		parent_commit = "\nparent " + string(t)
	}

	commit_data = slices.Concat(commit_data, []byte("tree " + tree_id + parent_commit +"\n" + "author " + os.Getenv("AUTHOR_NAME") + " " + os.Getenv("AUTHOR_EMAIL") + " " + seconds))
	
	commit_id := utils.MakeObject(commit_data, "commit", dir)

	os.WriteFile(path.Join(dir,"refs", "heads", "master"), []byte(commit_id), os.ModePerm)
	fmt.Println(commit_id)
}