package security

import (
	"math/rand"
	"musclemate/utils/base64"
	"strconv"
	"time"
)

func GenerateAuthentication(username string) string {
	return base64.Encode(username) + "." + base64.Encode(strconv.Itoa(int(Since_Epoch()))) + "." + base64.Encode(RString(25))
}

func RString(length int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	end_string := make([]rune, length)

	for i := range end_string {
		end_string[i] = letters[rand.Intn(len(letters))]
	}

	return string(end_string)
}

func Since_Epoch() int64 {
	epoch_time := time.Date(2023, time.October, 8, 0, 0, 0, 0, time.UTC)
	time_now := time.Now()

	return time_now.Sub(epoch_time).Milliseconds()
}
