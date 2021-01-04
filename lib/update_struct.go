/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type UpdateField struct {
	TargetField uint8
	Source      SourceVariable
}

type UpdateStruct struct {
	target       TargetVariable
	structToCopy SourceVariable
	updates      []UpdateField
}

func (o *UpdateStruct) String() string {
	return fmt.Sprintf("[UpdateStruct %v <= (%v) %v]", o.target, o.structToCopy, o.updates)
}
