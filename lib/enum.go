/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type Enum struct {
	target         TargetVariable
	enumFieldIndex int
	arguments      []SourceVariable
}

func (o *Enum) String() string {
	return fmt.Sprintf("[enum %v <= %v (%v)]", o.target, o.enumFieldIndex, o.arguments)
}
