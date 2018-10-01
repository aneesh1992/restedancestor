// Package lib contains various functions that may be useful for more than one part of the API.
package lib

import (
	"encoding/json"
)

/*
//BadRequest returns a ready to use badrequest text, not used for now
func BadRequest(requestText string) []byte{
	badRequest := map[string]string{"code": "400", "message": requestText}
	badRequestJson, _ := json.MarshalIndent(badRequest, "", "")
	return badRequestJson
}
*/

// NotFound returns a ready to write message for ResponseWriter when needed a 404 error.
func NotFound(notFoundWord string) []byte {
	notFound := map[string]string{"code": "404", "message": "'" + notFoundWord + "' was not found in the database"}
	notFoundJSON, _ := json.MarshalIndent(notFound, "", "")

	return notFoundJSON
}

// InternalServerError returns a ready to write message for ResponseWriter when needed a 503 error.
func InternalServerError(uuid string) []byte {
	internalError := map[string]string{"code": "503", "message": "Unable to delete phrase with ID = '" + uuid + "'"}
	internalErrorJSON, _ := json.MarshalIndent(internalError, "", "")

	return internalErrorJSON
}

// ResultMsg returns a ready to write message for ResponseWriter when needed a message.
func ResultMsg(message string) []byte {
	msg := map[string]string{"code": "200", "message": "'" + message + "'"}
	msgJSON, _ := json.MarshalIndent(msg, "", "")

	return msgJSON
}

// Filter returns a slice with all the words of a QuoteSlice after been filtered.
/*func Filter(filters []string, slice database.QuoteSlice) []string {
	var count int
	filteredWords := make([]string, 0)

		for _, quote := range slice {
			s := strings.Split(quote.Quote, " ")
			for _, word := range s {
				filtered := word
				for count < len(filters) {
					filtered = strings.Split(filtered, filters[count])[0]
					count++
				}
				count = 0
				filteredWords = append(filteredWords, filtered)
			}
	}
	return filteredWords
}*/
