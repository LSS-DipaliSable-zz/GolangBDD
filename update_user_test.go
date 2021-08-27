package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	s := gobdd.NewSuite(t, gobdd.WithBeforeScenario(func(ctx context.Context) {
	}))
	s.AddStep(`Call the update user api`, call_update_user_api)
	s.AddStep(`Validate the updated user response`, validate_updated_user_response)
	s.Run()
}

func call_update_user_api(t gobdd.TestingT, ctx context.Context) context.Context {
	jsonData := map[string]string{"name": "Luis", "job": "Team Leader"}
	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest(http.MethodPut, "https://reqres.in/api/users/5", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)

	ctx.Set("update_user_response", bodyString)

	assert.Equal(t, "200 OK", response.Status)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response.Body)
	return ctx
}
func validate_updated_user_response(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := ctx.Get("update_user_response")
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the interface into string
	var i interface{} = response
	str := fmt.Sprintf("%v", i)

	//Convert the string into bytes
	bodyBytes := []byte(str)

	var userData UpdateUserData
	json.Unmarshal(bodyBytes, &userData)

	assert.Equal(t, userData.Name, "Luis")
	assert.Equal(t, userData.Job, "Team Leader")

	return ctx
}

type UpdateUserData struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}
