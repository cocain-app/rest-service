package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	apiPath = "/api/"
)

func init() {
	//initialize enviromental variables
	err := godotenv.Load()
	checkErr(err, "Unable to load env variables!")
}

func main() {
	fmt.Println(calcTransScore(2.0, 0.5, 0.9887483727, 1.0, 1.0, 0.942, 0.9947259, 0.989, 0.7371124215, 0.678, 0, 0, 0, 0, 0, 0, 0, 0))

	router := mux.NewRouter()
	router.HandleFunc(apiPath+"suggestions/{id}/", getSuggestions).Methods("GET")
	router.HandleFunc(apiPath+"search/", getSongID).Methods("GET")
	router.HandleFunc(apiPath+"songs/{id}/", getSongData).Methods("GET")
	handler := cors.Default().Handler(router)
	fmt.Println(http.ListenAndServe(":8000", handler))
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
