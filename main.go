package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "songTitle"},
		Debug:            false, //TODO: Disable for production
	})

	handler := cors.Handler(router)

	server := &http.Server{
		Addr: "127.0.0.1:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler, // Pass our instance of gorilla/mux in.
	}
	log.Println("Configured server")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Started server")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	log.Println("Server Status: offline \n\n\n")
	os.Exit(0)
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
