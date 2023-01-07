# GO FILTER AND MAP

A map function applies a given function to each element of a collection, returning the results in a new collection.

A predicate is a single-argument function which returns a boolean value.

we'll apply the filter function on a slice of strings.
```go

package main

import (
    "fmt"
    "strings"
)

func filter(data []string, f func(string) bool) []string {

    fltd := make([]string, 0)

    for _, e := range data {

        if f(e) {
            fltd = append(fltd, e)
        }
    }

    return fltd
}

func main() {

    words := []string{"war", "water", "cup", "tree", "storm"}

    p := "w"

    res := filter(words, func(s string) bool {

        return strings.HasPrefix(s, p)
    })

    fmt.Println(res)

}

```

The filter function takes a slice of strings data and a function f as its arguments. It then creates a new empty slice fltd of strings.

The function then iterates over the elements in data and checks if f(v) is true for each element v. If f(v) is true, it appends v to the fltd slice.

Finally, the fltd slice is returned.

The main function then creates a slice of strings called words and a string called p. It calls the filter function and passes words and a function literal as arguments. The function literal takes a string s as an argument and returns true if s has the prefix p.

The result of the filter function is then printed using fmt.Println.

<br /> <br />
<b>function literal </b><br />
A function literal is an anonymous function that is defined inline. In this case, the function literal takes a string s as an argument and returns true if s has the prefix p.

The function literal is being passed as an argument to the filter function, and it will be used by the filter function to determine which elements from the words slice should be included in the resulting slice.

Here is the function literal again:
```go
func(s string) bool {
    return strings.HasPrefix(s, p)
}
```
#
<b>WE FILTER A SLICE OF STRUCT TYPES</b>
```go

package main

import "fmt"

type User struct {
    name       string
    occupation string
    country    string
}

func filter(data []User, f func(User) bool) []User {

    fltd := make([]User, 0)

    for _, user := range data {

        if f(user) {
            fltd = append(fltd, user)
        }
    }

    return fltd
}

func main() {

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

    res := filter(users, func(u User) bool {
        return u.country == country
    })

    fmt.Println(res)

}

```

The filter function takes a slice of User values and a function as its arguments, and returns a slice of User values. The function passed as an argument is a predicate function that takes a User value and returns a boolean value indicating whether or not the User value should be included in the output slice.

In the main function, a slice of User values is defined and assigned to the users variable. Then, a string value "Slovakia" is assigned to the country variable.

Next, the filter function is called with the users slice and a closure as its arguments. The closure is a function that takes a User value as its argument and returns true if the User value's country field is equal to the value of the country variable and false otherwise.

The filter function iterates over the users slice and calls the closure for each User value. If the closure returns true, the User value is appended to the output slice. When the loop completes, the output slice is returned by the filter function and printed to the console.