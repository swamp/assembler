/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type Curry struct {
	target    TargetVariable
	function  SourceVariable
	arguments []SourceVariable
}

func (o *Curry) String() string {
	return fmt.Sprintf("[curry %v <= %v (%v)]", o.target, o.function, o.arguments)
}
