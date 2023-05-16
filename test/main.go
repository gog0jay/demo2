package main

import "fmt"

func main() {
	a := "a"
	var s []string
	s = append(s, "s1")
	s = append(s, "s2")
	s = append(s, "s3")
	s[1] = "sad"

	a = "b"
	fmt.Println(a)
	fmt.Println(s)
}
