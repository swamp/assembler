package assembler_sp

import (
	"fmt"

	"github.com/swamp/opcodes/instruction_sp"
	"github.com/swamp/opcodes/opcode_sp"
)

type LoadRune struct {
	position opcode_sp.FilePosition
	target TargetStackPos
	rune   instruction_sp.ShortRune
}

func (o *LoadRune) String() string {
	return fmt.Sprintf("[loadrune %v <= %v]", o.target, o.rune)
}
