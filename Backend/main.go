package main

import (
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

	router.HandleFunc("/Bensalem", handleOwl).Methods("GET")
	router.HandleFunc("/Allentown", handleAllentown).Methods("GET")
	router.HandleFunc("/Hatboro", handleHatboro).Methods("GET")
	router.HandleFunc("/Houston", handleHouston).Methods("GET")
	router.HandleFunc("/Lehigh", handleLehigh).Methods("GET")
	router.HandleFunc("/Montgomery", handleMontgomery).Methods("GET")
	router.HandleFunc("/MountOlive", handleOlive).Methods("GET")
	router.HandleFunc("/Seneca", handleSeneca).Methods("GET")
	router.HandleFunc("/Springside", handleSpringside).Methods("GET")
	router.HandleFunc("/WarrenHills", handleWarren).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

func handleOwl(writer http.ResponseWriter, request *http.Request) {

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

	num := request.FormValue("TeamNum")
	part := request.FormValue("Part")

	tnum, err := strconv.Atoi(num)
	if err != nil {
		return
	}

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
