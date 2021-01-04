/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import (
	"fmt"
)

type ListAppend struct {
	target TargetVariable
	a      SourceVariable
	b      SourceVariable
}

func (o *ListAppend) String() string {
	return fmt.Sprintf("[listappend %v <= %v %v]", o.target, o.a, o.b)
}
