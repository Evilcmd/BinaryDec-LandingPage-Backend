package interfaces

import "github.com/Evilcmd/Hackup-backend/internal/models"

type UserDatabaseClient interface {
	AddUser(models.User) error
	FindUserWithEmail(string) (models.User, error)
}
