/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import (
	"fmt"
)

type ListConj struct {
	target TargetVariable
	item   SourceVariable
	list   SourceVariable
}

func (o *ListConj) String() string {
	return fmt.Sprintf("[ListConj %v <= item:%v list:%v]", o.target, o.item, o.list)
}
