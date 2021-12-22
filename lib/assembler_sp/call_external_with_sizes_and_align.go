/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"
	"github.com/swamp/opcodes/opcode_sp"
)

type CallExternalWithSizesAlign struct {
	position opcode_sp.FilePosition
	function       SourceStackPos
	newBasePointer TargetStackPos
	sizes          []VariableArgumentPosSizeAlign
}

func (o *CallExternalWithSizesAlign) String() string {
	return fmt.Sprintf("[callExternalWithSizesAlign %v %v %v]", o.newBasePointer, o.function, o.sizes)
}
