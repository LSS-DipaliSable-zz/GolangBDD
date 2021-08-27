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
	s.AddStep(`Call the user api`, callapi)
	s.AddStep(`Validate the response body`, validate_response)
	s.Run()
}

func callapi(t gobdd.TestingT, ctx context.Context) context.Context {
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
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response.Body)
	return ctx
}
func validate_response(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := ctx.Get("response")
	if err != nil {
		t.Error(err)
	}
	//Convert the interface into string
	var x interface{} = response
	str := fmt.Sprintf("%v", x)

	//Convert the string into bytes
	data := []byte(str)

	//Validate the response Body
	var responseObject Response
	json.Unmarshal(data, &responseObject)
	assert.Equal(t, responseObject.Page, 1)
	assert.Equal(t, responseObject.Per_page, 6)
	assert.Equal(t, responseObject.Total, 12)
	assert.Equal(t, responseObject.Total_pages, 2)

	//read the user data
	for i := 0; i < len(responseObject.User); i++ {
		fmt.Println(responseObject.User[i].Id)
		fmt.Println(responseObject.User[i].First_Name)
		fmt.Println(responseObject.User[i].Last_Name)
	}
	return ctx
}

type Response struct {
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
}
