package api

import (
	"encoding/json"

	"github.com/ErtelDG/user-service/config"
	"github.com/ErtelDG/user-service/models"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetUserbyId(userId int) (string, error) {
	var user models.User

	UserDB, err := config.ConnectToUserDB()
	check(err)
	defer UserDB.Close()

	row := UserDB.QueryRow("SELECT * FROM users WHERE user_id = $1", userId)
	err = row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.IsActive)
	check(err)
	userJSON, err := json.Marshal(user)
	check(err)

	return string(userJSON), err
}
