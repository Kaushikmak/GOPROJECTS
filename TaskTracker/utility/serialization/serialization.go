package serialization

import (
	"fmt"
	"strings"
	"time"

	"github.com/kaushikmak/go-projects/TaskTracker/models"
)

// preserve grammer fo JSON
func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}

// serialize the TASK
func Serialize(T models.Task) string {
	return fmt.Sprintf(
		"{ \"id\": %v, \"task\": \"%s\", \"status\": \"%s\", \"created at\": \"%s\", \"updated at\": \"%s\" }",
		T.Id.String(),
		escapeString(T.Description),
		models.StatusToString(T.Status),
		T.CreatedAt.Format(time.RFC3339),
		T.UpdatedAt.Format(time.RFC3339),
	)
}
