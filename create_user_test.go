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
	s.AddStep(`Call the create user api and verify the status code will be OK`, call_create_user_api)
	s.Run()
}

func call_create_user_api(t gobdd.TestingT, ctx context.Context) context.Context {
	jsonData := map[string]string{"name": "morpheus", "job": "leader"}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post("https://reqres.in/api/users", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	assert.Equal(t, "201 Created", response.Status)
	assert.Equal(t, 201, response.StatusCode)
	assert.NotNil(t, response.Body)
	return ctx
}
