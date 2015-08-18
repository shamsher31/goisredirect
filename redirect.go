// Package redirect Check if a number is a redirect HTTP status code
package redirect // import "github.com/shamsher31/goisredirect"

func Is(s int) bool {
	code := map[int]int{
		300: 300,
		301: 301,
		302: 302,
		303: 303,
		305: 305,
		307: 307,
		308: 308,
	}

	if _, ok := code[s]; ok {
		return true
	}
	return false
}
