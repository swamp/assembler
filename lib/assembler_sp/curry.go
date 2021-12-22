/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type Curry struct {
	position opcode_sp.FilePosition
	target              TargetStackPos
	typeIDConstant      uint16
	firstParameterAlign MemoryAlign
	function            SourceStackPos
	arguments           SourceStackPosRange
}

func (o *Curry) String() string {
	return fmt.Sprintf("[curry type:%v align:%v (%v) <= %v (%v)]", o.target, o.typeIDConstant, o.firstParameterAlign, o.function, o.arguments)
}
