package functions

import (
	"database/sql"
	"log"
	"strings"

	structure "realtimeforum/struct"
)

func LoginUser(user structure.UserLogin, db *sql.DB) structure.Cookie {
	id := user.UsernameMail
	password := user.Password

	var cookie structure.Cookie

	// On vérifies si l'id est un mail ou un username
	if len(strings.Split(id, "@")) > 1 {
		// L'id est un mail
		checkPassword, checkUsername := GetPasswordAndMailWithUsername(id, "mail", db)

		if checkPassword == password {
			cookie = structure.Cookie{
				Message: "cookie",
				Value: checkUsername,
			}
		}

	} else {
		// L'id est un username
		checkPassword, checkUsername := GetPasswordAndMailWithUsername(id, "username", db)

		if checkPassword == password {
			cookie = structure.Cookie{
				Message: "cookie",
				Value: checkUsername,
			}
		}
	}

	return cookie
}

func GetPasswordAndMailWithUsername(id, cequonveux string, db *sql.DB) (string, string) {
	var value1, value2 string

	if cequonveux == "mail" {

		err := db.QueryRow("SELECT password FROM users WHERE mail = ?", id).Scan(&value1)
		if err != nil {
			log.Println(err)
			return value1, value2
		}

		err = db.QueryRow("SELECT username FROM users WHERE mail = ?", id).Scan(&value2)
		if err != nil {
			log.Println(err)
			return value1, value2
		}

	} else {

		err := db.QueryRow("SELECT password FROM users WHERE username = ?", id).Scan(&value1)
		if err != nil {
			log.Println(err)
			return value1, value2
		}

		err = db.QueryRow("SELECT username FROM users WHERE username = ?", id).Scan(&value2)
		if err != nil {
			log.Println(err)
			return value1, value2
		}
	}
	

	return value1, value2
}
