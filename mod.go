package main

import (
	"fmt"
)

func get() {
	gunFiles := readDir(base + "/ammo")
	allAmmo := make([]*Ammo, 350)
	for _, file := range gunFiles {
		allAmmo = append(allAmmo, collectItems[Ammo](file)...)
	}
	ammoMap := make(map[string]*Ammo, len(allAmmo))
	for _, ammo := range allAmmo {
		if ammo == nil {
			continue
		}
		ammoMap[ammo.Id] = ammo
	}
	for _, ammo := range allAmmo {
		if ammo == nil {
			continue
		}
		if ammo.CopyFrom != "" {
			var templ *Ammo
			if ammoMap[ammo.CopyFrom] != nil {
				templ = ammoMap[ammo.CopyFrom]
			} else {
				fmt.Printf("Couldn't find template %s for %s\n", ammo.CopyFrom, ammo.Id)
				continue
			}
			mergeStructs[Ammo](ammo, templ)
		}
		fmt.Println(ammo, "\n")
	}
}
