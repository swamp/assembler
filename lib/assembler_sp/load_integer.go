package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type LoadInteger struct {
	position opcode_sp.FilePosition
	target   TargetStackPos
	intValue int32
}

func (o *LoadInteger) String() string {
	return fmt.Sprintf("[loadinteger %v <= %v]", o.target, o.intValue)
}
