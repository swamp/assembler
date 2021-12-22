/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"

	"github.com/swamp/opcodes/instruction_sp"
	"github.com/swamp/opcodes/opcode_sp"
)

type IntBinaryOperator struct {
	position opcode_sp.FilePosition
	target   TargetStackPos
	a        SourceStackPos
	b        SourceStackPos
	operator instruction_sp.BinaryOperatorType
}

func (o *IntBinaryOperator) String() string {
	return fmt.Sprintf("[ibinop %v <= %v %v %v]", o.target, o.operator, o.a, o.b)
}
