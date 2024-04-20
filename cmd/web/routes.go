package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	// Sights
	mux.Post("/api/sight/add", dynamicMiddleware.ThenFunc(app.createSight))
	mux.Get("/api/sight/get/:id", standardMiddleware.ThenFunc(app.getSight)) //http://localhost:4000/api/sight/get/2
	mux.Get("/api/sights/get-all", standardMiddleware.ThenFunc(app.getAllSights))

	return standardMiddleware.Then(mux)
}
