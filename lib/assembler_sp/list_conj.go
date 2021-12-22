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

type ListConj struct {
	position opcode_sp.FilePosition
	target         TargetStackPos
	item           SourceStackPos
	list           SourceStackPos
	debugItemSize  StackItemSize
	debugItemAlign opcode_sp_type.MemoryAlign
}

func (o *ListConj) String() string {
	return fmt.Sprintf("[ListConj %v <= item:%v (%d, %d) list:%v]", o.target, o.item, o.debugItemSize, o.debugItemAlign, o.list)
}
