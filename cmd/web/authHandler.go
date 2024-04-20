package main

import (
	"awesomeProject/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func (app *application) signupClient(w http.ResponseWriter, r *http.Request) {
	var newClient models.Client

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newClient)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.client.Insert(newClient.Email, newClient.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
}

func (app *application) loginClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	clientId, err := app.client.Authenticate(client.Email, client.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.clientError(w, http.StatusBadRequest)
			return
		} else {
			app.serverError(w, err)

			return
		}
	}

	responseUser, err := app.client.GetUserById(strconv.Itoa(clientId))

	_, err = w.Write(responseUser)
	if err != nil {
		return
	}
}
