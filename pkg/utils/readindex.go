package utils

import (
	"encoding/binary"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var FileName string = "index"

type IndexData struct{
	Path string
	Id string
	Size int
}

func ReadIndex(dir string) []IndexData{
	var data []IndexData

	if _, err := os.Stat(filepath.Join(dir, FileName)); os.IsNotExist(err) {
		return data
	}

	data_bytes, err := os.ReadFile(filepath.Join(dir, FileName))
	if(err!=nil) {log.Fatal(err)}

	index := 0
	num := binary.BigEndian.Uint32(data_bytes[0:4])
	data_bytes = data_bytes[4:]

	for i := uint32(0); i < num; i++ {
		var new IndexData

		index = Contains(data_bytes, byte('\000'))
		new.Path = string(data_bytes[:index])
		data_bytes = data_bytes[index+1:]

		index = Contains(data_bytes, byte('\000'))
		new.Id = string(data_bytes[:index])
		data_bytes = data_bytes[index+1:]

		index = Contains(data_bytes, byte('\000'))
		new.Size, _ = strconv.Atoi(string(data_bytes[:index]))
		data_bytes = data_bytes[index+1:]

		data = append(data, new)
	}
	return data
}

func WriteIndex(dir string, data []IndexData){
	data_bytes := []byte("")
	var num_bytes [4]byte
	binary.BigEndian.PutUint32(num_bytes[0:4], uint32(len(data)))
	for i := 0; i < len(data); i++ {
		data_bytes = append(data_bytes, []byte(data[i].Path + "\000" + data[i].Id + "\000" + strconv.Itoa(data[i].Size) + "\000")...)
	}
	data_bytes = append(num_bytes[:], data_bytes...)

	os.WriteFile(filepath.Join(dir, FileName), data_bytes, os.ModePerm)
}