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
	WorkoutID   int64      `json:"workout_id"`
	UserID      int64      `json:"user_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   string     `json:"created_at"`
	Exercises   []Exercise `json:"exercises"`
}

type Exercise struct {
	ExerciseID int64   `json:"exercise_id"`
	WorkoutID  int64   `json:"workout_id"`
	Name       string  `json:"name"`
	Sets       int     `json:"sets"`
	Reps       int     `json:"reps"`
	Weight     float64 `json:"weight"`
	Duration   int     `json:"duration"`
	VideoURL   string  `json:"video_url"`
	CreatedAt  string  `json:"created_at"`
}

type User struct {
	UserID       int    `json:"user_id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	ProfileImage string `json:"profile_image"`
	CreatedAt    string `json:"created_at"`
}
