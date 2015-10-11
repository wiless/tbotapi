package tbotapi

import (
	"bitbucket.org/mrd0ll4r/tbotapi/model"
)

type encodable interface {
	GetQueryString() model.Querystring
}

type emptyEncodable struct{}

func (emptyEncodable) GetQueryString() model.Querystring {
	return model.Querystring(map[string]string{})
}
