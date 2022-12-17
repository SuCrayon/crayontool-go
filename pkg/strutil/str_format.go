package strutil

import "fmt"

func SprintfRepeatedTimes(format string, v interface{}, repeatedTimes int) string {
	vs := make([]interface{}, 0, repeatedTimes)
	for i := 0; i < repeatedTimes; i++ {
		vs = append(vs, v)
	}
	return fmt.Sprintf(format, vs...)
}
