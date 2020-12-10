package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Task struct {
	ID int `json:ID`
	Name string `json:Name`
	Content string `json:Content`
}

type User struct {
	UserID int `json:UserID`
	UserName string `json:UserName`
	UserEmail string `json:UserEmail`

}

type allTasks []Task

var tasks = allTasks {
	{
		ID: 1,
		Name: "Tarea 1",
		Content: "Contenido de la tarea",
	},
	{
		ID: 2,
		Name: "Tarea 2",
		Content: "Contenido de la tarea",
	},
	{
		ID: 3,
		Name: "Tarea 3",
		Content: "Contenido de la tarea",
	},
	{
		ID: 4,
		Name: "Tarea 4",
		Content: "Contenido de la tarea",
	},
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Funciona la raíz")

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	// Recibimos el body desde el Request
	req, err := ioutil.ReadAll(r.Body)

	// Devolvemos un error si lo encontramos
	if err != nil {
		fmt.Fprintf(w, "Error, inserta datos válidos")
	}

	json.Unmarshal(req, &newTask)

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	router := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*/*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	router.HandleFunc("/", index)
	router.HandleFunc("/tareas", getTasks)
	router.HandleFunc("/crearTarea", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}