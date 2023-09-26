package helper

import (
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func Generator() string {
	return uuid.New().String()
}

func UnescapeUnicode(input string) string {
	re := regexp.MustCompile(`\\u([a-zA-Z0-9]{4})`)
	return re.ReplaceAllStringFunc(input, unescape)
}

func unescape(uStr string) string {
	code, _ := strconv.ParseInt(strings.TrimPrefix(uStr, "\\u"), 16, 32)
	return string(rune(code))
}

func GenerateRoomID() string {
	str1 := randomString(4)
	time.Sleep(time.Millisecond)
	str2 := randomString(4)
	return str1 + "-" + str2
}

func randomString(length int) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	characters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	result := make([]rune, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}

func TruncateString(str string, length int) string {
	if utf8.RuneCountInString(str) <= length {
		return str
	}

	truncated := []rune(str)[:length]
	if len(truncated) > 0 {
		// Check if the last character is a partial UTF-8 character
		if truncated[len(truncated)-1] != utf8.RuneError {
			// If it's not a partial character, we can return the truncated string as is
			return string(truncated) + "..."
		}
	}

	// If the last character is a partial UTF-8 character, remove it and return the truncated string
	return string(truncated[:len(truncated)-1]) + "..."
}
