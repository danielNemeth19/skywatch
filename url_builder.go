package main

import (
	"net/url"
)

type urlParts struct {
	base      string
	pathParam string
	airport   string
}

func (u urlParts) Compose() string {
	base, err := url.Parse(u.base)
	if err != nil {
		panic(err)
	}
	base.Path += u.pathParam
	params := url.Values{}
	if u.airport != "" {
		params.Add("airport", u.airport)
	}
	base.RawQuery = params.Encode()
	return base.String()
}
