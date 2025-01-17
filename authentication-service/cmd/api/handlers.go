package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Repo.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invalid credentials!"), http.StatusBadRequest)
		return
	}

	valid, err := app.Repo.PasswordMatches(requestPayload.Password, *user)

	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invalid credentials!"), http.StatusBadRequest)
		return
	}

	// log authentication
	err = app.logRequest("authentication", fmt.Sprintf("%s logged in!", user.Email))
	// err = app.logRequestBroker("authentication", fmt.Sprintf("%s logged in!", user.Email))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) logRequest(name, data string) error {

	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceUrl := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	_, err = app.Client.Do(request)

	if err != nil {
		return err
	}

	return nil
}

type Log struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type LogRequest struct {
	Action string `json:"action"`
	Log    Log    `json:"log"`
}

func (app *Config) logRequestBroker(name, data string) error {

	logRequest := LogRequest{
		Action: "log",
		Log: Log{
			Name: name,
			Data: data,
		},
	}

	jsonData, _ := json.MarshalIndent(logRequest, "", "\t")
	logServiceUrl := "http://broker-service:8080/handle"

	request, err := http.NewRequest("POST", logServiceUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)

	if err != nil {
		return err
	}

	return nil
}
