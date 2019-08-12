package main

type AvailableTile struct {
	Type    string   `json:type`
	Options []string `json:options`
	Death []string `json:death`
	Reproduce []string `json:reproduce`
}
