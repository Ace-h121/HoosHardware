package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var (
	ctx    context.Context
	conf   *firebase.Config
	opt    option.ClientOption
	app    *firebase.App
	client *db.Client
)

func main() {
	// Mux for routing stuff

	// All the funny little firebase stuff
	// Initialize Firebase
	ctx = context.Background()
	conf = &firebase.Config{
		DatabaseURL: "https://fro-part-sharing-default-rtdb.firebaseio.com/",
	}
	// Fetch the service account key JSON file contents
	opt = option.WithCredentialsFile("fro-part-sharing-firebase-adminsdk-2q8ix-24d2cddd19.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// set refrences to the actual
	ref := client.NewRef("")

	// create a varible to be able ot save data
	var data interface{}

	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}

	fmt.Println(data)

	fmt.Println("connected to server")

	// time to spin up the server
	StartServer()
}

func StartServer() {
	router := mux.NewRouter()

	//Functions used for Putting data

	router.HandleFunc("/Bensalem", handleOwl).Methods("POST")
	router.HandleFunc("/Allentown", handleAllentown).Methods("POST")
	router.HandleFunc("/Hatboro", handleHatboro).Methods("POST")
	router.HandleFunc("/Houston", handleHouston).Methods("POST")
	router.HandleFunc("/Lehigh", handleLehigh).Methods("POST")
	router.HandleFunc("/Montgomery", handleMontgomery).Methods("POST")
	router.HandleFunc("/MountOlive", handleOlive).Methods("POST")
	router.HandleFunc("/Seneca", handleSeneca).Methods("POST")
	router.HandleFunc("/Springside", handleSpringside).Methods("POST")
	router.HandleFunc("/WarrenHills", handleWarren).Methods("POST")

	//Functions used for Getting

	router.HandleFunc("/Bensalem", handleGetOwl).Methods("GET")
	router.HandleFunc("/Allentown", handleGetAllentown).Methods("GET")
	router.HandleFunc("/Hatboro", handleGetHatboro).Methods("GET")
	router.HandleFunc("/Houston", handleGetHouston).Methods("GET")
	router.HandleFunc("/Lehigh", handleGetLehigh).Methods("GET")
	router.HandleFunc("/Montgomery", handleGetMontgomery).Methods("GET")
	router.HandleFunc("/MountOlive", handleGetOlive).Methods("GET")
	router.HandleFunc("/Seneca", handleGetSeneca).Methods("GET")
	router.HandleFunc("/Springside", handleGetSpringside).Methods("GET")
	router.HandleFunc("/WarrenHills", handleGetWarren).Methods("GET")

	//Delete parts that have been used.

	router.HandleFunc("/BensalemRemove", handleDeleteOwl).Methods("POST")
	router.HandleFunc("/AllentownRemove", handleDeleteAllentown).Methods("POST")
	router.HandleFunc("/HatboroRemove", handleDeleteHatboro).Methods("POST")
	router.HandleFunc("/HoustonRemove", handleDeleteHouston).Methods("POST")
	router.HandleFunc("/LehighRemove", handleDeleteLehigh).Methods("POST")
	router.HandleFunc("/MontgomeryRemove", handleDeleteMontgomery).Methods("POST")
	router.HandleFunc("/MountOliveRemove", handleDeleteOlive).Methods("POST")
	router.HandleFunc("/SenecaRemove", handleDeleteSeneca).Methods("POST")
	router.HandleFunc("/SpringsideRemove", handleDeleteSpringside).Methods("POST")
	router.HandleFunc("/WarrenHillsRemove", handleDeleteWarren).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

func handleOwl(writer http.ResponseWriter, request *http.Request) {

	//Grab the Team number and part needed
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if request.Method == "OPTIONS" {
		http.Error(writer, "No Content", http.StatusNoContent)
		return
	}

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	ref := client.NewRef("/Bensalem/" + part)
	ref.Set(ctx, &num)

	writer.Write([]byte(num))
	writer.Write([]byte(part))
	fmt.Print("yo")
}

func handleAllentown(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//Grab the Team number and part needed

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	//Write to the Database
	ref := client.NewRef("/Allentown/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}
func handleSeneca(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Seneca/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}
func handleSpringside(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Springside/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}
func handleLehigh(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Lehigh/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}
func handleHatboro(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Hatboro/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}
func handleHouston(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Houston/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}

func handleMontgomery(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Montgomery/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}

func handleOlive(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Mount Olive/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}

func handleWarren(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Warren Hills/" + part)
	ref.Set(ctx, &tnum)

	fmt.Println(tnum)
	fmt.Println(part)
}

func handleDeleteOwl(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	writer.WriteHeader(http.StatusOK)

	part := request.FormValue("Part")

	ref := client.NewRef("/Bensalem/" + part)
	ref.Delete(ctx)

	fmt.Println(part)

}
func handleDeleteAllentown(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Allentown/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}
func handleDeleteSeneca(writer http.ResponseWriter, request *http.Request) {
	part := request.FormValue("Part")
	writer.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:5500")

	ref := client.NewRef("/Seneca/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}
func handleDeleteSpringside(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Springside/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}
func handleDeleteLehigh(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Lehigh/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}
func handleDeleteHatboro(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Hatboro/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}
func handleDeleteHouston(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Houston/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}

func handleDeleteMontgomery(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Montgomery/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}

func handleDeleteOlive(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Mount Olive/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}

func handleDeleteWarren(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	part := request.FormValue("Part")

	ref := client.NewRef("/Warren Hills/" + part)
	ref.Delete(ctx)

	fmt.Println(part)
}

func handleGetOwl(writer http.ResponseWriter, request *http.Request) {
	data := make(map[string]string)

	ref := client.NewRef("/Bensalem/")
	ref.Get(ctx, &data)

	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if request.Method == "OPTIONS" {
		http.Error(writer, "No Content", http.StatusNoContent)
		return
	}

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}

}
func handleGetAllentown(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Allentown/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
func handleGetSeneca(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Seneca/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
func handleGetSpringside(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Springside/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
func handleGetLehigh(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Lehigh/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
func handleGetHatboro(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Hatboro/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
func handleGetHouston(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Houston/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}

func handleGetMontgomery(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Montgomery/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}

func handleGetOlive(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Mount Olive/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}

func handleGetWarren(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	data := make(map[string]string)

	ref := client.NewRef("/Warren Hills/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
