/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import (
	"fmt"

	swampopcodetype "github.com/swamp/opcodes/type"
)

type Label struct {
	identifier  *VariableName
	debugString string
	opLabel     *swampopcodetype.Label
	offset      *swampopcodetype.Label
}

func (o *Label) String() string {
	if o.identifier != nil {
		return fmt.Sprintf("%v: # (%v)]", o.identifier, o.debugString)
	}
	return fmt.Sprintf("%v:", o.debugString)
}

func (o *Label) SetOpLabel(opLabel *swampopcodetype.Label) {
	o.opLabel = opLabel
}

func (o *Label) OpLabel() *swampopcodetype.Label {
	return o.opLabel
}

func (o *Label) OffsetLabel() *swampopcodetype.Label {
	return o.offset
}

func (o *Label) Name() string {
	if o.identifier != nil {
		return o.identifier.Name()
	}
	return o.debugString
}
