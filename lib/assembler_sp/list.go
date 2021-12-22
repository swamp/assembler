/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"

	opcode_sp_type "github.com/swamp/opcodes/type"
	"github.com/swamp/opcodes/opcode_sp"
)

type ListLiteral struct {
	position opcode_sp.FilePosition
	target    TargetStackPos
	itemSize  StackRange
	itemAlign opcode_sp_type.MemoryAlign
	values    []SourceStackPos
}

func (o *ListLiteral) String() string {
	return fmt.Sprintf("[list %v (%d, %d) <= %v]", o.target, o.itemSize, o.itemAlign, o.values)
}
