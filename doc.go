// Package tbotapi provides a Go wrapper for the Telegram Messenger Bot API.
//
// Note that, if the REST API returns an error, that error will be wrapped in a Go error and returned by the function call.
// This means that you will never have to examine the Ok value of responses, as the functions do that for you.
//
// An example bot is implemented in cmd/example.go, so check that out.
package tbotapi
