package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	db "realtimeforum/data"
	functions "realtimeforum/functions"
	structure "realtimeforum/struct"

	"github.com/gorilla/websocket"
)

// VARIABLES
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Home Page")
	RenderTemplate(w, "index")
}

func main() {
	db.InitDB()
	http.Handle("/script/", http.StripPrefix("/script/", http.FileServer(http.Dir("./script"))))
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
	log.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")

	reader(ws)
}

// LECTURE DES MESSAGES EN PROVENANCE DU CLIENT (JS) + RECUP DONNEES
func reader(conn *websocket.Conn) {
	for {

		db, err := sql.Open("sqlite3", "data.db") // ouvre la DB
		if err != nil {
			log.Println(err)
		}
		defer db.Close()

		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var title structure.Message2
		json.Unmarshal(p, &title)

		if title.Message == "register" {
			var infoUser structure.UserRegister
			json.Unmarshal(p, &infoUser)

			log.Println(infoUser)

			functions.RegisterUser(infoUser)
		}
		if title.Message == "login" {
			var identificationUser structure.UserLogin
			json.Unmarshal(p, &identificationUser)

			log.Println(identificationUser)

			cookie := functions.LoginUser(identificationUser, db)

			data, _ := json.Marshal(cookie)

			err = conn.WriteMessage(messageType, data)
			if err != nil {
				log.Println(err)
				return
			}

			// functions.LoginUser(identificationUser)
		}

		if title.Message == "post" {
			var contentOfPost structure.Posts
			json.Unmarshal(p, &contentOfPost)
			// functions.LoginUser(contentOfPost)
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./static/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
