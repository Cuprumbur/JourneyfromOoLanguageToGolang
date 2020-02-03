package api

import (
	"context"
	"log"
	"net/http"

	"github.com/CuprumBur/JourneyfromOoLanguageToGolang/storage"
	"github.com/julienschmidt/httprouter"

)

type API struct {
	storage *storage.Storage
	server  *http.Server
}

func NewAPI(storage *storage.Storage) *API {
	return &API{
		storage: storage,
	}
}

func (a *API) Start(port string) error {
	a.server = &http.Server{
		Addr:    ":" + port,
		Handler: a.bootRouter(),
	}

	return a.server.ListenAndServe()
}

func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

func (a *API) bootRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/pages", a.GetAll)
	//TODO
	// router.PUT("/pages/:id", a.Update)
	// router.POST("/pages", a.Create)
	// router.GET("/pages/:id", a.Get)
	// router.DELETE("/pages/:id", a.Delete)

	return router
}

func (a *API) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	write(w, 200, []byte(`{"stub message":"Ok"}`))
}

func okResponce() []byte {
	return []byte(`{"message":"Ok"}`)
}

func write(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Printf("failed to write: %v", err)
	}
}
