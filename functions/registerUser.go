package functions

import (
	//"fmt"

	data "realtimeforum/data"
	structure "realtimeforum/struct"
)

func RegisterUser(user structure.UserRegister) {
	//fmt.Println("utilisateur:", user)
	data.InitDB()
	data.InsertDB(user)
}
