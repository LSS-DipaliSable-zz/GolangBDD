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
	s.AddStep(`Call the user api`, callapi)
	s.AddStep(`Satus code should be Ok`, verify_satus)
	s.Run()
}

func callapi(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	ctx.Set("response", responseData)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))
	response.StatusCode = 200
	return ctx
}

func verify_satus(t gobdd.TestingT, ctx context.Context) context.Context {
	response, err := ctx.Get("response")
	fmt.Println(response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	if err != nil {
		t.Error(err)
	}
	return ctx
}
