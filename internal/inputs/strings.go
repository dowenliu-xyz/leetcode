package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

func ReadStringSlice(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "[") || !strings.HasSuffix(input, "]") {
		return nil, fmt.Errorf("input should be wrapped by '[' and ']'")
	}
	input = input[1 : len(input)-1] // remove '[' and ']'
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	quotedStrings := strings.Split(input, ",")
	strs := make([]string, 0, len(quotedStrings))
	for i := range quotedStrings {
		str, err := strconv.Unquote(strings.TrimSpace(quotedStrings[i]))
		if err != nil {
			return nil, err
		}
		strs = append(strs, str)
	}
	return strs, nil
}
