package main

import (
	"awesomeProject/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) createClientEvent(w http.ResponseWriter, r *http.Request) {
	var newClientEvent models.Client_Events

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newClientEvent)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.clientevent.Insert(&newClientEvent)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getClientEventById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":client_id")

	eventData, err := app.clientevent.GetClientEventByClientId(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(eventData)
}
