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
	// Create users table first
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

	// Create workouts table
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

	// Create exercises table
	_, err = databaseConnection.Exec(`
	CREATE TABLE IF NOT EXISTS exercises (
		exercise_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		workout_id BIGINT REFERENCES workouts(workout_id) NOT NULL,
		name VARCHAR(100) NOT NULL,
		sets INT,
		reps INT,
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

	// Create nutrition table
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

	// Create progress table
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

	// Create goals table
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

	// Create analytics table
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
