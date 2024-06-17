package workouts

import (
	"encoding/json"
	"musclemate/database"
	ReturnModule "musclemate/utils/return_module"
	"net/http"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		WorkoutID int64 `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.WorkoutID == 0 {
		ReturnModule.BadRequest(w, "workout_id is required")
		return
	}

	workout := database.GetWorkout(requestData.WorkoutID)
	if workout.UserID == 0 {
		ReturnModule.NotFound(w, "workout not found")
		return
	}

	ReturnModule.WorkoutReturn(w, workout)

}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		UserID      int64  `json:"user_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
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

	if requestData.Name == "" {
		ReturnModule.BadRequest(w, "name is required")
		return
	}

	if requestData.Description == "" {
		ReturnModule.BadRequest(w, "description is required")
		return
	}

	workoutID := database.CreateWorkout(requestData.UserID, requestData.Name, requestData.Description)
	if workoutID == 0 {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	ReturnModule.WorkoutCreated(w, workoutID)
}

func PutRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		WorkoutID   int64  `json:"workout_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.WorkoutID == 0 {
		ReturnModule.BadRequest(w, "workout_id is required")
		return
	}

	if requestData.Name == "" {
		ReturnModule.BadRequest(w, "name is required")
		return
	}

	if requestData.Description == "" {
		ReturnModule.BadRequest(w, "description is required")
		return
	}

	if !database.UpdateWorkout(requestData.WorkoutID, requestData.Name, requestData.Description) {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	ReturnModule.Success(w)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		WorkoutID int64 `json:"workout_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.WorkoutID == 0 {
		ReturnModule.BadRequest(w, "workout_id is required")
		return
	}

	if !database.DeleteWorkout(requestData.WorkoutID) {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	ReturnModule.Success(w)
}

