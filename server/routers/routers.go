package routers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/praveenmahasena647/vue-todo/dbs"
	"github.com/praveenmahasena647/vue-todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var handleCors = handlers.CORS(
	handlers.AllowedHeaders([]string{"Content-Type"}),
	handlers.AllowedMethods([]string{"GET", "DELETE", "PUT", "POST"}),
	handlers.AllowedOrigins([]string{"*"}),
)

func RunServer() {
	var router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getAll).Methods("GET")
	router.HandleFunc("/", postOne).Methods("POST")
	router.HandleFunc("/{id}", deleteOne).Methods("DELETE")
	router.HandleFunc("/", deleteAll).Methods("DELETE")
	http.ListenAndServe(":42069", handleCors(router))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	var cursor, err = dbs.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(400)
		log.Println("error during getting all")
		json.NewEncoder(w).Encode(false)
		return
	}

	var Activities []primitive.M

	for cursor.Next(context.Background()) {
		var activity bson.M
		var decodeErr = cursor.Decode(&activity)
		if decodeErr != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(false)
			log.Println(decodeErr)
			return
		}

		Activities = append(Activities, activity)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Activities)
	return
}
func postOne(w http.ResponseWriter, r *http.Request) {
	var data *models.ToDo = &models.ToDo{}

	json.NewDecoder(r.Body).Decode(&data)
	var done, err = dbs.Collection.InsertOne(context.Background(), data)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(false)
		return
	}

	w.WriteHeader(202)
	json.NewEncoder(w).Encode(done.InsertedID)
}
func deleteOne(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)["id"]
	var id, idErr = primitive.ObjectIDFromHex(params)
	if idErr != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(false)
		return
	}
	var deleted, deleteErr = dbs.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if deleteErr != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(false)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(true)
	log.Println(deleted)
}
func deleteAll(w http.ResponseWriter, r *http.Request) {
	var done, err = dbs.Collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(false)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(true)
	log.Println(done)
}
