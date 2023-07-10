package data

import (
	"database/sql"
	"log"

	structure "realtimeforum/struct"

	_ "github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	db, err := sql.Open("sqlite3", "data.db") // ouvre la DB
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	// ferme la DB
	sts, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		firstName TEXT, 
		lastName TEXT, 
		username TEXT, 
		gender TEXT, 
		mail TEXT, 
		password TEXT, 
		age INT
		)`) // commande à exécuter afin de créer la table, PREPARE permet d'éviter les injections SQL (sécurité)                                                                                                                                           // execution de la commande précédente
	if err != nil {
		log.Println(err)
		return
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		firstName TEXT, 
		lastName TEXT, 
		username TEXT, 
		gender TEXT, 
		mail TEXT, 
		password TEXT, 
		age TEXT
		)`)
	//_, err = db.Exec("INSERT INTO users (firstName, lastName, username, gender, mail, password, age) VALUES (?,?,?,?,?,?,?)")
	defer sts.Close()
}

func InsertDB(object structure.UserRegister) {
	db, err := sql.Open("sqlite3", "data.db") // ouvre la DB
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	sts, err := db.Prepare("INSERT INTO users(firstName, lastName, username, gender, mail, password, age) VALUES (?,?,?,?,?,?,?)") // commande à exécuter afin de créer la table, PREPARE permet d'éviter les injections SQL (sécurité)                                                                                                                                           // execution de la commande précédente
	if err != nil {
		log.Println(err)
		return
	}

	check, _ := CheckUsernameExists(object.Username, db)

	log.Println(check)

	if !check {
		log.Println(object.FirstName, object.LastName, object.Username, object.Gender, object.Email, object.Password, object.Age)
		_, err = db.Exec("INSERT INTO users(firstName, lastName, username, gender, mail, password, age) VALUES (?,?,?,?,?,?,?)", object.FirstName, object.LastName, object.Username, object.Gender, object.Email, object.Password, object.Age)
		log.Println(err)
		if err != nil {
			return
		}
	}

	defer sts.Close()
}

func CheckUsernameExists(username string, db *sql.DB) (bool, error) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, err
}