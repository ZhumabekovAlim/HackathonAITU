package main

import (
	"awesomeProject/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) createClientSight(w http.ResponseWriter, r *http.Request) {
	var newClientSight models.Client_Sights

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newClientSight)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.clientsight.Insert(&newClientSight)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getClientSightById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":client_id")

	eventData, err := app.clientsight.GetClientSightByClientId(id)
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
