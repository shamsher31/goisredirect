// Package redirect Check if a number is a redirect HTTP status code
package redirect // import "github.com/shamsher31/goisredirect"

func get(n int) bool {
	code := []int{300, 301, 302, 303, 305, 307, 308}

	for _, val := range code {
		if n == val {
			return true
		}
	}
	return false
}

func Is(n int) bool {

	if get(n) {
		return true
	}
	return false
}
