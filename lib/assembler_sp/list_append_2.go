package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type ListAppend struct {
	position opcode_sp.FilePosition
	target TargetStackPos
	a      SourceStackPos
	b      SourceStackPos
}

func (o *ListAppend) String() string {
	return fmt.Sprintf("[listappend %v <= %v %v]", o.target, o.a, o.b)
}
