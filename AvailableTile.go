package main

type AvailableTile struct {
	Type    string   `json:grass`
	Options []string `json:options`
}
