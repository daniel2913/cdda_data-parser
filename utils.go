package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
)

func mergeStructs[T Gunish](targ *T, src *T) {
	targR := reflect.ValueOf(*targ)
	typeOfSrc := targR.Type()
	srcR := reflect.ValueOf(*src)
	for i := 0; i < targR.NumField(); i++ {
		if targR.Field(i).Interface() == reflect.New(typeOfSrc.Field(i).Type) {
			targR.Field(i).Set(srcR.Field(i))
		}
	}
}

func readDir(_path string) []*os.File {
	dir, err := os.ReadDir(_path)
	if err != nil {
		panic(fmt.Sprintf("Can't read directory %s\n", _path))
	}
	files := make([]*os.File, 20)
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		fileDesc, err := os.Open(path.Join(_path, file.Name()))
		if err != nil {
			println("Couldnt read file", file.Name())
			continue
		}
		files = append(files, fileDesc)
	}
	return files
}
