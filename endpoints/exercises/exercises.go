package exercises

import (
	"encoding/json"
	"musclemate/database"
	ReturnModule "musclemate/utils/return_module"
	"net/http"
)

// Exercise represents a single exercise in a workout

/*
type Exercise struct {
	ExerciseID  int    `json:"exercise_id"`
	WorkoutID   int    `json:"workout_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Reps        int    `json:"reps"`
	Sets        int    `json:"sets"`
	Weight      int    `json:"weight"`
	Duration    int    `json:"duration"`
	VideoURL    string `json:"video_url"`
	CreatedAt   string `json:"created_at"`
}
*/

func GetRequest(w http.ResponseWriter, r *http.Request) {
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

	exercises := database.GetExercises(requestData.WorkoutID)
	if len(exercises) == 0 {
		ReturnModule.NotFound(w, "exercises not found")
		return
	}

	ReturnModule.SendResponce(w, exercises, http.StatusOK)
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		WorkoutID int64   `json:"workout_id"`
		Name      string  `json:"name"`
		Sets      int     `json:"sets"`
		Reps      int     `json:"reps"`
		Weight    float64 `json:"weight"`
		Duration  int     `json:"duration"`
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

	if requestData.Sets == 0 {
		ReturnModule.BadRequest(w, "sets is required")
		return
	}

	if requestData.Reps == 0 {
		ReturnModule.BadRequest(w, "reps is required")
		return
	}

	if requestData.Weight == 0 {
		ReturnModule.BadRequest(w, "weight is required")
		return
	}

	if requestData.Duration == 0 {
		ReturnModule.BadRequest(w, "duration is required")
		return
	}

	exercise := database.CreateExercise(requestData.WorkoutID, requestData.Name, requestData.Sets, requestData.Reps, requestData.Weight, requestData.Duration, "")
	ReturnModule.SendResponce(w, exercise, http.StatusOK)
}

func PutRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ExerciseID int64   `json:"exercise_id"`
		Name       string  `json:"name"`
		Sets       int     `json:"sets"`
		Reps       int     `json:"reps"`
		Weight     float64 `json:"weight"`
		Duration   int     `json:"duration"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.ExerciseID == 0 {
		ReturnModule.BadRequest(w, "exercise_id is required")
		return
	}

	if requestData.Name == "" {
		ReturnModule.BadRequest(w, "name is required")
		return
	}

	if requestData.Sets == 0 {
		ReturnModule.BadRequest(w, "sets is required")
		return
	}

	if requestData.Reps == 0 {
		ReturnModule.BadRequest(w, "reps is required")
		return
	}

	if requestData.Weight == 0 {
		ReturnModule.BadRequest(w, "weight is required")
		return
	}

	if requestData.Duration == 0 {
		ReturnModule.BadRequest(w, "duration is required")
		return
	}

	exercise := database.UpdateExercise(requestData.ExerciseID, requestData.Name, requestData.Sets, requestData.Reps, requestData.Weight, requestData.Duration, "")
	ReturnModule.SendResponce(w, exercise, http.StatusOK)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ExerciseID int64 `json:"exercise_id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	if requestData.ExerciseID == 0 {
		ReturnModule.BadRequest(w, "exercise_id is required")
		return
	}

	if !database.DeleteExercise(requestData.ExerciseID) {
		ReturnModule.InternalServerError(w, "Internal server error")
		return
	}

	ReturnModule.Success(w)
}
