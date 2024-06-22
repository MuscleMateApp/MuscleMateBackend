package account

import (
	"encoding/json"
	"musclemate/database"
	ReturnModule "musclemate/utils/return_module"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {

	var requestData struct {
		UserID int64 `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.UserID == 0 {
		ReturnModule.BadRequest(w, "user_id is required")
		return
	} else if !database.UserExists(requestData.UserID) {
		ReturnModule.BadRequest(w, "user does not exist")
		return
	}

	user := database.GetUser(requestData.UserID)
	ReturnModule.SendResponce(w, user, http.StatusOK)

}

func PostRequest(w http.ResponseWriter, r *http.Request) {

	var requestData struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		PhoneNumber  string `json:"phone_number"`
		ProfileImage string `json:"profile_image"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.Email == "" {
		ReturnModule.BadRequest(w, "email is required")
		return
	}

	if requestData.Password == "" {
		ReturnModule.BadRequest(w, "password is required")
		return
	}

	if requestData.FirstName == "" {
		ReturnModule.BadRequest(w, "first_name is required")
		return
	}

	if requestData.LastName == "" {
		ReturnModule.BadRequest(w, "last_name is required")
		return
	}

	userID := database.CreateUser(requestData.Email, requestData.Password, requestData.FirstName, requestData.LastName, requestData.PhoneNumber, requestData.ProfileImage)
	ReturnModule.SendResponce(w, userID, http.StatusOK)

}

func PutRequest(w http.ResponseWriter, r *http.Request) {

	var requestData struct {
		UserID       int64  `json:"user_id"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		PhoneNumber  string `json:"phone_number"`
		ProfileImage string `json:"profile_image"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.UserID == 0 {
		ReturnModule.BadRequest(w, "user_id is required")
		return
	}

	if requestData.Email == "" {
		ReturnModule.BadRequest(w, "email is required")
		return
	}

	if requestData.Password == "" {
		ReturnModule.BadRequest(w, "password is required")
		return
	}

	if requestData.FirstName == "" {
		ReturnModule.BadRequest(w, "first_name is required")
		return
	}

	if requestData.LastName == "" {
		ReturnModule.BadRequest(w, "last_name is required")
		return
	}

	if requestData.PhoneNumber == "" {
		ReturnModule.BadRequest(w, "phone_number is required")
		return
	}

	if requestData.ProfileImage == "" {
		ReturnModule.BadRequest(w, "profile_image is required")
		return
	}

	if !database.UserExists(requestData.UserID) {
		ReturnModule.BadRequest(w, "user does not exist")
		return
	}

	database.UpdateUser(requestData.UserID, requestData.Email, requestData.Password, requestData.FirstName, requestData.LastName, requestData.PhoneNumber, requestData.ProfileImage)
	ReturnModule.SendResponce(w, "User updated", http.StatusOK)

}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {

	var requestData struct {
		UserID int64 `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.UserID == 0 {
		ReturnModule.BadRequest(w, "user_id is required")
		return
	}

	if !database.UserExists(requestData.UserID) {
		ReturnModule.BadRequest(w, "user does not exist")
		return
	}

	database.DeleteUser(requestData.UserID)
	ReturnModule.SendResponce(w, "User deleted", http.StatusOK)
}
