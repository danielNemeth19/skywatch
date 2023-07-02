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
	params.Add("airport", u.airport)
	params.Add("zoomLevel", u.zoomLevel)
	params.Add("depart", u.departDate)
	params.Add("return", u.returnDate)
	params.Add("budget", u.budget)
	params.Add("stopsFilterActive", u.stopsFilterActive)
	params.Add("duration", "")

	params.Add("topRightLat", "68.58212830775821")
	params.Add("topRightLon", "180")
	params.Add("bottomLeftLat", "-6.168763628541718")
	params.Add("bottomLeftLon", "-180")
	base.RawQuery = params.Encode()
	fmt.Printf("URL is: %s\n", base.String())
	return base.String()
}
