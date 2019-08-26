package db

import (
	"fmt"
)

// sqllike returns the arg to be `%arg%`, so it can be used in SQL Query `LIKE %arg%`
func sqllike(arg string) string {
	return fmt.Sprintf("%%%s%%", arg)
}

// validatePage is used to check the limit and offset so it can be used in the query.
// It changes the value of limit and offset
func validatePage(limit, offset *int) {
	if *limit < 0 {
		*limit = DefaultLimit
	}

	// Offset < 0 will throw an error
	if *offset < 0 {
		*offset = 0
	}
}
