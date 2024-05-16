package readdir

import (
	"fmt"
	"os"
	"path"
)

func ReadDir(dir string) ([]string, []string, [][]byte, []bool, int){
	var datas [][]byte
	var names []string
	var perms []string
	var isdirs []bool
	
	files, _ := os.ReadDir(dir)
	for _, file := range files{
		filename := file.Name()

		data, _ := os.ReadFile(path.Join(dir, filename))
		datas = append(datas, data)
		
		fileinfo, _ := file.Info()
		perm:=fmt.Sprintf("%o", fileinfo.Mode())
		perms = append(perms, perm)

		names = append(names, filename)

		isdirs = append(isdirs, file.IsDir())
	}
	return names, perms, datas, isdirs, len(files)
}