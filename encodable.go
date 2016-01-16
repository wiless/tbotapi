package tbotapi

type encodable interface {
	queryString() Querystring
}

type emptyEncodable struct{}

func (emptyEncodable) queryString() Querystring {
	return Querystring(map[string]string{})
}
