package main

import (
	"fmt"
)

func filter1(data []string, f func(string) bool) []string {
	fltd := make([]string, 0)

	for _, v := range data {
		fmt.Println("calling F()")
		if f(v) {
			fltd = append(fltd, v)
		}

		fmt.Println("called F()")
		fmt.Println("")
		fmt.Println("")
	}

	return fltd
}

type User struct {
	name       string
	occupation string
	country    string
}

func filter2(data []User, f func(User) bool) []User {

	fltd := make([]User, 0)

	for _, user := range data {

		if f(user) {
			fltd = append(fltd, user)
		}
	}

	return fltd
}

func main() {

	// Filter 1 -------------------------------------------------

	// words := []string{"war", "water", "cup", "tree", "storm"}

	// p := "w"

	// res := filter1(words, func(s string) bool {

	// 	fmt.Println(s, p)

	// 	return strings.HasPrefix(s, p)
	// })

	// fmt.Println(res)

	// Filter 1 -------------------------------------------------

	// Filter 2 -------------------------------------------------
	users := []User{

		{"John Doe", "gardener", "USA"},
		{"Roger Roe", "driver", "UK"},
		{"Paul Smith", "programmer", "Canada"},
		{"Lucia Mala", "teacher", "Slovakia"},
		{"Patrick Connor", "shopkeeper", "USA"},
		{"Tim Welson", "programmer", "Canada"},
		{"Tomas Smutny", "programmer", "Slovakia"},
	}

	country := "Slovakia"

	res := filter2(users, func(u User) bool {
		return u.country == country
	})

	fmt.Println(res)

}
