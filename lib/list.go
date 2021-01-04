/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type ListLiteral struct {
	target TargetVariable
	values []SourceVariable
}

func (o *ListLiteral) String() string {
	return fmt.Sprintf("[list %v <= %v]", o.target, o.values)
}
