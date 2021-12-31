/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package assembler_sp

import (
	"fmt"

	swampdisasmsp "github.com/swamp/disassembler/lib"
	"github.com/swamp/opcodes/instruction_sp"
	"github.com/swamp/opcodes/opcode_sp"
	opcode_sp_type "github.com/swamp/opcodes/type"
)

type VariableType uint

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

func (c *Code) Label(identifier VariableName, debugString string) *Label {
	o := &Label{identifier: identifier, debugString: debugString}
	c.labels = append(c.labels, o)
	c.addStatement(o)
	return o
}

func (c *Code) VariableStart(name VariableName, posRange SourceStackPosRange) *VariableImpl {
	o := &VariableImpl{identifier: name, VariableNode: VariableNode{debugString: ""}}
//	c.labels = append(c.labels, o)
	c.addStatement(o)

	return o
}

func (c *Code) VariableEnd(refer *VariableImpl) *VariableEnd {
	o := &VariableEnd{refer:refer}
	//	c.labels = append(c.labels, o)
	c.addStatement(o)

	return o
}

func (c *Code) ListAppend(target TargetStackPos, a SourceStackPos, b SourceStackPos, position opcode_sp.FilePosition) {
	o := &ListAppend{target: target, a: a, b: b, position: position}
	c.addStatement(o)
}

func (c *Code) StringAppend(target TargetStackPos, a SourceStackPos, b SourceStackPos, position opcode_sp.FilePosition) {
	o := &StringAppend{target: target, a: a, b: b, position: position}
	c.addStatement(o)
}

func (c *Code) ListConj(target TargetStackPos, item SourceStackPos, itemSize StackItemSize, itemAlign opcode_sp_type.MemoryAlign, list SourceStackPos, position opcode_sp.FilePosition) {
	o := &ListConj{target: target, item: item, debugItemSize: itemSize, debugItemAlign: itemAlign, list: list, position: position}
	c.addStatement(o)
}

func (c *Code) IntBinaryOperator(target TargetStackPos, a SourceStackPos, b SourceStackPos, operator instruction_sp.BinaryOperatorType, position opcode_sp.FilePosition) {
	o := &IntBinaryOperator{target: target, a: a, b: b, operator: operator, position: position}
	c.addStatement(o)
}

func (c *Code) StringBinaryOperator(target TargetStackPos, a SourceStackPos, b SourceStackPos, operator instruction_sp.BinaryOperatorType, position opcode_sp.FilePosition) {
	o := &StringBinaryOperator{target: target, a: a, b: b, operator: operator, position: position}
	c.addStatement(o)
}

func (c *Code) EnumBinaryOperator(target TargetStackPos, a SourceStackPos, b SourceStackPos, operator instruction_sp.BinaryOperatorType, position opcode_sp.FilePosition) {
	o := &EnumBinaryOperator{target: target, a: a, b: b, operator: operator, position: position}
	c.addStatement(o)
}

func (c *Code) BooleanBinaryOperator(target TargetStackPos, a SourceStackPos, b SourceStackPos, operator instruction_sp.BinaryOperatorType, position opcode_sp.FilePosition) {
	o := &BooleanBinaryOperator{target: target, a: a, b: b, operator: operator, position: position}
	c.addStatement(o)
}

func (c *Code) UnaryOperator(target TargetStackPos, a SourceStackPos, operator instruction_sp.UnaryOperatorType, position opcode_sp.FilePosition) {
	o := &UnaryOperator{target: target, a: a, operator: operator, position: position}
	c.addStatement(o)
}

func (c *Code) ListLiteral(target TargetStackPos, values []SourceStackPos, itemSize StackRange, itemAlign opcode_sp_type.MemoryAlign, position opcode_sp.FilePosition) {
	o := &ListLiteral{target: target, values: values, itemSize: itemSize, itemAlign: itemAlign, position: position}
	c.addStatement(o)
}

func (c *Code) ArrayLiteral(target TargetStackPos, values []SourceStackPos, itemSize StackRange, itemAlign opcode_sp_type.MemoryAlign, position opcode_sp.FilePosition) {
	o := &ArrayLiteral{target: target, values: values, itemSize: itemSize, itemAlign: itemAlign, position: position}
	c.addStatement(o)
}

func (c *Code) CaseEnum(test SourceStackPos, consequences []*CaseConsequence, defaultConsequence *CaseConsequence, position opcode_sp.FilePosition) {
	o := &Case{test: test, consequences: consequences, defaultConsequence: defaultConsequence, position: position}
	c.addStatement(o)
}

func (c *Code) CasePatternMatchingInt(test SourceStackPos, consequences []*CaseConsequencePatternMatchingInt, defaultConsequence *Label, position opcode_sp.FilePosition) {
	o := &CasePatternMatchingInt{test: test, consequences: consequences, defaultConsequence: defaultConsequence, position: position}
	c.addStatement(o)
}

func (c *Code) CopyConstant(target TargetStackPos, source SourceDynamicMemoryPos, position opcode_sp.FilePosition) {
	o := &CopyConstant{target: target, source: source}
	c.addStatement(o)
}

func (c *Code) LoadInteger(target TargetStackPos, intValue int32, position opcode_sp.FilePosition) {
	o := &LoadInteger{target: target, intValue: intValue, position: position}
	c.addStatement(o)
}

func (c *Code) LoadRune(target TargetStackPos, runeValue instruction_sp.ShortRune, position opcode_sp.FilePosition) {
	o := &LoadRune{target: target, rune: runeValue, position: position}
	c.addStatement(o)
}

func (c *Code) LoadBool(target TargetStackPos, boolValue bool, position opcode_sp.FilePosition) {
	o := &LoadBool{target: target, boolean: boolValue, position: position}
	c.addStatement(o)
}

func (c *Code) SetEnum(target TargetStackPos, enumIndex uint8, position opcode_sp.FilePosition) {
	o := &SetEnum{target: target, enumIndex: enumIndex, position: position}
	c.addStatement(o)
}

func (c *Code) LoadZeroMemoryPointer(target TargetStackPos, zeroMemoryPointer SourceDynamicMemoryPos, position opcode_sp.FilePosition) {
	o := &LoadZeroMemoryPointer{target: target, sourceZeroMemory: zeroMemoryPointer, position: position}
	c.addStatement(o)
}

func (c *Code) CopyMemory(target TargetStackPos, source SourceStackPosRange, position opcode_sp.FilePosition) {
	o := NewCopyMemory(target, source, position)
	c.addStatement(o)
}

func (c *Code) BranchFalse(condition SourceStackPos, jump *Label, position opcode_sp.FilePosition) {
	if jump == nil {
		panic("swamp assembler: null jump")
	}
	o := &BranchFalse{condition: condition, jump: jump, position: position}
	c.addStatement(o)
}

func (c *Code) BranchTrue(condition SourceStackPos, jump *Label, position opcode_sp.FilePosition) {
	if jump == nil {
		panic("swamp assembler: null jump")
	}
	o := &BranchTrue{condition: condition, jump: jump, position: position}
	c.addStatement(o)
}

func (c *Code) Jump(jump *Label, position opcode_sp.FilePosition) {
	if jump == nil {
		panic("swamp assembler: null jump")
	}
	o := &Jump{jump: jump, position: position}
	c.addStatement(o)
}

func (c *Code) Return(position opcode_sp.FilePosition) {
	o := &Return{position: position}
	c.addStatement(o)
}

func (c *Code) Call(function SourceStackPos, newBasePointer TargetStackPos, position opcode_sp.FilePosition) {
	o := &Call{function: function, newBasePointer: newBasePointer, position: position}
	c.addStatement(o)
}

func (c *Code) Recur(position opcode_sp.FilePosition) {
	o := &Recur{position: position}
	c.addStatement(o)
}

func (c *Code) CallExternal(function SourceStackPos, newBasePointer TargetStackPos, position opcode_sp.FilePosition) {
	o := &CallExternal{function: function, newBasePointer: newBasePointer, position: position}
	c.addStatement(o)
}

func (c *Code) CallExternalWithSizes(function SourceStackPos, newBasePointer TargetStackPos, sizes []VariableArgumentPosSize, position opcode_sp.FilePosition) {
	o := &CallExternalWithSizes{function: function, newBasePointer: newBasePointer, sizes: sizes, position: position}
	c.addStatement(o)
}

func (c *Code) CallExternalWithSizesAndAlign(function SourceStackPos, newBasePointer TargetStackPos, sizes []VariableArgumentPosSizeAlign, position opcode_sp.FilePosition) {
	o := &CallExternalWithSizesAlign{function: function, newBasePointer: newBasePointer, sizes: sizes, position: position}
	c.addStatement(o)
}

func (c *Code) Curry(target TargetStackPos, typeIDConstant uint16, firstParameterAlign MemoryAlign, function SourceStackPos, startArgument SourceStackPosRange, position opcode_sp.FilePosition) {
	o := &Curry{target: target, typeIDConstant: typeIDConstant, firstParameterAlign: firstParameterAlign, function: function, arguments: startArgument, position: position}
	c.addStatement(o)
}

func targetStackPosition(pos TargetStackPos) opcode_sp_type.TargetStackPosition {
	return opcode_sp_type.TargetStackPosition(pos)
}

func sourceStackPosition(pos SourceStackPos) opcode_sp_type.SourceStackPosition {
	return opcode_sp_type.SourceStackPosition(pos)
}

func convertAlign(assemblerAlign MemoryAlign) opcode_sp_type.MemoryAlign {
	return opcode_sp_type.MemoryAlign(assemblerAlign)
}

func argOffsetSizes(args []VariableArgumentPosSize) []opcode_sp_type.ArgOffsetSize {
	offsetSizes := make([]opcode_sp_type.ArgOffsetSize, len(args))
	for index, arg := range args {
		offsetSizes[index] = opcode_sp_type.ArgOffsetSize{
			Offset: arg.Offset,
			Size:   arg.Size,
		}
	}
	return offsetSizes
}

func argOffsetSizesAlign(args []VariableArgumentPosSizeAlign) []opcode_sp_type.ArgOffsetSizeAlign {
	offsetSizes := make([]opcode_sp_type.ArgOffsetSizeAlign, len(args))
	for index, arg := range args {
		offsetSizes[index] = opcode_sp_type.ArgOffsetSizeAlign{
			Offset: arg.Offset,
			Size:   arg.Size,
			Align:  arg.Align,
		}
	}
	return offsetSizes
}

func sourceDynamicMemoryPos(pos SourceDynamicMemoryPos) opcode_sp_type.SourceDynamicMemoryPosition {
	return opcode_sp_type.SourceDynamicMemoryPosition(pos)
}

func sourceStackRange(size SourceStackRange) opcode_sp_type.SourceStackRange {
	return opcode_sp_type.SourceStackRange(size)
}

func stackRange(size StackRange) opcode_sp_type.StackRange {
	return opcode_sp_type.StackRange(size)
}

func stackRangeFromItemSize(size StackItemSize) opcode_sp_type.StackRange {
	return opcode_sp_type.StackRange(size)
}

func sourceStackPositionRange(pos SourceStackPosRange) opcode_sp_type.SourceStackPositionRange {
	return opcode_sp_type.SourceStackPositionRange{Position: sourceStackPosition(pos.Pos), Range: sourceStackRange(pos.Size)}
}

func writeUnaryOperator(stream *opcode_sp.Stream, operator *UnaryOperator) {
	stream.IntUnaryOperator(targetStackPosition(operator.target), operator.operator, sourceStackPosition(operator.a), operator.position)
}

func writeListAppend(stream *opcode_sp.Stream, operator *ListAppend) {
	stream.ListAppend(targetStackPosition(operator.target), sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeStringAppend(stream *opcode_sp.Stream, operator *StringAppend) {
	stream.StringAppend(targetStackPosition(operator.target), sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeListConj(stream *opcode_sp.Stream, operator *ListConj) {
	stream.ListConj(targetStackPosition(operator.target), sourceStackPosition(operator.list), sourceStackPosition(operator.item), stackRangeFromItemSize(operator.debugItemSize), operator.debugItemAlign, operator.position)
}

func writeBinaryOperator(stream *opcode_sp.Stream, operator *BinaryOperator) {
	stream.BinaryOperator(targetStackPosition(operator.target), operator.operator, sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeIntBinaryOperator(stream *opcode_sp.Stream, operator *IntBinaryOperator) {
	stream.IntBinaryOperator(targetStackPosition(operator.target), operator.operator, sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeStringBinaryOperator(stream *opcode_sp.Stream, operator *StringBinaryOperator) {
	stream.StringBinaryOperator(targetStackPosition(operator.target), operator.operator, sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeEnumBinaryOperator(stream *opcode_sp.Stream, operator *EnumBinaryOperator) {
	stream.EnumBinaryOperator(targetStackPosition(operator.target), operator.operator, sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeBooleanBinaryOperator(stream *opcode_sp.Stream, operator *BooleanBinaryOperator) {
	stream.BooleanBinaryOperator(targetStackPosition(operator.target), operator.operator, sourceStackPosition(operator.a), sourceStackPosition(operator.b), operator.position)
}

func writeCopyMemory(stream *opcode_sp.Stream, operator *CopyMemory) {
	stream.MemoryCopy(targetStackPosition(operator.target), sourceStackPositionRange(operator.source), operator.position)
}

func writeZeroMemoryPointer(stream *opcode_sp.Stream, operator *LoadZeroMemoryPointer) {
	stream.LoadZeroMemoryPointer(targetStackPosition(operator.target), sourceDynamicMemoryPos(operator.sourceZeroMemory), operator.position)
}

func writeBranchFalse(stream *opcode_sp.Stream, branch *BranchFalse) {
	stream.BranchFalse(sourceStackPosition(branch.Condition()), branch.Jump().OpLabel(), branch.position)
}

func writeBranchTrue(stream *opcode_sp.Stream, branch *BranchTrue) {
	stream.BranchTrue(sourceStackPosition(branch.Condition()), branch.Jump().OpLabel(), branch.position)
}

func writeJump(stream *opcode_sp.Stream, jump *Jump) {
	stream.Jump(jump.Jump().OpLabel(), jump.position)
}

func writeCase(stream *opcode_sp.Stream, caseExpr *Case) {
	var opLabels []instruction_sp.EnumCaseJump

	for _, consequence := range caseExpr.consequences {
		label := consequence.label.OpLabel()

		caseJump := instruction_sp.NewEnumCaseJump(consequence.InternalEnumIndex(), label)
		opLabels = append(opLabels, caseJump)
	}

	defaultCons := caseExpr.defaultConsequence

	if caseExpr.defaultConsequence != nil {
		label := defaultCons.label.OpLabel()
		caseJump := instruction_sp.NewEnumCaseJump(0xff, label)
		opLabels = append(opLabels, caseJump)
	}

	stream.EnumCase(sourceStackPosition(caseExpr.test), opLabels, caseExpr.position)
}

func writeCasePatternMatchingInt(stream *opcode_sp.Stream, caseExpr *CasePatternMatchingInt) {
	var opLabels []instruction_sp.EnumCasePatternMatchingIntJump

	for _, consequence := range caseExpr.consequences {
		label := consequence.label.OpLabel()

		caseJump := instruction_sp.NewEnumCasePatternMatchingIntJump(consequence.ConstantInteger(), label)
		opLabels = append(opLabels, caseJump)
	}

	defaultCons := caseExpr.defaultConsequence

	defaultLabel := defaultCons.OpLabel()

	stream.PatternMatchingInt(sourceStackPosition(caseExpr.test), opLabels, defaultLabel, caseExpr.position)
}

func writeList(stream *opcode_sp.Stream, listLiteral *ListLiteral) {
	var registers []opcode_sp_type.SourceStackPosition

	for _, argument := range listLiteral.values {
		registers = append(registers, sourceStackPosition(argument))
	}

	stream.CreateList(targetStackPosition(listLiteral.target), stackRange(listLiteral.itemSize), listLiteral.itemAlign, registers, listLiteral.position)
}

func writeLoadInteger(stream *opcode_sp.Stream, loadInteger *LoadInteger) {
	stream.LoadInteger(targetStackPosition(loadInteger.target), loadInteger.intValue, loadInteger.position)
}

func writeLoadRune(stream *opcode_sp.Stream, loadRune *LoadRune) {
	stream.LoadRune(targetStackPosition(loadRune.target), loadRune.rune, loadRune.position)
}

func writeLoadBool(stream *opcode_sp.Stream, loadBool *LoadBool) {
	stream.LoadBool(targetStackPosition(loadBool.target), loadBool.boolean, loadBool.position)
}

func writeSetEnum(stream *opcode_sp.Stream, setEnum *SetEnum) {
	stream.SetEnum(targetStackPosition(setEnum.target), setEnum.enumIndex, setEnum.position)
}

func writeCreateArray(stream *opcode_sp.Stream, arrayLiteral *ArrayLiteral) {
	var registers []opcode_sp_type.SourceStackPosition

	for _, argument := range arrayLiteral.values {
		registers = append(registers, sourceStackPosition(argument))
	}

	stream.CreateArray(targetStackPosition(arrayLiteral.target), stackRange(arrayLiteral.itemSize), arrayLiteral.itemAlign, registers, arrayLiteral.position)
}

func writeCallExternal(stream *opcode_sp.Stream, call *CallExternal) {
	stream.CallExternal(targetStackPosition(call.newBasePointer), sourceStackPosition(call.function), call.position)
}

func writeCallExternalWithSizes(stream *opcode_sp.Stream, call *CallExternalWithSizes) {
	stream.CallExternalWithSizes(targetStackPosition(call.newBasePointer), sourceStackPosition(call.function), argOffsetSizes(call.sizes), call.position)
}

func writeCallExternalWithSizesAndAlign(stream *opcode_sp.Stream, call *CallExternalWithSizesAlign) {
	stream.CallExternalWithSizesAndAlign(targetStackPosition(call.newBasePointer), sourceStackPosition(call.function), argOffsetSizesAlign(call.sizes), call.position)
}

func writeCall(stream *opcode_sp.Stream, call *Call) {
	stream.Call(targetStackPosition(call.newBasePointer), sourceStackPosition(call.function), call.position)
}

func writeRecur(stream *opcode_sp.Stream, call *Recur) {
	stream.TailCall(call.position)
}

func writeCurry(stream *opcode_sp.Stream, call *Curry) {
	stream.Curry(targetStackPosition(call.target), call.typeIDConstant, convertAlign(call.firstParameterAlign), sourceStackPosition(call.function), sourceStackPositionRange(call.arguments), call.position)
}

func writeReturn(stream *opcode_sp.Stream, ret *Return) {
	stream.Return(ret.position)
}

func handleStatement(cmd CodeCommand, opStream *opcode_sp.Stream) {
	switch t := cmd.(type) {
	case *Label:
		opStream.Label(t.OpLabel())
	case *VariableImpl:
	case *VariableEnd:
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
	case *CasePatternMatchingInt:
		writeCasePatternMatchingInt(opStream, t)
	case *ListLiteral:
		writeList(opStream, t)
	case *LoadInteger:
		writeLoadInteger(opStream, t)
	case *LoadBool:
		writeLoadBool(opStream, t)
	case *LoadRune:
		writeLoadRune(opStream, t)
	case *ArrayLiteral:
		writeCreateArray(opStream, t)
	case *ListAppend:
		writeListAppend(opStream, t)
	case *ListConj:
		writeListConj(opStream, t)
	case *StringAppend:
		writeStringAppend(opStream, t)
	case *Return:
		writeReturn(opStream, t)
	case *Call:
		writeCall(opStream, t)
	case *Recur:
		writeRecur(opStream, t)
	case *CallExternal:
		writeCallExternal(opStream, t)
	case *CallExternalWithSizes:
		writeCallExternalWithSizes(opStream, t)
	case *CallExternalWithSizesAlign:
		writeCallExternalWithSizesAndAlign(opStream, t)
	case *Curry:
		writeCurry(opStream, t)
	case *IntBinaryOperator:
		writeIntBinaryOperator(opStream, t)
	case *StringBinaryOperator:
		writeStringBinaryOperator(opStream, t)
	case *EnumBinaryOperator:
		writeEnumBinaryOperator(opStream, t)
	case *CopyMemory:
		writeCopyMemory(opStream, t)
	case *LoadZeroMemoryPointer:
		writeZeroMemoryPointer(opStream, t)
	case *SetEnum:
		writeSetEnum(opStream, t)
	default:
		panic(fmt.Sprintf("swamp assembler: unknown cmd %v", cmd))
	}
}

func (c *Code) Resolve(verboseFlag bool) ([]byte, []opcode_sp.OpcodeInfo, error) {
	if verboseFlag {
		// context.ShowSummary()
	}

	opStream := opcode_sp.NewStream()

	for _, label := range c.labels {
		opLabel := opStream.CreateLabel(label.Name())
		label.SetOpLabel(opLabel)
	}

	for _, cmd := range c.statements {
		handleStatement(cmd, opStream)
	}

	octets, opcodeInfos, err := opStream.Serialize()

	if verboseFlag {
		fmt.Println("--- disassembly ---")

		stringLines := swampdisasmsp.Disassemble(octets)
		for _, line := range stringLines {
			fmt.Printf("%s\n", line)
		}
	}

	return octets, opcodeInfos, err
}
