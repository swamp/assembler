/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type StructSplit struct {
	source  SourceVariable
	targets []TargetVariable
}

func (o *StructSplit) String() string {
	return fmt.Sprintf("[structsplit %v > %v]", o.source, o.targets)
}
