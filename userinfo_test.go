package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
	"github.com/stretchr/testify/assert"
)

func TestFeatures(t *testing.T) {
	s := gobdd.NewSuite(t, gobdd.WithBeforeScenario(func(ctx context.Context) {
	}))
	s.AddStep(`Call the user api`, call_api)
	s.AddStep(`Validate the response body`, validate_response)
	s.Run()
}

func call_api(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	res := string(responseData)
	fmt.Println(res)

	ctx.Set("response", res)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "200 OK", response.Status)
	return ctx
}

func validate_response(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := ctx.Get("response")
	fmt.Println(response)

	responseObject := `{"page":1,"per_page":6,"total":12,"total_pages":2,"data":[{"id":1,"email":"george.bluth@reqres.in","first_name":"George","last_name":"Bluth","avatar":"https://reqres.in/img/faces/1-image.jpg"},{"id":2,"email":"janet.weaver@reqres.in","first_name":"Janet","last_name":"Weaver","avatar":"https://reqres.in/img/faces/2-image.jpg"},{"id":3,"email":"emma.wong@reqres.in","first_name":"Emma","last_name":"Wong","avatar":"https://reqres.in/img/faces/3-image.jpg"},{"id":4,"email":"eve.holt@reqres.in","first_name":"Eve","last_name":"Holt","avatar":"https://reqres.in/img/faces/4-image.jpg"},{"id":5,"email":"charles.morris@reqres.in","first_name":"Charles","last_name":"Morris","avatar":"https://reqres.in/img/faces/5-image.jpg"},{"id":6,"email":"tracey.ramos@reqres.in","first_name":"Tracey","last_name":"Ramos","avatar":"https://reqres.in/img/faces/6-image.jpg"}],"support":{"url":"https://reqres.in/#support-heading","text":"To keep ReqRes free, contributions towards server costs are appreciated!"}}`
	responseBytes := []byte(responseObject)

	var res ResponseData

	json.Unmarshal(responseBytes, &res)
	assert.Equal(t, res.Page, 1)
	assert.Equal(t, res.Per_page, 6)
	assert.Equal(t, res.Total, 12)
	assert.Equal(t, res.Total_pages, 2)

	//read the user data
	for i := 0; i < len(res.User); i++ {
		fmt.Println(res.User[i].Id)
		fmt.Println(res.User[i].First_Name)
		fmt.Println(res.User[i].Last_Name)
	}

	if err != nil {
		t.Error(err)
	}
	return ctx
}

type ResponseData struct {
	Page        int    `json:"page"`
	Per_page    int    `json:"per_page"`
	Total       int    `json:"total"`
	Total_pages int    `json:"total_pages"`
	User        []User `json:"data"`
}

type User struct {
	Id         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}
