package ReturnModule

import (
	"encoding/json"
	"musclemate/structs"
	"net/http"
	"strings"

	"github.com/pterm/pterm"
)

func SendResponce(w http.ResponseWriter, response any, statusCode int) {
	responseMarshal, err := json.Marshal(response)
	if err != nil {
		pterm.Error.Printf("Error converting %+v to json: %s\n", response, err)
	}

	w.WriteHeader(statusCode)

	_, err = w.Write(responseMarshal)
	if err != nil {
		pterm.Error.Println("Error writing to HTTP: ", err)
	}
}

func CustomError(w http.ResponseWriter, errorMessage string, errorCode int) {
	errorResponse := structs.ErrorResponse{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
	SendResponce(w, errorResponse, errorCode)
}

func MethodNotAllowed(w http.ResponseWriter) {
	errorMessage := "That method is not accepted at this endpoint."
	CustomError(w, errorMessage, http.StatusInternalServerError)
}

func Unauthorized(w http.ResponseWriter) {
	errorMessage := "You are unauthorized to perform this operation."
	CustomError(w, errorMessage, http.StatusUnauthorized)
}

func NotFound(w http.ResponseWriter, message ...string) {
	errorMessage := "Not found."
	if len(message) != 0 {
		errorMessage = strings.Join(message, " ")
	}
	CustomError(w, errorMessage, http.StatusNotFound)
}

func MissingData(w http.ResponseWriter, message ...string) {
	errorMessage := "There was missing data inside the request."
	if len(message) != 0 {
		errorMessage = strings.Join(message, " ")
	}
	CustomError(w, errorMessage, http.StatusInternalServerError)
}

func InternalServerError(w http.ResponseWriter, message ...string) {
	errorMessage := "There was an internal server error while trying to handle your request."
	if len(message) != 0 {
		errorMessage = strings.Join(message, " ")
	}
	CustomError(w, errorMessage, http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, message ...string) {
	errorMessage := "Bad request."
	if len(message) != 0 {
		errorMessage = strings.Join(message, " ")
	}
	CustomError(w, errorMessage, http.StatusBadRequest)
}

func SessionCreated(w http.ResponseWriter, sessionAuthentication string) {
	errorResponse := structs.SessionCreated{
		SessionAuthentication: sessionAuthentication,
	}
	SendResponce(w, errorResponse, http.StatusOK)
}

func Success(w http.ResponseWriter) {
	successResponse := structs.SuccessResponse{
		Success: true,
	}
	SendResponce(w, successResponse, http.StatusOK)
}

func WorkoutReturn(w http.ResponseWriter, workout structs.Workout) {
	SendResponce(w, workout, http.StatusOK)
}

func WorkoutCreated(w http.ResponseWriter, workoutID int64) {
	SendResponce(w, workoutID, http.StatusCreated)
}
