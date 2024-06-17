package structs

type HashedAndSaltedPassword struct {
	HashedPassword string
	RandomSalt     string
}

type SessionCreated struct {
	SessionAuthentication string `json:"session_authentication"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type Workout struct {
	WorkoutID   int    `json:"workout_id"`
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

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
