package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Yoga main application
type Yoga struct {
	handlers map[string]http.HandlerFunc
}

// Data dummy data struct
type Data struct {
	Value string
}

//NewYoga Create a new app
func NewYoga(cors bool) Yoga {
	yoga := Yoga{
		handlers: make(map[string]http.HandlerFunc),
	}
	dataHandler := disableCors(yoga.GetData)
	yoga.handlers["/api/data"] = dataHandler
	yoga.handlers["/"] = http.FileServer(http.Dir("client/angular/static/")).ServeHTTP
	return yoga
}

//Serve Serve the server
func (a *Yoga) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 3030")
	return http.ListenAndServe(":3030", nil)
}

// GetData endpoint
func (a *Yoga) GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := Data{
		Value: "Hello",
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
