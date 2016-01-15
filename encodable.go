package tbotapi

type encodable interface {
	GetQueryString() Querystring
}

type emptyEncodable struct{}

func (emptyEncodable) GetQueryString() Querystring {
	return Querystring(map[string]string{})
}
