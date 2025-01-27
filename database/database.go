package database

import (
	"database/sql"
	"fmt"
	"musclemate/structs"

	_ "github.com/lib/pq"
)

var databaseConnection *sql.DB

func Init() bool {
	connection, databaseConnectionError := sql.Open("postgres", "postgresql://testing:AmHxYcFu7Pdr@178.128.39.4/testing")

	if databaseConnectionError != nil {
		fmt.Println("Error connecting to database:", databaseConnectionError)
		return false
	} else {
		databaseConnection = connection
		return true
	}
}

func CreateTables() bool {
	_, err := databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating users table:", err)
		return false
	}

	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS workouts (
		workout_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT REFERENCES users(user_id) NOT NULL,
		name VARCHAR(100) NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating workouts table:", err)
		return false
	}

	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS exercises (
		exercise_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		workout_id BIGINT REFERENCES workouts(workout_id) NOT NULL,
		name VARCHAR(100) NOT NULL,
		reps INT,
		sets INT,
		weight DECIMAL(5,2),
		duration INT, -- duration in seconds
		video_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating exercises table:", err)
		return false
	}

	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS nutrition (
		nutrition_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT REFERENCES users(user_id) NOT NULL,
		meal_type VARCHAR(50), -- e.g., Breakfast, Lunch, Dinner, Snack
		food_name VARCHAR(100) NOT NULL,
		calories INT,
		protein DECIMAL(5,2),
		carbs DECIMAL(5,2),
		fats DECIMAL(5,2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating nutrition table:", err)
		return false
	}

	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS progress (
		progress_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT REFERENCES users(user_id) NOT NULL,
		weight DECIMAL(5,2),
		body_fat_percentage DECIMAL(5,2),
		chest DECIMAL(5,2),
		waist DECIMAL(5,2),
		hips DECIMAL(5,2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating progress table:", err)
		return false
	}

	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS goals (
		goal_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT REFERENCES users(user_id) NOT NULL,
		goal_type VARCHAR(100), -- e.g., Weight Loss, Muscle Gain
		target_value DECIMAL(5,2),
		start_date DATE,
		end_date DATE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating goals table:", err)
		return false
	}

	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS analytics (
		analytics_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT REFERENCES users(user_id) NOT NULL,
		date DATE NOT NULL,
		workout_duration INT, -- total workout duration in seconds
		calories_burned INT,
		calories_consumed INT,
		protein_consumed DECIMAL(5,2),
		carbs_consumed DECIMAL(5,2),
		fats_consumed DECIMAL(5,2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		fmt.Println("Error creating analytics table:", err)
		return false
	}

	return true
}
func GetWorkout(workoutID int64) structs.Workout {
	var workout structs.Workout

	err := databaseConnection.QueryRow("SELECT workout_id, user_id, name, description, created_at FROM workouts WHERE workout_id = $1", workoutID).Scan(&workout.WorkoutID, &workout.UserID, &workout.Name, &workout.Description, &workout.CreatedAt)

	if err != nil {
		fmt.Println("Error getting workout by workout id:", err)
		return workout
	}

	workout.Exercises = GetExercises(workoutID)

	return workout
}

func CreateWorkout(userID int64, name string, description string) int64 {
	var workoutID int64

	err := databaseConnection.QueryRow("INSERT INTO workouts (user_id, name, description) VALUES ($1, $2, $3) RETURNING workout_id", userID, name, description).Scan(&workoutID)

	if err != nil {
		fmt.Println("Error creating workout:", err)
		return 0
	}

	return workoutID
}

// add existing exercise to workout via workout id and exercise id

func AddExerciseToWorkout(workoutID int64, exerciseID int64) bool {
	_, err := databaseConnection.Exec("UPDATE exercises SET workout_id = $1 WHERE exercise_id = $2", workoutID, exerciseID)

	if err != nil {
		fmt.Println("Error adding exercise to workout:", err)
		return false
	}

	return true
}

func UpdateWorkout(workoutID int64, name string, description string) bool {
	_, err := databaseConnection.Exec("UPDATE workouts SET name = $1, description = $2 WHERE workout_id = $3", name, description, workoutID)

	if err != nil {
		fmt.Println("Error updating workout:", err)
		return false
	}

	return true
}

func DeleteWorkout(workoutID int64) bool {
	_, err := databaseConnection.Exec("DELETE FROM workouts WHERE workout_id = $1", workoutID)

	if err != nil {
		fmt.Println("Error deleting workout:", err)
		return false
	}

	return true
}

func GetWorkouts(userID int64) []structs.Workout {
	var workouts []structs.Workout

	rows, err := databaseConnection.Query("SELECT workout_id, user_id, name, description, created_at FROM workouts WHERE user_id = $1", userID)
	if err != nil {
		fmt.Println("Error getting workouts by user id:", err)
		return workouts
	}

	for rows.Next() {
		var workout structs.Workout
		err := rows.Scan(&workout.WorkoutID, &workout.UserID, &workout.Name, &workout.Description, &workout.CreatedAt)
		if err != nil {
			fmt.Println("Error scanning workout row:", err)
			return workouts
		}

		workout.Exercises = GetExercises(workout.WorkoutID)
		workouts = append(workouts, workout)
	}

	return workouts
}

func GetExercises(workoutID int64) []structs.Exercise {
	var exercises []structs.Exercise

	rows, err := databaseConnection.Query("SELECT exercise_id, workout_id, name, sets, reps, weight, duration, video_url, created_at FROM exercises WHERE workout_id = $1", workoutID)
	if err != nil {
		fmt.Println("Error getting exercises by workout id:", err)
		return exercises
	}

	for rows.Next() {
		var exercise structs.Exercise
		err := rows.Scan(&exercise.ExerciseID, &exercise.WorkoutID, &exercise.Name, &exercise.Sets, &exercise.Reps, &exercise.Weight, &exercise.Duration, &exercise.VideoURL, &exercise.CreatedAt)
		if err != nil {
			fmt.Println("Error scanning exercise row:", err)
			return exercises
		}

		exercises = append(exercises, exercise)
	}

	return exercises
}

func GetExercise(exerciseID int64) structs.Exercise {
	var exercise structs.Exercise

	err := databaseConnection.QueryRow("SELECT exercise_id, workout_id, name, sets, reps, weight, duration, video_url, created_at FROM exercises WHERE exercise_id = $1", exerciseID).Scan(&exercise.ExerciseID, &exercise.WorkoutID, &exercise.Name, &exercise.Sets, &exercise.Reps, &exercise.Weight, &exercise.Duration, &exercise.VideoURL, &exercise.CreatedAt)

	if err != nil {
		fmt.Println("Error getting exercise by exercise id:", err)
		return exercise
	}

	return exercise
}

func CreateExercise(workoutID int64, name string, sets int, reps int, weight float64, duration int, videoURL string) int64 {
	var exerciseID int64

	err := databaseConnection.QueryRow("INSERT INTO exercises (workout_id, name, sets, reps, weight, duration, video_url) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING exercise_id", workoutID, name, sets, reps, weight, duration, videoURL).Scan(&exerciseID)

	if err != nil {
		fmt.Println("Error creating exercise:", err)
		return 0
	}

	return exerciseID
}

func UpdateExercise(exerciseID int64, name string, sets int, reps int, weight float64, duration int, videoURL string) bool {
	_, err := databaseConnection.Exec("UPDATE exercises SET name = $1, sets = $2, reps = $3, weight = $4, duration = $5, video_url = $6 WHERE exercise_id = $7", name, sets, reps, weight, duration, videoURL, exerciseID)

	if err != nil {
		fmt.Println("Error updating exercise:", err)
		return false
	}

	return true
}

func DeleteExercise(exerciseID int64) bool {
	_, err := databaseConnection.Exec("DELETE FROM exercises WHERE exercise_id = $1", exerciseID)

	if err != nil {
		fmt.Println("Error deleting exercise:", err)
		return false
	}

	return true
}

func GetUser(userID int64) structs.User {

	var user structs.User

	err := databaseConnection.QueryRow("SELECT user_id, email, password, first_name, last_name, phone_number, profile_image, created_at FROM users WHERE user_id = $1", userID).Scan(&user.UserID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.ProfileImage, &user.CreatedAt)

	if err != nil {
		fmt.Println("Error getting user by user id:", err)
		return user
	}

	return user

}

func UserExists(userID int64) bool {
	var user structs.User

	err := databaseConnection.QueryRow("SELECT user_id FROM users WHERE user_id = $1", userID).Scan(&user.UserID)

	if err != nil {
		return false
	}

	return true
}

func CreateUser(email string, password string, firstName string, lastName string, phoneNumber string, profileImage string) int64 {
	var userID int64

	err := databaseConnection.QueryRow("INSERT INTO users (email, password, first_name, last_name, phone_number, profile_image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id", email, password, firstName, lastName, phoneNumber, profileImage).Scan(&userID)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return 0
	}

	return userID
}

func UpdateUser(userID int64, email string, password string, firstName string, lastName string, phoneNumber string, profileImage string) bool {
	_, err := databaseConnection.Exec("UPDATE users SET email = $1, password = $2, first_name = $3, last_name = $4, phone_number = $5, profile_image = $6 WHERE user_id = $7", email, password, firstName, lastName, phoneNumber, profileImage, userID)

	if err != nil {
		fmt.Println("Error updating user:", err)
		return false
	}

	return true
}

func DeleteUser(userID int64) bool {
	_, err := databaseConnection.Exec("DELETE FROM users WHERE user_id = $1", userID)

	if err != nil {
		fmt.Println("Error deleting user:", err)
		return false
	}

	return true
}
