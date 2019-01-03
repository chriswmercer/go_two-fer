//A package containining the solution for "Two Fer" for the Exercism GO track
package twofer

import (
	"fmt"
	"strings"
)

//Return a string after processing
func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 { name = "you" }
	return fmt.Sprintf("One for %s, one for me.", name)
}
