// Code generated by "stringer -type Op"; DO NOT EDIT.

package edit

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Del - -1]
	_ = x[Eq-0]
	_ = x[Ins-1]
}

const _Op_name = "DelEqIns"

var _Op_index = [...]uint8{0, 3, 5, 8}

func (i Op) String() string {
	i -= -1
	if i < 0 || i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
