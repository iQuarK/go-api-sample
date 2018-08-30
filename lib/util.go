package lib

import "reflect"

// ContainsPerson returns true if a Person belongs to an array
func ContainsPerson(haystack []Person, needle Person) bool {
	for _, item := range haystack {
		if reflect.DeepEqual(item, needle) {
			return true
		}
	}
	return false
}

// ContainsNumber returns true if an Integer belongs to an array
func ContainsNumber(haystack []uint, needle uint) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
