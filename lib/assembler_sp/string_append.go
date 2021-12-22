/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type StringAppend struct {
	position opcode_sp.FilePosition
	target TargetStackPos
	a      SourceStackPos
	b      SourceStackPos
}

func (o *StringAppend) String() string {
	return fmt.Sprintf("[stringappend %v <= %v %v]", o.target, o.a, o.b)
}
