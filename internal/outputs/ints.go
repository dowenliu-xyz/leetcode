package outputs

import (
	"bytes"
	"strconv"
)

func SprintIntSlice(result []int) string {
	buf := &bytes.Buffer{}
	buf.WriteString("[")
	for i, v := range result {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString("]")
	return buf.String()
}

func SprintIntMatrix(result [][]int) string {
	buf := &bytes.Buffer{}
	buf.WriteString("[")
	for i, v := range result {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(SprintIntSlice(v))
	}
	buf.WriteString("]")
	return buf.String()
}
