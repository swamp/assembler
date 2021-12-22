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

type BinaryOperator struct {
	position opcode_sp.FilePosition
	target   TargetStackPos
	a        SourceStackPos
	b        SourceStackPos
	r        SourceStackRange
	operator instruction_sp.BinaryOperatorType
}

func (o *BinaryOperator) String() string {
	return fmt.Sprintf("[binop %v <= %v %v %v]", o.target, o.operator, o.a, o.b)
}
