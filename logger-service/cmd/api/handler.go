package main

import (
	"logger/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	// read json var
	var requstPayload JSONPayload

	_ = app.readJSON(w, r, &requstPayload)

	// insert data

	event := data.LogEntry{
		Name: requstPayload.Name,
		Data: requstPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJOSN(w, http.StatusAccepted, resp)
}
