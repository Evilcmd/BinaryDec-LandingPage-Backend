package apis

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Evilcmd/Hackup-backend/internal/models"
	"github.com/Evilcmd/Hackup-backend/internal/utilities"
)

func (apiCfg *ApiConfig) Presignup(res http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, "error reading request body")
		return
	}

	user := models.User{}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, "error unmarshalling request body")
		return
	}

	if user.Name == "" || user.Email == "" || user.Company == "" {
		utilities.RespondWithError(res, http.StatusBadRequest, "did not find name, email or company")
		return
	}

	err = apiCfg.UserDbClient.AddUser(user)
	if err == models.ErrorUserExists {
		utilities.RespondWithError(res, 409, "user already exists")
		return
	}
	if err != nil {
		utilities.RespondWithError(res, http.StatusInternalServerError, "error adding user to databse")
		return
	}

	utilities.RespondWithJson(res, 200, "OK")
}
