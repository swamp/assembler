/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import (
	"fmt"

	swampopcodeinst "github.com/swamp/opcodes/instruction"
)

type UnaryOperator struct {
	target   TargetVariable
	a        SourceVariable
	operator swampopcodeinst.UnaryOperatorType
}

func (o *UnaryOperator) String() string {
	return fmt.Sprintf("[unary %v <= %v %v]", o.target, o.operator, o.a)
}
