// Package tbotapi provides a Go wrapper for the Telegram Messenger Bot API.
//
// Note that, if the REST API returns an error, that error will be wrapped in a Go error and returned by the function call.
// This means that you will never have to examine the Ok value of responses, as the functions do that for you.
//
// We currently only support long polling (i.e. no webhooks). Feature-wise, everything up to and including the January 4
// changes should be implemented.
//
// An example bot is implemented in cmd/example.go, so check that out.
//
// The Bot API imposes certain limitations, these are especially interesting for inline query results and files. This
// library does not keep track of those limitations, so you'll have to perform checks yourself.
package tbotapi
