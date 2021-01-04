/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type Recur struct {
	arguments []SourceVariable
}

func (o *Recur) String() string {
	return fmt.Sprintf("[rcall (%v)]", o.arguments)
}
