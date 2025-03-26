package main

import "slices"

func Contains[T comparable](arr []T, item T) bool {
	return slices.Contains(arr, item)
}
