package main_test

import (
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
	s.AddStep(`Call the user api and verify the Status code should be Ok`, callapi)
	s.Run()
}

func callapi(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseData))
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response.Body)
	return ctx
}
