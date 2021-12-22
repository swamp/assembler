/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type Recur struct {
	position opcode_sp.FilePosition
}

func (o *Recur) String() string {
	return fmt.Sprintf("[rcall]")
}
