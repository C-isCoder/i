package test

import (
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"testing"
	"i/main/test/handlers"
)

const hCheckMark = "\u2713"
const hBallotX = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShoud be able to create a request.", hBallotX, err)
		}
		t.Log("\tShould be able to create a request.", hCheckMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", hBallotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", hCheckMark)
		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShoule decode the response.", hBallotX)
		}
		t.Log("\tShould decode the response.", hCheckMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", hCheckMark)
		} else {
			t.Error("\tShould have a Name.", hBallotX, u.Name)
		}
		if u.Email == "bill@mail.com" {
			t.Log("\tShould have a Email.", hCheckMark)
		} else {
			t.Error("\tShould have a Email.", hBallotX, u.Email)
		}

	}
}
