package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io/ioutil"

	_ "github.com/gorilla/mux"
)

type Task struct {
	ID int `json:ID`
	Name string `json:Name`
	Content string `json:Content`
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/tareas", getTasks)
	router.HandleFunc("/crearTarea", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}