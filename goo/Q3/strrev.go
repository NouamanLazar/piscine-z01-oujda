package main

import (
	"fmt"
)

func StrRev(str string) string {
	// kan3tiw Qima l string tanya
	var str2 string
	// i heya range nta3 string bach tkteb 7arf b 7arf
	for i := range str {
		// str2 kan3tiwha hena Qima
		str2 = str2 + string(str[len(str)-i-1]) // "-i-i" heya li katkhli characters ytktbo tanqossiyan
	}
	return str2
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
