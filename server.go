package main

import (
	"fmt"
	"musclemate/database"
	api "musclemate/endpoints"
	"musclemate/utils/logger"
	"net/http"
	"os"
	"strconv"
)

var enableRateLimiting bool
var rateLimitThreshold int

func init() {
	enableRateLimiting = os.Getenv("ENABLE_RATE_LIMITING") == "true"
	threshold, err := strconv.Atoi(os.Getenv("RATE_LIMIT_THRESHOLD"))
	if err != nil {
		// default threshold if environment variable not set or invalid
		rateLimitThreshold = 100
	} else {
		rateLimitThreshold = threshold
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, "Request received at", r.URL.Path)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if enableRateLimiting {
			// ipAddress := r.Header.Get("X-Real-Ip")
			// if ipAddress == "" {
			// 	ipAddress = r.Header.Get("X-Forwarded-For")
			// }
			// if ipAddress == "" {
			// 	ipAddress = r.RemoteAddr
			// }
			// rateLimit, err := database.CheckRateLimit(ipAddress)
			// if err != nil {
			// 	logger.LogError(err)
			// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			// 	return
			// }
			// if rateLimit == nil {
			// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			// 	return
			// }
			// rateLimitVal := *rateLimit
			// if rateLimitVal.RequestCount >= rateLimitThreshold {
			// 	http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			// 	return
			// }
			// rateLimitVal.RequestCount++
			// rateLimitVal.LastRequest = time.Now()
			// if err := database.UpdateRateLimit(rateLimitVal); err != nil {
			// 	logger.LogError(err)
			// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			// 	return
			// }
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	if database.Init() && database.CreateTables() {
		logger.LogSuccess("[Database] Connected successfully!")

		http.HandleFunc("/api/accounts", rateLimitMiddleware(api.HandleAccounts))
		http.HandleFunc("/api/workouts", rateLimitMiddleware(api.HandleWorkouts))
		http.HandleFunc("/api/exercises", rateLimitMiddleware(api.HandleExercises))
	} else {
		logger.LogWarning("[Database] Failed to connect!")
		return
	}
	logger.LogSuccess("[Server] Started server successfully!")
	// logger.LogSuccess("[Twilio] Connected!")
	// logger.LogSuccess("[SendGrid] Connected!")
	fmt.Println(http.ListenAndServe(":6969", corsMiddleware(http.DefaultServeMux)))
}
