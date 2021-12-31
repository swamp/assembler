/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"
)

type TypeString string
type TypeID uint32

type VariableName string

func (o VariableName) String() string {
	return fmt.Sprintf("[var %s]", string(o))
}

func (o VariableName) Name() string {
	return string(o)
}

type VariableNode struct {
	debugString string
	source      SourceStackPosRange
	startLabel* Label
	endLabel* Label
}

type VariableImpl struct {
	VariableNode
	identifier VariableName
	typeString TypeString
	typeID TypeID
}


func NewVariable(identifier VariableName, source SourceStackPosRange, typeID TypeID, typeString TypeString, startLabel* Label) *VariableImpl {
	return &VariableImpl{identifier: identifier, typeID: typeID, typeString: typeString, VariableNode: VariableNode{source: source, startLabel: startLabel}}
}

func (v *VariableImpl) String() string {
	return fmt.Sprintf("[var %v #%v]", v.identifier, v.VariableNode.source)
}

func (v *VariableImpl) EndLabel(endLabel *Label) {
	if v.endLabel != nil {
		panic("already defined")
	}

	v.endLabel = endLabel
}

func (v *VariableImpl) EndIsDefined() bool {
	return v.endLabel != nil
}

type VariableEnd struct {
	refer *VariableImpl
}

func NewVariableEnd(refer *VariableImpl) *VariableEnd {
	return &VariableEnd{refer: refer}
}

func (v *VariableEnd) String() string {
	return fmt.Sprintf("[varend %v]", v.refer)
}

