package main

import "strconv"

type Address struct {
	Protocol string
	Hostname string
	Port     int
}

func (a *Address) createPostString() string {
	return a.Protocol + "://" + a.Hostname + ":" + strconv.Itoa(a.Port)
}
