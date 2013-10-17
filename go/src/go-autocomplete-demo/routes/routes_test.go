package routes

import (
	"testing"
)

func Test_Routes(t *testing.T) {
	nameCounter := make(map[string]int)

	for _, route := range Routes {
		nameCounter[route.Name]++

		if nameCounter[route.Name] > 1 {
			t.Fatalf("duplicate route name: %q", route.Name)
		}
	}
}
