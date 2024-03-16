package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Gunish interface {
	Gun | GunAbstract | Ammo
}

type StringOrSlice struct {
	Value []string
}

func (this *StringOrSlice) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = append([]byte{'['}, data...)
		data = append(data, ']')
	} else if data[0] != '[' && data[len(data)-1] != ']' {
		return errors.New("Not a string or slice")
	}
	return json.Unmarshal(data, &this.Value)
}

type NameOrFucingStruct struct {
	Value string
}

func (this *NameOrFucingStruct) UnmarshalJSON(data []byte) error {
	if data[0] == '{' && data[len(data)-1] == '}' {
		schema := map[string]string{"str": ""}
		err := json.Unmarshal(data, &schema)
		this.Value = schema["str"]
		return err
	} else if data[0] != '"' && data[len(data)-1] != '"' {
		return errors.New("Not a string or struct")
	}
	return json.Unmarshal(data, &this.Value)
}

type ToHit struct {
	Value float64
}

func (this *ToHit) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &this.Value)
	if err != nil {
		this.Value = 0
	}
	return nil
}

type Material struct {
	Value []string
}

func (this *Material) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = append([]byte{'['}, data...)
		data = append(data, ']')
	} else if data[0] != '[' && data[len(data)-1] != ']' {
		return errors.New("Not a string or slice")
	}
	err := json.Unmarshal(data, &this.Value)
	if err != nil {
		raw := []json.RawMessage{}
		err := json.Unmarshal(data, &raw)
		if err != nil {
			return nil
		}
		materials := make([]string, len(raw))
		schema := map[string]string{"type": ""}
		for _, bytes := range raw {
			err = json.Unmarshal(bytes, &schema)
			if err != nil {
				fmt.Printf("Couldn't read materials from %s\n", bytes)
				continue
			}
			materials = append(materials, schema["type"])
		}
		this.Value = materials
	}
	return nil
}
