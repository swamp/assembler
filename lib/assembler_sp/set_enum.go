package assembler_sp

import (
	"fmt"

	"github.com/swamp/opcodes/opcode_sp"
)

type SetEnum struct {
	position  opcode_sp.FilePosition
	itemSize  StackRange
	target    TargetStackPos
	enumIndex uint8
}

func (o *SetEnum) String() string {
	return fmt.Sprintf("[setenum %v <= %v %v]", o.target, o.enumIndex, o.itemSize)
}
