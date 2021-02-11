/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import "fmt"

type CaseConsequencePatternMatching struct {
	caseLiteral SourceVariable
	label       *Label
}

func NewCaseConsequencePatternMatching(caseLiteral SourceVariable, label *Label) *CaseConsequencePatternMatching {
	return &CaseConsequencePatternMatching{caseLiteral: caseLiteral, label: label}
}

func (c *CaseConsequencePatternMatching) Label() *Label {
	return c.label
}

func (c *CaseConsequencePatternMatching) LiteralVariable() SourceVariable {
	return c.caseLiteral
}

func (c *CaseConsequencePatternMatching) String() string {
	return fmt.Sprintf("[caseconpm %v %v]", c.caseLiteral, c.label)
}

type CasePatternMatching struct {
	target             TargetVariable
	test               SourceVariable
	consequences       []*CaseConsequencePatternMatching
	defaultConsequence *CaseConsequencePatternMatching
}

func (o *CasePatternMatching) String() string {
	return fmt.Sprintf("[casepm  %v and then jump %v (%v)]", o.test, o.consequences, o.defaultConsequence)
}
