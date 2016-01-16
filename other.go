package tbotapi

// Querystring is a type to represent querystring-applicable data
type querystring map[string]string

type querystringer interface {
	querystring() querystring
}

type file struct {
	fieldName string
	path      string
}
