/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler

import (
	"fmt"

	swampdisasm "github.com/swamp/disassembler/lib"
	swampopcodeinst "github.com/swamp/opcodes/instruction"
	swampopcode "github.com/swamp/opcodes/opcode"
	swampopcodetype "github.com/swamp/opcodes/type"
)

type CodeCommand interface {
	String() string
}

type Code struct {
	statements []CodeCommand
	labels     []*Label
}

func (c *Code) Commands() []CodeCommand {
	return c.statements
}

func (c *Code) PrintOut() {
	for _, cmd := range c.statements {
		fmt.Printf("%v\n", cmd)
	}
}

func NewCode() *Code {
	return &Code{}
}

type CopyVariable struct {
	target TargetVariable
	source SourceVariable
}

func (o *CopyVariable) String() string {
	return fmt.Sprintf("[copyvar %v <= %v]", o.target, o.source)
}

func (c *Code) addStatement(cmd CodeCommand) {
	c.statements = append(c.statements, cmd)
}

func (c *Code) Copy(other *Code) {
	for _, cmd := range other.statements {
		lbl, _ := cmd.(*Label)
		if lbl != nil {
			c.labels = append(c.labels, lbl)
		}
		c.addStatement(cmd)
	}
}

func (c *Code) Label(identifier *VariableName, debugString string) *Label {
	o := &Label{identifier: identifier, debugString: debugString}
	c.labels = append(c.labels, o)
	c.addStatement(o)
	return o
}

func (c *Code) ListAppend(target TargetVariable, a SourceVariable, b SourceVariable) {
	o := &ListAppend{target: target, a: a, b: b}
	c.addStatement(o)
}

func (c *Code) StringAppend(target TargetVariable, a SourceVariable, b SourceVariable) {
	o := &StringAppend{target: target, a: a, b: b}
	c.addStatement(o)
}

func (c *Code) ListConj(target TargetVariable, item SourceVariable, list SourceVariable) {
	o := &ListConj{target: target, item: item, list: list}
	c.addStatement(o)
}

func (c *Code) BinaryOperator(target TargetVariable, a SourceVariable, b SourceVariable, operator swampopcodeinst.BinaryOperatorType) {
	o := &BinaryOperator{target: target, a: a, b: b, operator: operator}
	c.addStatement(o)
}

func (c *Code) UnaryOperator(target TargetVariable, a SourceVariable, operator swampopcodeinst.UnaryOperatorType) {
	o := &UnaryOperator{target: target, a: a, operator: operator}
	c.addStatement(o)
}

func (c *Code) Lookups(target TargetVariable, a SourceVariable, indexLookups []uint8) {
	o := &Lookups{target: target, a: a, indexLookups: indexLookups}
	c.addStatement(o)
}

func (c *Code) ListLiteral(target TargetVariable, values []SourceVariable) {
	o := &ListLiteral{target: target, values: values}
	c.addStatement(o)
}

func (c *Code) Constructor(target TargetVariable, values []SourceVariable) {
	o := &Constructor{target: target, values: values}
	c.addStatement(o)
}

func (c *Code) StructSplit(source SourceVariable, targets []TargetVariable) {
	o := &StructSplit{source: source, targets: targets}
	c.addStatement(o)
}

func (c *Code) UpdateStruct(target TargetVariable, structToCopy SourceVariable, updates []UpdateField) {
	o := &UpdateStruct{target: target, structToCopy: structToCopy, updates: updates}
	c.addStatement(o)
}

func (c *Code) Case(test SourceVariable, consequences []*CaseConsequence, defaultConsequence *CaseConsequence) {
	o := &Case{test: test, consequences: consequences, defaultConsequence: defaultConsequence}
	c.addStatement(o)
}

func (c *Code) CasePatternMatching(test SourceVariable, consequences []*CaseConsequencePatternMatching, defaultConsequence *CaseConsequencePatternMatching) {
	o := &CasePatternMatching{test: test, consequences: consequences, defaultConsequence: defaultConsequence}
	c.addStatement(o)
}

func (c *Code) BranchFalse(condition SourceVariable, jump *Label) {
	if jump == nil {
		panic("swamp assembler: null jump")
	}
	o := &BranchFalse{condition: condition, jump: jump}
	c.addStatement(o)
}

func (c *Code) BranchTrue(condition SourceVariable, jump *Label) {
	if jump == nil {
		panic("swamp assembler: null jump")
	}
	o := &BranchTrue{condition: condition, jump: jump}
	c.addStatement(o)
}

func (c *Code) Jump(jump *Label) {
	if jump == nil {
		panic("swamp assembler: null jump")
	}
	o := &Jump{jump: jump}
	c.addStatement(o)
}

func (c *Code) Return() {
	o := &Return{}
	c.addStatement(o)
}

func (c *Code) CopyVariable(target TargetVariable, source SourceVariable) {
	o := &CopyVariable{target: target, source: source}
	c.addStatement(o)
}

func (c *Code) Call(target TargetVariable, function SourceVariable, arguments []SourceVariable) {
	o := &Call{target: target, function: function, arguments: arguments}
	c.addStatement(o)
}
func (c *Code) Recur(arguments []SourceVariable) {
	o := &Recur{arguments: arguments}
	c.addStatement(o)
}

func (c *Code) CallExternal(target TargetVariable, function SourceVariable, arguments []SourceVariable) {
	o := &CallExternal{target: target, function: function, arguments: arguments}
	c.addStatement(o)
}

func (c *Code) CreateEnum(target TargetVariable, enumFieldIndex int, arguments []SourceVariable) {
	o := &Enum{target: target, enumFieldIndex: enumFieldIndex, arguments: arguments}
	c.addStatement(o)
}

func (c *Code) Curry(target TargetVariable, typeIDConstant uint16, function SourceVariable, arguments []SourceVariable) {
	o := &Curry{target: target, typeIDConstant: typeIDConstant, function: function, arguments: arguments}
	c.addStatement(o)
}

func writeUnaryOperator(stream *swampopcode.Stream, operator *UnaryOperator) {
	stream.IntUnaryOperator(operator.target.Register(), operator.operator, operator.a.Register())
}

func writeListAppend(stream *swampopcode.Stream, operator *ListAppend) {
	stream.ListAppend(operator.target.Register(), operator.a.Register(), operator.b.Register())
}

func writeStringAppend(stream *swampopcode.Stream, operator *StringAppend) {
	stream.StringAppend(operator.target.Register(), operator.a.Register(), operator.b.Register())
}

func writeListConj(stream *swampopcode.Stream, operator *ListConj) {
	stream.ListConj(operator.target.Register(), operator.list.Register(), operator.item.Register())
}

func writeBinaryOperator(stream *swampopcode.Stream, operator *BinaryOperator) {
	stream.BinaryOperator(operator.target.Register(), operator.operator, operator.a.Register(), operator.b.Register())
}

func writeBranchFalse(stream *swampopcode.Stream, branch *BranchFalse) {
	stream.BranchFalse(branch.Condition().Register(), branch.Jump().OpLabel())
}

func writeBranchTrue(stream *swampopcode.Stream, branch *BranchTrue) {
	stream.BranchTrue(branch.Condition().Register(), branch.Jump().OpLabel())
}

func writeJump(stream *swampopcode.Stream, jump *Jump) {
	stream.Jump(jump.Jump().OpLabel())
}

func writeCase(stream *swampopcode.Stream, caseExpr *Case) {
	var opLabels []swampopcodeinst.EnumCaseJump

	for _, consequence := range caseExpr.consequences {
		label := consequence.label.OpLabel()

		var arguments []swampopcodetype.Register

		for _, argument := range consequence.arguments {
			arguments = append(arguments, argument.Register())
		}

		caseJump := swampopcodeinst.NewEnumCaseJump(consequence.InternalEnumIndex(), arguments, label)
		opLabels = append(opLabels, caseJump)
	}

	defaultCons := caseExpr.defaultConsequence

	if caseExpr.defaultConsequence != nil {
		label := defaultCons.label.OpLabel()
		caseJump := swampopcodeinst.NewEnumCaseJump(0xff, nil, label)
		opLabels = append(opLabels, caseJump)
	}

	stream.EnumCase(caseExpr.test.Register(), opLabels)
}

func writeCasePatternMatching(stream *swampopcode.Stream, caseExpr *CasePatternMatching) {
	var opLabels []swampopcodeinst.CasePatternMatchingJump

	for _, consequence := range caseExpr.consequences {
		label := consequence.label.OpLabel()

		caseJump := swampopcodeinst.NewCasePatternMatchingJump(consequence.LiteralVariable().Register(), label)
		opLabels = append(opLabels, caseJump)
	}

	defaultCons := caseExpr.defaultConsequence

	if caseExpr.defaultConsequence != nil {
		label := defaultCons.label.OpLabel()
		caseJump := swampopcodeinst.NewCasePatternMatchingJump(swampopcodetype.Register{}, label)
		opLabels = append(opLabels, caseJump)
	}

	stream.CasePatternMatching(caseExpr.test.Register(), opLabels)
}

func writeConstructor(stream *swampopcode.Stream, constructor *Constructor) {
	var registers []swampopcodetype.Register

	for _, argument := range constructor.values {
		registers = append(registers, argument.Register())
	}

	stream.CreateStruct(constructor.target.Register(), registers)
}

func writeStructSplit(stream *swampopcode.Stream, constructor *StructSplit) {
	var targets []swampopcodetype.Register

	for _, argument := range constructor.targets {
		targets = append(targets, argument.Register())
	}

	stream.StructSplit(constructor.source.Register(), targets)
}

func writeUpdateStruct(stream *swampopcode.Stream, copyStruct *UpdateStruct) {
	var copyFields []swampopcodeinst.CopyToFieldInfo

	for _, update := range copyStruct.updates {
		copyField := swampopcodeinst.CopyToFieldInfo{Source: update.Source.Register(),
			Target: swampopcodetype.NewField(update.TargetField)}
		copyFields = append(copyFields, copyField)
	}

	stream.UpdateStruct(copyStruct.target.Register(), copyStruct.structToCopy.Register(), copyFields)
}

func writeList(stream *swampopcode.Stream, listLiteral *ListLiteral) {
	var registers []swampopcodetype.Register

	for _, argument := range listLiteral.values {
		registers = append(registers, argument.Register())
	}

	stream.CreateList(listLiteral.target.Register(), registers)
}

func writeCallExternal(stream *swampopcode.Stream, call *CallExternal) {
	var arguments []swampopcodetype.Register

	for _, argument := range call.arguments {
		arguments = append(arguments, argument.Register())
	}

	stream.CallExternal(call.target.Register(), call.function.Register(), arguments)
}

func writeCall(stream *swampopcode.Stream, call *Call) {
	var arguments []swampopcodetype.Register

	for _, argument := range call.arguments {
		arguments = append(arguments, argument.Register())
	}

	stream.Call(call.target.Register(), call.function.Register(), arguments)
}

func writeRecur(stream *swampopcode.Stream, call *Recur) {
	var arguments []swampopcodetype.Register

	for _, argument := range call.arguments {
		arguments = append(arguments, argument.Register())
	}

	stream.TailCall(arguments)
}

func writeCurry(stream *swampopcode.Stream, call *Curry) {
	var arguments []swampopcodetype.Register

	for _, argument := range call.arguments {
		arguments = append(arguments, argument.Register())
	}

	stream.Curry(call.target.Register(), call.typeIDConstant, call.function.Register(), arguments)
}

func writeEnum(stream *swampopcode.Stream, enumConstructor *Enum) {
	var arguments []swampopcodetype.Register

	for _, argument := range enumConstructor.arguments {
		arguments = append(arguments, argument.Register())
	}

	stream.Enum(enumConstructor.target.Register(), enumConstructor.enumFieldIndex, arguments)
}

func writeLookups(stream *swampopcode.Stream, lookups *Lookups) {
	var fields []swampopcodetype.Field

	for _, indexLookup := range lookups.indexLookups {
		fld := swampopcodetype.NewField(indexLookup)
		fields = append(fields, fld)
	}

	stream.GetStruct(lookups.target.Register(), lookups.a.Register(), fields)
}

func writeCopyVar(stream *swampopcode.Stream, copyVar *CopyVariable) {
	stream.RegCopy(copyVar.target.Register(), copyVar.source.Register())
}

func writeReturn(stream *swampopcode.Stream) {
	stream.Return()
}

func handleStatement(cmd CodeCommand, opStream *swampopcode.Stream) {
	switch t := cmd.(type) {
	case *BinaryOperator:
		writeBinaryOperator(opStream, t)
	case *UnaryOperator:
		writeUnaryOperator(opStream, t)
	case *BranchFalse:
		writeBranchFalse(opStream, t)
	case *BranchTrue:
		writeBranchTrue(opStream, t)
	case *Jump:
		writeJump(opStream, t)
	case *Case:
		writeCase(opStream, t)
	case *CasePatternMatching:
		writeCasePatternMatching(opStream, t)
	case *Constructor:
		writeConstructor(opStream, t)
	case *StructSplit:
		writeStructSplit(opStream, t)
	case *UpdateStruct:
		writeUpdateStruct(opStream, t)
	case *ListLiteral:
		writeList(opStream, t)
	case *ListAppend:
		writeListAppend(opStream, t)
	case *ListConj:
		writeListConj(opStream, t)
	case *StringAppend:
		writeStringAppend(opStream, t)
	case *Lookups:
		writeLookups(opStream, t)
	case *CopyVariable:
		writeCopyVar(opStream, t)
	case *Return:
		writeReturn(opStream)
	case *Label:
		opStream.Label(t.OpLabel())
	case *Call:
		writeCall(opStream, t)
	case *Recur:
		writeRecur(opStream, t)
	case *CallExternal:
		writeCallExternal(opStream, t)
	case *Curry:
		writeCurry(opStream, t)
	case *Enum:
		writeEnum(opStream, t)
	default:
		panic(fmt.Sprintf("swamp assembler: unknown cmd %v", cmd))
	}
}

func (c *Code) Resolve(context *FunctionRootContext, verboseFlag bool) ([]byte, error) {
	variableSpan := context.Layouter().HighestUsedRegisterValue()
	startConstantIndex := uint8(variableSpan + 1)

	for constantIndex, constant := range context.Constants().Constants() {
		r := swampopcodetype.NewRegister(uint8(startConstantIndex + uint8(constantIndex)))
		constant.SetRegister(r)
	}

	if verboseFlag {
		context.ShowSummary()
	}

	opStream := swampopcode.NewStream()

	for _, label := range c.labels {
		opLabel := opStream.CreateLabel(label.Name())
		label.SetOpLabel(opLabel)
	}

	for _, cmd := range c.statements {
		handleStatement(cmd, opStream)
	}

	octets, err := opStream.Serialize()

	if verboseFlag {
		fmt.Println("--- disassembly ---")

		stringLines := swampdisasm.Disassemble(octets)
		for _, line := range stringLines {
			fmt.Printf("%s\n", line)
		}
	}

	return octets, err
}
