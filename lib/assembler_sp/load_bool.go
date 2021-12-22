package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type LoadBool struct {
	position opcode_sp.FilePosition
	target  TargetStackPos
	boolean bool
}

func (o *LoadBool) String() string {
	return fmt.Sprintf("[loadbool %v <= %v]", o.target, o.boolean)
}
