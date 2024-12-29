package main

import (
	"encoding/json"
	"fmt"

	"io"
	"log"
	"net/http"
	"os"

	// "a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Initialize the services
var fileService = &service.FileService{}
var aiService = &service.AIService{Client: &http.Client{}}
var store = sessions.NewCookieStore([]byte("my-key"))

func getSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "chat-session")
	return session
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the Hugging Face token from the environment variables
	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		log.Fatal("HUGGINGFACE_TOKEN is not set in the .env file")
	}

	// Set up the router
	router := mux.NewRouter()

	// File upload endpoint
	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Please Upload the File First!", http.StatusBadRequest)
			return
		}

		question := r.FormValue("initial_query")
		if question == "" {
			http.Error(w, "Atleast Give an Some Question about the CSV!", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		filePath, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Please Upload the CSV File Only!", http.StatusBadRequest)
			return
		}

		respFile, err := fileService.ProcessFile(string(filePath))
		if err != nil {
			http.Error(w, "Unable to Read File. Please Upload with Format Correctly!", http.StatusInternalServerError)
			return
		}

		resp, err := aiService.AnalyzeData(respFile, question, token)
		if err != nil {
			http.Error(w, "Unable to Analyst File. Please Upload with Format Correctly!", http.StatusInternalServerError)
			return
		}

		fmt.Println(resp)

		response := map[string]string{
			"status": "success",
			"answer": resp,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		fmt.Println(response)
	}).Methods("POST")

	// Chat endpoint
	router.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		question := r.FormValue("query")
		if question == "" {
			http.Error(w, "Please Atleast Give me An Question!", http.StatusBadRequest)
		}

		resAI, _ := aiService.ChatWithAI("Text from previous conversation", question, token)
		// if err != nil {
		// 	http.Error(w, "Error with Model / Your Question. Please Try Again!", http.StatusInternalServerError)
		// 	return
		// }

		response := map[string]string{
			"status": "success",
			"answer": resAI.GeneratedText,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("POST")

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Allow your React app's origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))

}
