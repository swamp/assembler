/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type CaseConsequence struct {
	caseValue uint8
	arguments []SourceVariable
	label     *Label
}

func NewCaseConsequence(caseValue uint8, arguments []SourceVariable, label *Label) *CaseConsequence {
	return &CaseConsequence{caseValue: caseValue, arguments: arguments, label: label}
}

func (c *CaseConsequence) Label() *Label {
	return c.label
}

func (c *CaseConsequence) InternalEnumIndex() uint8 {
	return c.caseValue
}

func (c *CaseConsequence) String() string {
	return fmt.Sprintf("[casecon %v %v %v]", c.caseValue, c.arguments, c.label)
}

type Case struct {
	test               SourceVariable
	consequences       []*CaseConsequence
	defaultConsequence *CaseConsequence
}

func (o *Case) String() string {
	return fmt.Sprintf("[case %v and then jump %v (%v)]", o.test, o.consequences, o.defaultConsequence)
}
