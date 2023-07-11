package main

import (
	"fmt"
	"net/url"
)

type urlParts struct {
	base              string
	pathParam         string
	airport           string
	zoomLevel         string
	departDate        string
	returnDate        string
	budget            string
	stopsFilterActive string
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
	if u.zoomLevel != "" {
		params.Add("zoomLevel", u.zoomLevel)
	}
	if u.departDate != "" {
		params.Add("depart", u.departDate)
	}
	if u.returnDate != "" {
		params.Add("return", u.returnDate)
	}
	if u.budget != "" {
		params.Add("budget", u.budget)
	}
	if u.stopsFilterActive != "" {
		params.Add("stopsFilterActive", u.stopsFilterActive)
	}
	base.RawQuery = params.Encode()
	fmt.Printf("URL is: %s\n", base.String())
	return base.String()
}
