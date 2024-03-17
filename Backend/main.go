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

	router.HandleFunc("/Bensalem", handleOwl).Methods("PUT")
	router.HandleFunc("/Allentown", handleAllentown).Methods("PUT")
	router.HandleFunc("/Hatboro", handleHatboro).Methods("PUT")
	router.HandleFunc("/Houston", handleHouston).Methods("PUT")
	router.HandleFunc("/Lehigh", handleLehigh).Methods("PUT")
	router.HandleFunc("/Montgomery", handleMontgomery).Methods("PUT")
	router.HandleFunc("/MountOlive", handleOlive).Methods("PUT")
	router.HandleFunc("/Seneca", handleSeneca).Methods("PUT")
	router.HandleFunc("/Springside", handleSpringside).Methods("PUT")
	router.HandleFunc("/WarrenHills", handleWarren).Methods("PUT")

	//Functions used for Getting

	router.HandleFunc("/Bensalem", handleGetOwl).Methods("PUT")
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

	router.HandleFunc("/Bensalem", handleDeleteOwl).Methods("DELETE")
	router.HandleFunc("/Allentown", handleDeleteAllentown).Methods("DELETE")
	router.HandleFunc("/Hatboro", handleDeleteHatboro).Methods("DELETE")
	router.HandleFunc("/Houston", handleDeleteHouston).Methods("DELETE")
	router.HandleFunc("/Lehigh", handleDeleteLehigh).Methods("DELETE")
	router.HandleFunc("/Montgomery", handleDeleteMontgomery).Methods("DELETE")
	router.HandleFunc("/MountOlive", handleDeleteOlive).Methods("DELETE")
	router.HandleFunc("/Seneca", handleDeleteSeneca).Methods("DELETE")
	router.HandleFunc("/Springside", handleDeleteSpringside).Methods("DELETE")
	router.HandleFunc("/WarrenHills", handleDeleteWarren).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

func handleOwl(writer http.ResponseWriter, request *http.Request) {

	//Grab the Team number and part needed

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	ref := client.NewRef("/Bensalem/" + part)
	ref.Set(ctx, &num)

	fmt.Println(tnum)
	fmt.Println(part)
}

func handleAllentown(writer http.ResponseWriter, request *http.Request) {

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

	part := request.FormValue("Part")

	ref := client.NewRef("/Bensalem/" + part)
	ref.Delete(ctx)

	fmt.Println(part)

}
func handleDeleteAllentown(writer http.ResponseWriter, request *http.Request) {

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
func handleDeleteSeneca(writer http.ResponseWriter, request *http.Request) {

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
func handleDeleteSpringside(writer http.ResponseWriter, request *http.Request) {

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
func handleDeleteLehigh(writer http.ResponseWriter, request *http.Request) {

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
func handleDeleteHatboro(writer http.ResponseWriter, request *http.Request) {

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
func handleDeleteHouston(writer http.ResponseWriter, request *http.Request) {

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

func handleDeleteMontgomery(writer http.ResponseWriter, request *http.Request) {

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

func handleDeleteOlive(writer http.ResponseWriter, request *http.Request) {

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

func handleDeleteWarren(writer http.ResponseWriter, request *http.Request) {

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

func handleGetOwl(writer http.ResponseWriter, request *http.Request) {
	data := make(map[string]string)

	ref := client.NewRef("/Bensalem/")
	ref.Get(ctx, &data)

	for key, value := range data {
		kvw := bytes.NewBufferString(key + ":" + value + "\n")
		if _, err := kvw.WriteTo(writer); err != nil {
			log.Fatal("Error: ", err)
		}
	}

}
func handleGetAllentown(writer http.ResponseWriter, request *http.Request) {

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
