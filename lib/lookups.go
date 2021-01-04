/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type Lookups struct {
	target       TargetVariable
	a            SourceVariable
	indexLookups []uint8
}

func (o *Lookups) String() string {
	return fmt.Sprintf("[lookups %v <= %v %v]", o.target, o.a, o.indexLookups)
}
