package sessions

import (
	"github.com/LizzieM2003/golang-web-programming/27_mongodb/06_handsOn/03/models"
)

func GetSession() map[string]models.User {
	// get data from file and store in map
	m := models.LoadUsers()
	return m
}
