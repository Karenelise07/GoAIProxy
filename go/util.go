package swagger

import (
	"log"
	"net/http"
)




// logError is a helper function to log errors and send an HTTP error response
func logError(w http.ResponseWriter, message string, statusCode int) {
	log.Println(message)               // Log the error message
	http.Error(w, message, statusCode) // Send the HTTP error response
}
