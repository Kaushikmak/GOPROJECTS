package serialization

import (
	"fmt"
	"strings"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
)

// escapeString preserves JSON grammar by escaping backslashes and quotes
func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}

// Serialize converts a Task struct into a JSON string
func Serialize(T models.Task) string {
	// FIX: Added \" around %s for the id field
	return fmt.Sprintf(
		"{ \"id\": \"%s\", \"task\": \"%s\", \"status\": \"%s\", \"created_at\": \"%s\", \"updated_at\": \"%s\" }",
		T.Id.String(), // ID is a string, so it needs quotes in JSON
		escapeString(T.Description),
		models.StatusToString(T.Status),
		T.CreatedAt.Format(time.RFC3339),
		T.UpdatedAt.Format(time.RFC3339),
	)
}
