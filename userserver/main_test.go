package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var db = getDb()

func clearDatabase(t *testing.T) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		t.Errorf("Error clearing database: %v", err)
	}
}
func TestRegisterFieldValidation(t *testing.T) {
	clearDatabase(t)

	cases := []struct {
		body           string
		expectedStatus int
		expectedError  string
	}{
		{"", http.StatusBadRequest, ErrorInvalidJSON},
		{"{\"email\":\"\", \"password\":\"bar\"}", http.StatusBadRequest, ErrorEmailBlank},
		{"{\"email\":\"foo@email.com\", \"password\":\"\"}", http.StatusBadRequest, ErrorPasswordBlank},
		{"{\"email\":\"foo\", \"password\":\"bar\"}", http.StatusBadRequest, ErrorEmailInvalid},
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", "/register", strings.NewReader(c.body))
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		rr := httptest.NewRecorder()
		http.HandlerFunc(registerHandler(db)).ServeHTTP(rr, req)

		if status := rr.Code; status != c.expectedStatus {
			t.Errorf("Status code differs. Expected \"%d\", Got \"%d\"", c.expectedStatus, status)
		}

		var restError restError
		err = json.NewDecoder(rr.Body).Decode(&restError)
		if err != nil {
			t.Errorf("Could not read error json: %v", err)
		}

		if restError.Error != c.expectedError {
			t.Errorf("Error differs. Expected \"%s\", Got: \"%s\"", c.expectedError, restError)
		}
	}
}

func TestSuccessfullRegistration(t *testing.T) {
	clearDatabase(t)

	req, err := http.NewRequest("GET", "/register", strings.NewReader("{\"email\":\"foo@email.com\", \"password\":\"bar\"}"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(registerHandler(db)).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected \"%d\", Got \"%d\"", http.StatusOK, status)
	}

	var userResponce userResponce
	err = json.NewDecoder(rr.Body).Decode(&userResponce)
	if err != nil {
		t.Errorf("Could not read userResponce json: %v", err)
	}

	if userResponce.Email != "foo@email.com" {
		t.Errorf("Email differs. Expected \"%s\", Got: \"%s\"", "foo@email.com", userResponce.Email)
	}
}

func TestRegisterDuplicateEmails(t *testing.T) {
	clearDatabase(t)

	req, _ := http.NewRequest("GET", "/register", strings.NewReader("{\"email\":\"foo@email.com\", \"password\":\"bar\"}"))
	rr := httptest.NewRecorder()
	http.HandlerFunc(registerHandler(db)).ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected OK from first /register")
	}

	req, err := http.NewRequest("GET", "/register", strings.NewReader("{\"email\":\"foo@email.com\", \"password\":\"bar\"}"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	rr = httptest.NewRecorder()
	http.HandlerFunc(registerHandler(db)).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected \"%d\", Got \"%d\"", http.StatusBadRequest, status)
	}

	var restError restError
	err = json.NewDecoder(rr.Body).Decode(&restError)
	if err != nil {
		t.Errorf("Could not read error json: %v", err)
	}

	if restError.Error != ErrorEmailExists {
		t.Errorf("Error differs. Expected \"%s\", Got: \"%s\"", ErrorEmailExists, restError)
	}
}

func TestSuccessfulLogin(t *testing.T) {
	clearDatabase(t)

	req, _ := http.NewRequest("GET", "/register", strings.NewReader("{\"email\":\"foo@email.com\", \"password\":\"bar\"}"))
	rr := httptest.NewRecorder()
	http.HandlerFunc(registerHandler(db)).ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected OK from /register")
	}

	req, err := http.NewRequest("GET", "/login", strings.NewReader("{\"email\":\"foo@email.com\", \"password\":\"bar\"}"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	rr = httptest.NewRecorder()
	http.HandlerFunc(loginHandler(db)).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected \"%d\", Got \"%d\"", http.StatusOK, status)
	}

	var tokenResponce tokenResponce
	err = json.NewDecoder(rr.Body).Decode(&tokenResponce)
	if err != nil {
		t.Errorf("Could not read tokenResponce json: %v", err)
	}

	if tokenResponce.Token == "" {
		t.Errorf("Token is blank, tokenResponce: %+v", tokenResponce)
	}
}

func TestLoginFailure(t *testing.T) {
	clearDatabase(t)

	req, _ := http.NewRequest("GET", "/register", strings.NewReader("{\"email\":\"foo@email.com\", \"password\":\"bar\"}"))
	rr := httptest.NewRecorder()
	http.HandlerFunc(registerHandler(db)).ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected OK from /register")
	}

	cases := []struct {
		body               string
		expectedError      string
		expectedReturnCode int
	}{
		{"", ErrorInvalidJSON, http.StatusBadRequest},
		{"{\"email\":\"wrongemail\", \"password\":\"bar\"}", ErrorEmailDoesntExist, http.StatusBadRequest},
		{"{\"email\":\"foo@email.com\", \"password\":\"wrongpassword\"}", ErrorPasswordIncorrect, http.StatusBadRequest},
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", "/login", strings.NewReader(c.body))
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		rr = httptest.NewRecorder()
		http.HandlerFunc(loginHandler(db)).ServeHTTP(rr, req)

		if status := rr.Code; status != c.expectedReturnCode {
			t.Errorf("Status code differs. Expected \"%d\", Got \"%d\"", c.expectedReturnCode, status)
		}

		var restError restError
		err = json.NewDecoder(rr.Body).Decode(&restError)
		if err != nil {
			t.Errorf("Could not read restError json: %v", err)
		}

		if restError.Error != c.expectedError {
			t.Errorf("Error differs. Expected \"%s\", Got: \"%s\"", c.expectedError, restError.Error)
		}
	}
}
