package languages

import (
	"amwaywave.io/adp/server/models"
	_ "amwaywave.io/adp/server/pkg"
	"fmt"
	"io/ioutil"
	"testing"
)

func getAPI() models.API {
	bytes, e := ioutil.ReadFile("assets/api/examples.yaml")
	if e != nil {
		panic(e)
	}

	api := models.API{}
	e = ToAPI("yaml", bytes, &api)
	if e != nil {
		panic(e)
	}
	return api
}

func TestYaml_ToAPI(t *testing.T) {
	api := getAPI()
	fmt.Printf("%#v/n", api)
}

func TestYaml_FromAPI(t *testing.T) {
	api := getAPI()
	newAPI, err := FromAPI("yaml", &api, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(newAPI[:]))
}

func TestJava_GetExtraAttr(t *testing.T) {
	data, err := GetExtraAttr("java")
	if err != nil {
		t.Error(err)
	}
	for _, e := range data {
		fmt.Printf("%#v\n", e)
	}
}

func TestJava_FromAPI(t *testing.T) {
	api := getAPI()
	out, e := FromAPI("java", &api, map[string]string{
		"Package": "io.eightpigs.models",
		"Indent":  "4",
	})
	if e != nil {
		panic(e)
	}
	fmt.Println(string(out))
}
