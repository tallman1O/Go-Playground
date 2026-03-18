package main

import "fmt"

// maps data type - a map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
// use maps to find things quickly.
// If you have a list of items and you want to check if a particular item is in the list, you can use a map to store the items as keys and then check for the presence of the key in the map.

type Vertex struct {
	Lat, Long float64
}

// map[type of key]type of value
var m map[string]Vertex // maps must be initialized before use. The make function returns a map of the given type, initialized and ready for use.

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// map literals
	var m1 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m1)

	// mutating maps
	m1["Bell Labs"] = Vertex{1, 2}
	fmt.Println(m1)

	delete(m1, "Bell Labs")
	fmt.Println(m1)

	// the optional second return value when retrieving a value from a map indicates whether the key was present in the map.
	// This can be used to disambiguate between a missing key and a key with a zero value.
	v, ok := m1["Bell Labs"]
	fmt.Println(v, ok) //{0 0} false

	v, ok = m1["Google"]
	fmt.Println(v, ok) //{37.42202 -122.08408} true
}
