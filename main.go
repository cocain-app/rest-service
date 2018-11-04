package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	apiPath = "/api"
)

func init() {
	//initialize enviromental variables
	err := godotenv.Load()
	checkErr(err, "Unable to load env variables!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc(apiPath, isOnline).Methods("GET")
	router.HandleFunc(apiPath+"/search", getSongs).Methods("GET")
	router.HandleFunc(apiPath+"/songs/transitions/{id}", getTransitions).Methods("GET")
	router.HandleFunc(apiPath+"/songs/get/{id}", getSongDetails).Methods("GET")
	router.HandleFunc(apiPath+"/songs/get/{id}/all", getAllSongDetails).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{[]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router))
	fmt.Println("Started server!")
}

func initDB() *sql.DB {
	//get database variables
	DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS := getDBConfig()
	//build connection detail string
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	//try to establish database connection
	db, err := sql.Open("postgres", dbConfig)
	checkErr(err, "Unable to establish database connection!")

	err = db.Ping()
	checkErr(err, "Database isn't reachable!")

	fmt.Println("Successfully connected!")
	return db
}

func getDBConfig() (DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS string) {
	//retrieve enviromental variables
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")

	return DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS
}

func checkErr(err error, message string) {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err)
		panic(err)
	}
}
