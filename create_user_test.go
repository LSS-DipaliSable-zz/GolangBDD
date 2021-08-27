package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	s := gobdd.NewSuite(t, gobdd.WithBeforeScenario(func(ctx context.Context) {
	}))
	s.AddStep(`Call the create user api`, call_create_user_api)
	s.AddStep(`Validate created user response body`, validate_created_user_response)
	s.Run()
}

func call_create_user_api(t gobdd.TestingT, ctx context.Context) context.Context {
	jsonData := map[string]string{"name": "Morpheus", "job": "Leader"}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	res := string(data)

	fmt.Println(res)

	ctx.Set("newUserInfo", res)

	assert.Equal(t, "201 Created", response.Status)
	assert.Equal(t, 201, response.StatusCode)
	assert.NotNil(t, response.Body)

	return ctx
}
func validate_created_user_response(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := ctx.Get("newUserInfo")
	if err != nil {
		t.Error(err)
	}
	//Convert the interface into string
	var x interface{} = response
	str := fmt.Sprintf("%v", x)

	//Convert the string into bytes
	data := []byte(str)

	var responseObject UserData
	json.Unmarshal(data, &responseObject)
	assert.Equal(t, responseObject.Name, "Morpheus")
	assert.Equal(t, responseObject.Job, "Leader")

	return ctx
}

type UserData struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}
