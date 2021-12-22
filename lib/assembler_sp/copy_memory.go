package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type CopyMemory struct {
	position opcode_sp.FilePosition
	target TargetStackPos
	source SourceStackPosRange
}

func NewCopyMemory(target TargetStackPos, source SourceStackPosRange, position opcode_sp.FilePosition) *CopyMemory {
	if source.Size == 0 {
		panic("not allowed copy zero size")
	}
	return &CopyMemory{
		position: position,
		target: target,
		source: source,
	}
}

func (o *CopyMemory) String() string {
	return fmt.Sprintf("[copymemory %v <= %v]", o.target, o.source)
}
