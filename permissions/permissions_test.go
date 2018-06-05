package permissions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/armory/plank/client"
)

func TestAdmin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockBody := `
		{
			"name": "somebody",
			"admin": true
		}`
		fmt.Fprintln(w, mockBody)
	}))
	defer ts.Close()
	var err error
	defaultFiatClient, err = client.New(client.BaseURL(ts.URL))
	assert.Nil(t, err)
	output, _ := Admin("somebody")
	assert.True(t, output)
}

func TestWrite(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockBody := `
		{
			"name": "somebody",
			"admin": false,
			"applications": [
				{
					"name": "myapp",
					"authorizations" : ["WRITE", "READ"]
				}
			]
		}`
		fmt.Fprintln(w, mockBody)
	}))
	defer ts.Close()
	var err error
	defaultFiatClient, err = client.New(client.BaseURL(ts.URL))
	assert.Nil(t, err)
	output, err := WriteAccess("somebody", "myapp")
	assert.Nil(t, err)
	assert.True(t, output)
}
