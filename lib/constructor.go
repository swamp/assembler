/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type Constructor struct {
	target TargetVariable
	values []SourceVariable
}

func (o *Constructor) String() string {
	return fmt.Sprintf("[constructor %v <= %v]", o.target, o.values)
}
