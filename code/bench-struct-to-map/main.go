// package jsonhack
// the result we want is just a map (json is irrelevent in our usecase)
// just each struct field mapped to an interace type of the it's respective value
package jsonhack

import (
	"encoding/json"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Bar struct {
	Bar string `json:"bar"`
}

type Foo struct {
	A Bar    `json:"a"`
	B Bar    `json:"b"`
	C Bar    `json:"c"`
	D Bar    `json:"d"`
	E Bar    `json:"e"`
	F int    `json:"f"`
	G uint   `json:"g"`
	H string `json:"h"`
}

func base(foo Foo) map[string]interface{} {

	result := make(map[string]interface{})

	result["A"] = foo.A
	result["B"] = foo.B
	result["C"] = foo.C
	result["D"] = foo.D
	result["E"] = foo.E
	result["F"] = foo.F
	result["G"] = foo.G
	result["H"] = foo.H
	return result
}

func hack(foo Foo) map[string]interface{} {

	result := make(map[string]interface{})

	bytes, _ := json.Marshal(foo)
	json.Unmarshal(bytes, result)

	return result
}

func reflec(foo Foo) map[string]interface{} {

	result := make(map[string]interface{})

	val := reflect.ValueOf(foo)
	typ := reflect.TypeOf(foo)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		result[typ.Field(i).Name] = field.Interface()
	}

	return result
}

func mitchellh(foo Foo) map[string]interface{} {
	result := make(map[string]interface{})

	mapstructure.Decode(&foo, &result)

	return result
}
