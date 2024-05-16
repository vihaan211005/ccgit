package makeobject

import (
	"ccgit/compress"
	"ccgit/sha1"
	"os"
	"path"
	"slices"
	"strconv"
)

func MakeObject(data []byte, option string, dir string) string{
	tostart := []byte(option + " " + strconv.Itoa(len(data)) + "\000")
	towrite := slices.Concat(tostart, data)
	towrite_compressed := compress.Compress(towrite)
	id := sha1.Sha1(towrite)

	os.Mkdir(path.Join(dir, "objects", id[:2]), os.ModePerm)
	os.WriteFile(path.Join(dir, "objects", id[:2], id[2:]), towrite_compressed, os.ModePerm)
	
	return id
}