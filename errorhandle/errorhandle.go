package errorhandle

import "log"

// Check for handling errors
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
