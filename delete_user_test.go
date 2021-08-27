package main_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	s := gobdd.NewSuite(t, gobdd.WithBeforeScenario(func(ctx context.Context) {
	}))
	s.AddStep(`Call the delete user api and verify the status code`, call_delete_user_api)
	s.Run()
}

func call_delete_user_api(t gobdd.TestingT, ctx context.Context) context.Context {

	request, _ := http.NewRequest(http.MethodDelete, "https://reqres.in/api/users/2", nil)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, "204 No Content", response.Status)
	assert.Equal(t, 204, response.StatusCode)
	assert.NotNil(t, response.Body)
	return ctx
}
