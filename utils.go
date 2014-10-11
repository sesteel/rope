package rope

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	p = fmt.Printf
)

func (r *Rope) Equal(r2 *Rope) bool {
	if r == nil && r2 == nil {
		return true
	}
	if r == nil && r2 != nil || r != nil && r2 == nil {
		return false
	}
	if r.weight != r2.weight {
		return false
	}
	if !(bytes.Equal(r.content, r2.content)) {
		return false
	}
	if !r.left.Equal(r2.left) {
		return false
	}
	if !r.right.Equal(r2.right) {
		return false
	}
	return true
}

func (r *Rope) Dump() {
	r.dump(0)
}

func (r *Rope) dump(level int) {
	p("%s%d |%s|\n", strings.Repeat("  ", level), r.weight, r.content)
	if r.left != nil {
		r.left.dump(level + 1)
	}
	if r.right != nil {
		r.right.dump(level + 1)
	}
}