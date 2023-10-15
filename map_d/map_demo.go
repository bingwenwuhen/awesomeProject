package main

import "fmt"

func main() {
	var mapa = make(map[string]string, 5)
	mapa["1"] = ""
	mapa["2"] = "fdsgsd"
	for s := range mapa {
		nodeValue := mapa[s]
		if nodeValue == "" {
			fmt.Println("empty " + nodeValue)
		} else {
			fmt.Println("not null " + nodeValue)
		}
	}
}
