package inputs

import (
	"strconv"
	"strings"

	"github.com/dowenliu-xyz/leetcode/internal/types"
)

func LevelOrderRestore(input string) (*types.TreeNode, error) {
	values, err := readValues(input)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, nil
	}
	if values[0] == nil {
		return nil, nil
	}
	root := &types.TreeNode{Val: *values[0]}
	q := []*types.TreeNode{root}
	values = values[1:]
	for len(values) > 0 && len(q) > 0 {
		var p []*types.TreeNode
		for _, node := range q {
			if len(values) == 0 {
				break
			}
			if values[0] == nil {
				node.Left = nil
			} else {
				node.Left = &types.TreeNode{Val: *values[0]}
				p = append(p, node.Left)
			}
			if len(values) == 1 {
				values = values[1:]
				break
			}
			if values[1] == nil {
				node.Right = nil
			} else {
				node.Right = &types.TreeNode{Val: *values[1]}
				p = append(p, node.Right)
			}
			values = values[2:]
		}
		q = p
	}
	return root, nil
}

func readValues(input string) ([]*int, error) {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	parts := strings.Split(input, ",")
	values := make([]*int, 0, len(parts))
	for _, s := range parts {
		trim := strings.TrimSpace(s)
		if trim == "null" {
			values = append(values, nil)
			continue
		}
		val64, err := strconv.ParseInt(trim, 10, 64)
		if err != nil {
			return nil, err
		}
		val := int(val64)
		values = append(values, &val)
	}
	return values, nil
}
