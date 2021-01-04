/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type CallExternal struct {
	target    TargetVariable
	function  SourceVariable
	arguments []SourceVariable
}

func (o *CallExternal) String() string {
	return fmt.Sprintf("[CallExternal %v <= %v (%v)]", o.target, o.function, o.arguments)
}
