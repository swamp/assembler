package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type LoadZeroMemoryPointer struct {
	position opcode_sp.FilePosition
	target           TargetStackPos
	sourceZeroMemory SourceDynamicMemoryPos
}

func (o *LoadZeroMemoryPointer) String() string {
	return fmt.Sprintf("[loadzeromem %v <= %v]", o.target, o.sourceZeroMemory)
}
