package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func collectItems[T Gunish](file *os.File) []*T {
	text, err := io.ReadAll(file)
	if err != nil {
		//panic(fmt.Sprintf("Can't read %s\n", file.Name()))
		return *new([]*T)
	}
	resSlice := make([]*T, 20)
	err = json.Unmarshal(text, &resSlice)
	if err != nil {
		fmt.Println(file.Name())
		panic(err.Error())
	}
	return resSlice
}

func getAbstracts() map[string]*GunAbstract {
	file, err := os.Open(base + "/classes/gun.json")
	if err != nil {
		panic("Cant open gun abstracts\n")
	}
	abstractSlice := collectItems[GunAbstract](file)
	abstractMap := make(map[string]*GunAbstract, 20)
	for _, abstract := range abstractSlice {
		abstractMap[abstract.Abstract] = abstract
	}
	for _, abstract := range abstractSlice {
		if abstract.CopyFrom != "" {
			mergeStructs(abstract, abstractMap[abstract.CopyFrom])
		}
	}
	return abstractMap
}

func getGuns() {
	abstracts := getAbstracts()
	gunFiles := readDir(base + "/gun")
	allGuns := make([]*Gun, 350)
	for _, file := range gunFiles {
		allGuns = append(allGuns, collectItems[Gun](file)...)
	}
	gunMap := make(map[string]*Gun, len(allGuns))
	for _, gun := range allGuns {
		if gun == nil {
			continue
		}
		gunMap[gun.Id] = gun
	}
	for _, gun := range allGuns {
		if gun == nil {
			continue
		}
		if gun.CopyFrom != "" {
			var templ *Gun
			if abstracts[gun.CopyFrom] != nil {
				templ = abstracts[gun.CopyFrom].Gun
			} else if gunMap[gun.CopyFrom] != nil {
				templ = gunMap[gun.CopyFrom]
			} else {
				fmt.Printf("Couldn't find template %s for %s\n", gun.CopyFrom, gun.Id)
				continue
			}
			mergeStructs[Gun](gun, templ)
		}
		//fmt.Println(gun, "\n")
	}
}
