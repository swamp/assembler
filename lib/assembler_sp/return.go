/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type Return struct {
	position opcode_sp.FilePosition
	//stackPointerAdd uint32
}

func (o *Return) String() string {
	return fmt.Sprintf("[ret]")
}
