package main

import (
	"amwaywave.io/adp/server/models"
	"amwaywave.io/adp/server/pkg/languages"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, e := ioutil.ReadFile("conf/api_demo.yaml")
	if e != nil {
		panic(e)
	}

	api := models.API{}
	e = languages.ToAPI("yaml", bytes, &api)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%#v/n", api)

	newAPI, e := languages.FromAPI("yaml", &api)

	fmt.Println(string(newAPI[:]))

}
