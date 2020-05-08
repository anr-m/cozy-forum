package errorhandle

import (
	"log"
)

// Fatal for handling errors
func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
