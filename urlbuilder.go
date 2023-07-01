package main

import (
	"fmt"
	"net/url"
)

type urlParts struct {
	base       string
	pathParam  string
	airport    string
	zoomLevel  string
	departDate string
	returnDate string
	budget     string
}

func (u urlParts) Compose() string {
	base, err := url.Parse(u.base)
	if err != nil {
		panic(err)
	}
	base.Path += u.pathParam
	params := url.Values{}
	params.Add("airport", u.airport)
	params.Add("zoomLevel", u.zoomLevel)
	params.Add("depart", u.departDate)
	params.Add("return", u.returnDate)
	params.Add("budget", u.budget)
	base.RawQuery = params.Encode()
	fmt.Printf("URL is: %s\n", base.String())
	return base.String()
}
