package assembler_sp

import (
	"encoding/binary"
	"fmt"
	"log"
	"strings"

	"github.com/swamp/opcodes/opcode_sp"
	opcode_sp_type "github.com/swamp/opcodes/type"
)

type PackageConstants struct {
	constants             []*Constant
	functions             []*Constant
	externalFunctions     []*Constant
	strings               []*Constant
	resourceNames         []*Constant
	dynamicMapper         *DynamicMemoryMapper
	someConstantIDCounter uint
}

func NewPackageConstants() *PackageConstants {
	return &PackageConstants{
		dynamicMapper: DynamicMemoryMapperNew(256 * 1024),
	}
}

func (c *PackageConstants) String() string {
	s := "\n"
	for _, constant := range c.constants {
		if constant == nil {
			panic("swamp assembler: nil constant")
		}
		s += fmt.Sprintf("%v\n", constant)
	}
	return strings.TrimSpace(s)
}

func (c *PackageConstants) DebugString(filterTypes []ConstantType) string {
	s := "\n"
	for _, constant := range c.constants {
		if constant == nil {
			panic("swamp assembler: nil constant")
		}
		for _, filterTypeToInclude := range filterTypes {
			if constant.constantType == filterTypeToInclude {
				s += fmt.Sprintf("%v\n", constant)
				break;
			}
		}

	}
	return strings.TrimSpace(s)
}

func (c *PackageConstants) Constants() []*Constant {
	return c.constants
}

func (c *PackageConstants) Finalize() {
	if len(c.resourceNames) == 0 {
		return
	}
	pointerArea := c.dynamicMapper.Allocate(uint(int(opcode_sp_type.Sizeof64BitPointer)*len(c.resourceNames)), uint32(opcode_sp_type.Alignof64BitPointer), "Resource name chunk")
	for index, resourceName := range c.resourceNames {
		writePosition := pointerArea.Position + SourceDynamicMemoryPos(index*int(opcode_sp_type.Sizeof64BitPointer))
		binary.LittleEndian.PutUint64(c.dynamicMapper.memory[writePosition:writePosition+SourceDynamicMemoryPos(opcode_sp_type.Sizeof64BitPointer)], uint64(resourceName.PosRange().Position))
	}

	var resourceNameChunkOctets [16]byte
	binary.LittleEndian.PutUint64(resourceNameChunkOctets[0:8], uint64(pointerArea.Position))
	binary.LittleEndian.PutUint64(resourceNameChunkOctets[8:16], uint64(len(c.resourceNames)))
	resourceNameChunkPointer := c.dynamicMapper.WriteAlign(resourceNameChunkOctets[:], 8,"ResourceNameChunk struct (character-pointer-pointer, resourceNameCount)")

	c.constants = append(c.constants, &Constant{
		constantType:   ConstantTypeResourceNameChunk,
		str:            "",
		source:         resourceNameChunkPointer,
		debugString:    "resource name chunk",
		resourceNameId: 0,
	})
}

func (c *PackageConstants) DynamicMemory() *DynamicMemoryMapper {
	return c.dynamicMapper
}

func AllocateStringOctets(memoryMapper *DynamicMemoryMapper,s string) SourceDynamicMemoryPosRange {
	stringOctets := []byte(s)
	stringOctets = append(stringOctets, byte(0))
	stringOctetsPointer := memoryMapper.Write(stringOctets, "string:"+s)

	return stringOctetsPointer
}

const SizeofSwampString = 16

func (c *PackageConstants) AllocateStringConstant(s string) *Constant {
	for _, constant := range c.strings {
		if constant.str == s {
			return constant
		}
	}

	stringOctetsPointer := AllocateStringOctets(c.dynamicMapper, s)

	var swampStringOctets [SizeofSwampString]byte
	binary.LittleEndian.PutUint64(swampStringOctets[0:8], uint64(stringOctetsPointer.Position))
	binary.LittleEndian.PutUint64(swampStringOctets[8:16], uint64(len(s)))

	swampStringPointer := c.dynamicMapper.Write(swampStringOctets[:], "SwampString struct (character-pointer, characterCount) for:"+s)

	newConstant := NewStringConstant("string", s, swampStringPointer)
	c.constants = append(c.constants, newConstant)
	c.strings = append(c.strings, newConstant)

	return newConstant
}

func (c *PackageConstants) AllocateResourceNameConstant(s string) *Constant {
	for _, resourceNameConstant := range c.resourceNames {
		if resourceNameConstant.str == s {
			return resourceNameConstant
		}
	}

	stringOctetsPointer := AllocateStringOctets(c.dynamicMapper, s)

	newConstant := NewResourceNameConstant(c.someConstantIDCounter, s, stringOctetsPointer)
	c.someConstantIDCounter++
	c.constants = append(c.constants, newConstant)
	c.resourceNames = append(c.resourceNames, newConstant)

	return newConstant
}

const (
	SizeofSwampFunc         = 13 * 8
	SizeofSwampExternalFunc = 18 * 8
	SizeofSwampDebugInfoLines = 2 * 8
	SizeofSwampDebugInfoFiles = 2 * 8
	SizeofSwampDebugInfoScopes = 2 * 8
	SizeofSwampDebugInfoScopesEntry = 6 * 2 + 4 + 8
	AlignOfSwampDebugInfoScopesEntry = 8
)

func (c *PackageConstants) AllocateFunctionStruct(uniqueFullyQualifiedFunctionName string,
	opcodesPointer SourceDynamicMemoryPosRange, returnOctetSize opcode_sp_type.MemorySize,
	returnAlignSize opcode_sp_type.MemoryAlign, parameterCount uint, parameterOctetSize opcode_sp_type.MemorySize, typeIndex uint) (*Constant, error) {
	var swampFuncStruct [SizeofSwampFunc]byte

	fullyQualifiedStringPointer := AllocateStringOctets(c.dynamicMapper, uniqueFullyQualifiedFunctionName)

	binary.LittleEndian.PutUint32(swampFuncStruct[0:4], uint32(0))
	binary.LittleEndian.PutUint64(swampFuncStruct[8:16], uint64(parameterCount))      // parameterCount
	binary.LittleEndian.PutUint64(swampFuncStruct[16:24], uint64(parameterOctetSize)) // parameters octet size

	binary.LittleEndian.PutUint64(swampFuncStruct[24:32], uint64(opcodesPointer.Position))
	binary.LittleEndian.PutUint64(swampFuncStruct[32:40], uint64(opcodesPointer.Size))

	binary.LittleEndian.PutUint64(swampFuncStruct[40:48], uint64(returnOctetSize)) // returnOctetSize
	binary.LittleEndian.PutUint64(swampFuncStruct[48:56], uint64(returnAlignSize)) // returnAlign

	binary.LittleEndian.PutUint64(swampFuncStruct[56:64], uint64(fullyQualifiedStringPointer.Position)) // debugName
	binary.LittleEndian.PutUint64(swampFuncStruct[64:72], uint64(typeIndex))                            // typeIndex

	binary.LittleEndian.PutUint64(swampFuncStruct[72:80], uint64(opcodesPointer.Position))
	binary.LittleEndian.PutUint64(swampFuncStruct[80:88], uint64(opcodesPointer.Size))

	funcPointer := c.dynamicMapper.WriteAlign(swampFuncStruct[:], 8, "function Struct for:"+uniqueFullyQualifiedFunctionName)

	newConstant := NewFunctionReferenceConstantWithDebug("fn", uniqueFullyQualifiedFunctionName, funcPointer)
	c.constants = append(c.constants, newConstant)
	c.functions = append(c.functions, newConstant)

	return newConstant, nil
}

func (c *PackageConstants) AllocateExternalFunctionStruct(uniqueFullyQualifiedFunctionName string, returnValue SourceStackPosRange, parameters []SourceStackPosRange) (*Constant, error) {
	var swampFuncStruct [SizeofSwampExternalFunc]byte

	fullyQualifiedStringPointer := AllocateStringOctets(c.dynamicMapper, uniqueFullyQualifiedFunctionName)
	if len(parameters) == 0 {
		// panic(fmt.Errorf("not allowed to have zero paramters for %v", uniqueFullyQualifiedFunctionName))
	}

	binary.LittleEndian.PutUint32(swampFuncStruct[0:4], uint32(1))                  // external type
	binary.LittleEndian.PutUint64(swampFuncStruct[8:16], uint64(len(parameters)))   // parameterCount
	binary.LittleEndian.PutUint32(swampFuncStruct[16:20], uint32(returnValue.Pos))  // return pos
	binary.LittleEndian.PutUint32(swampFuncStruct[20:24], uint32(returnValue.Size)) // return size

	for index, param := range parameters {
		first := 24 + index*8
		firstEnd := first + 8
		second := 28 + index*8
		secondEnd := second + 8
		binary.LittleEndian.PutUint32(swampFuncStruct[first:firstEnd], uint32(param.Pos))    // params pos
		binary.LittleEndian.PutUint32(swampFuncStruct[second:secondEnd], uint32(param.Size)) // params size
	}

	binary.LittleEndian.PutUint64(swampFuncStruct[120:128], uint64(fullyQualifiedStringPointer.Position)) // debugName

	funcPointer := c.dynamicMapper.WriteAlign(swampFuncStruct[:], 8, fmt.Sprintf("external function Struct for: '%s' param Count: %d", uniqueFullyQualifiedFunctionName, len(parameters)))

	newConstant := NewExternalFunctionReferenceConstantWithDebug("fn", uniqueFullyQualifiedFunctionName, funcPointer)
	c.constants = append(c.constants, newConstant)
	c.externalFunctions = append(c.externalFunctions, newConstant)

	return newConstant, nil
}

func (c *PackageConstants) allocateDebugLinesStruct(count uint, debugLineOctets []byte, uniqueFullyQualifiedFunctionName string) SourceDynamicMemoryPosRange {
	var swampFuncStruct [SizeofSwampDebugInfoLines]byte

	debugLinesLinesPointer := c.dynamicMapper.WriteAlign(debugLineOctets, 2,"debug lines lines")

	binary.LittleEndian.PutUint32(swampFuncStruct[0:4], uint32(count))
	binary.LittleEndian.PutUint64(swampFuncStruct[8:16], uint64(debugLinesLinesPointer.Position))

	pointerToDebugLines := c.dynamicMapper.WriteAlign(swampFuncStruct[:], 4, "debug lines lines:"+uniqueFullyQualifiedFunctionName)

	return pointerToDebugLines
}

func (c *PackageConstants) AllocateDebugInfoFiles(fileUrls []*FileUrl) (*Constant, error) {
	var debugInfoFilesStruct [SizeofSwampDebugInfoFiles]byte

	stringPointers := make([]SourceDynamicMemoryPosRange, len(fileUrls))

	for index, fileUrl := range fileUrls {
		ptr := AllocateStringOctets(c.dynamicMapper, fileUrl.File)
		stringPointers[index] = ptr
	}

	spaceForArrayWithPointers := make([]byte, 8 * len(stringPointers))
	for index, stringPointer := range stringPointers {
		binary.LittleEndian.PutUint64(spaceForArrayWithPointers[index*8: index*8+8], uint64(stringPointer.Position))
	}
	arrayStart := c.dynamicMapper.WriteAlign(spaceForArrayWithPointers, 8, "array with pointers")

	binary.LittleEndian.PutUint32(debugInfoFilesStruct[0:4], uint32(len(fileUrls)))
	binary.LittleEndian.PutUint64(debugInfoFilesStruct[8:8+8], uint64(arrayStart.Position))

	debugInfoFilesPtr := c.dynamicMapper.WriteAlign(debugInfoFilesStruct[:], 8,"debug info files")

	newConstant := NewDebugInfoFilesWithDebug("debug info files", debugInfoFilesPtr)

	c.constants = append(c.constants, newConstant)

	return newConstant, nil
}

/*
func (c *PackageConstants) AllocateDebugInfoLines(instructions []*opcode_sp.Instruction) (*Constant, error) {
	var swampFuncStruct [SizeofSwampDebugInfoLines]byte

	binary.LittleEndian.PutUint32(swampFuncStruct[0:4], uint32(0))

	funcPointer := c.dynamicMapper.Write(swampFuncStruct[:], "function Struct for:")

	newConstant := NewFunctionReferenceConstantWithDebug("fn", "uniqueFullyQualifiedFunctionName", funcPointer)
	c.constants = append(c.constants, newConstant)

	return newConstant, nil
}
*/


const SwampFuncOpcodeOffset = 24
const SwampFuncDebugLinesOffset = 72
const SwampFuncDebugScopesOffset = 88

func (c *PackageConstants) FetchOpcodes(functionConstant *Constant) []byte {
	readSection := SourceDynamicMemoryPosRange{
		Position: SourceDynamicMemoryPos(uint(functionConstant.source.Position + SwampFuncOpcodeOffset)),
		Size:     DynamicMemoryRange(8 + 8),
	}
	opcodePointerAndSize := c.dynamicMapper.Read(readSection)
	opcodePosition := binary.LittleEndian.Uint64(opcodePointerAndSize[0:8])
	opcodeSize := binary.LittleEndian.Uint64(opcodePointerAndSize[8:16])

	readOpcodeSection := SourceDynamicMemoryPosRange{
		Position: SourceDynamicMemoryPos(opcodePosition),
		Size:     DynamicMemoryRange(opcodeSize),
	}

	return c.dynamicMapper.Read(readOpcodeSection)
}

func (c *PackageConstants) AllocatePrepareFunctionConstant(uniqueFullyQualifiedFunctionName string,
	returnSize opcode_sp_type.MemorySize, returnAlign opcode_sp_type.MemoryAlign,
	parameterCount uint, parameterOctetSize opcode_sp_type.MemorySize, typeId uint) (*Constant, error) {
	pointer := SourceDynamicMemoryPosRange{
		Position: 0,
		Size:     0,
	}

	return c.AllocateFunctionStruct(uniqueFullyQualifiedFunctionName, pointer, returnSize, returnAlign,
		parameterCount, parameterOctetSize, typeId)
}

func (c *PackageConstants) AllocatePrepareExternalFunctionConstant(uniqueFullyQualifiedFunctionName string, returnValue SourceStackPosRange, parameters []SourceStackPosRange) (*Constant, error) {
	return c.AllocateExternalFunctionStruct(uniqueFullyQualifiedFunctionName, returnValue, parameters)
}

func (c *PackageConstants) DefineFunctionOpcodes(funcConstant *Constant, opcodes []byte) error {
	opcodesPointer := c.dynamicMapper.Write(opcodes, "opcodes for:"+funcConstant.str)

	overwritePointer := SourceDynamicMemoryPos(uint(funcConstant.PosRange().Position) + SwampFuncOpcodeOffset)

	var opcodePointerOctets [16]byte

	binary.LittleEndian.PutUint64(opcodePointerOctets[0:8], uint64(opcodesPointer.Position))
	binary.LittleEndian.PutUint64(opcodePointerOctets[8:16], uint64(opcodesPointer.Size))

	c.dynamicMapper.Overwrite(overwritePointer, opcodePointerOctets[:], "opcodepointer"+funcConstant.str)

	return nil
}

func (c *PackageConstants) DefineFunctionDebugLines(funcConstant *Constant, count uint, debugInfoOctets []byte) error {
	overwritePointer := SourceDynamicMemoryPos(uint(funcConstant.PosRange().Position) + SwampFuncDebugLinesOffset)

	var debugLineOctets [16]byte

	debugLinesStructPointer := c.allocateDebugLinesStruct(count, debugInfoOctets, funcConstant.FunctionReferenceFullyQualifiedName())

	binary.LittleEndian.PutUint64(debugLineOctets[0:8], uint64(debugLinesStructPointer.Position))
	binary.LittleEndian.PutUint64(debugLineOctets[8:16], uint64(debugLinesStructPointer.Size))

	c.dynamicMapper.Overwrite(overwritePointer, debugLineOctets[:], "debugInfoOctets"+funcConstant.str)

	return nil
}


func serializeVariableInfoArray(memoryMapper *DynamicMemoryMapper, variables []opcode_sp.VariableInfo) SourceDynamicMemoryPosRange {
	var variableInfoEntry [SizeofSwampDebugInfoScopesEntry]byte
	if len(variables) == 0 {
		return SourceDynamicMemoryPosRange{
			SourceDynamicMemoryPosNull,
			0,
		}
	}
	startOfEntryArray := memoryMapper.Allocate(SizeofSwampDebugInfoScopesEntry * uint(len(variables)), AlignOfSwampDebugInfoScopesEntry, "variableInfos")
	for i, variable := range variables {
		nameStringPointer := AllocateStringOctets(memoryMapper, variable.Name)
		binary.LittleEndian.PutUint16(variableInfoEntry[0:2], uint16(variable.StartOpcodePosition))
		binary.LittleEndian.PutUint16(variableInfoEntry[2:4], uint16(variable.EndOpcodePosition))
		binary.LittleEndian.PutUint16(variableInfoEntry[4:6], uint16(variable.TypeID))
		binary.LittleEndian.PutUint16(variableInfoEntry[6:8], uint16(variable.ScopeID))
		binary.LittleEndian.PutUint16(variableInfoEntry[8:10], uint16(variable.StackPositionRange.Position))
		binary.LittleEndian.PutUint16(variableInfoEntry[10:12], uint16(variable.StackPositionRange.Range))
		binary.LittleEndian.PutUint64(variableInfoEntry[16:24], uint64(nameStringPointer.Position))
		overwritePosition := SourceDynamicMemoryPos(int(startOfEntryArray.Position) + SizeofSwampDebugInfoScopesEntry * i)
		memoryMapper.Overwrite(overwritePosition, variableInfoEntry[:], "variableInfo entry")
	}

	return startOfEntryArray
}


func (c *PackageConstants) allocateDebugScopesStruct(variableInfos []opcode_sp.VariableInfo, uniqueFullyQualifiedFunctionName string) SourceDynamicMemoryPosRange {
	var swampFuncStruct [SizeofSwampDebugInfoScopes]byte

	debugLinesLinesPointer := serializeVariableInfoArray(c.dynamicMapper, variableInfos)

	binary.LittleEndian.PutUint32(swampFuncStruct[0:4], uint32(len(variableInfos)))
	binary.LittleEndian.PutUint64(swampFuncStruct[8:16], uint64(debugLinesLinesPointer.Position))

	pointerToDebugScopes := c.dynamicMapper.WriteAlign(swampFuncStruct[:], 4, "debug scopes:"+uniqueFullyQualifiedFunctionName)

	return pointerToDebugScopes
}


func (c *PackageConstants) DefineFunctionDebugScopes(funcConstant *Constant, variables []opcode_sp.VariableInfo) error {
	overwritePointer := SourceDynamicMemoryPos(uint(funcConstant.PosRange().Position) + SwampFuncDebugScopesOffset)

	var debugInfoScopeStructOctets [16]byte

	debugInfoScopesPointerAndSize := c.allocateDebugScopesStruct(variables, funcConstant.FunctionReferenceFullyQualifiedName())

	binary.LittleEndian.PutUint64(debugInfoScopeStructOctets[0:8], uint64(debugInfoScopesPointerAndSize.Position))
	binary.LittleEndian.PutUint64(debugInfoScopeStructOctets[8:16], uint64(debugInfoScopesPointerAndSize.Size))

	c.dynamicMapper.Overwrite(overwritePointer, debugInfoScopeStructOctets[:], "debugInfoScopesOctets"+funcConstant.str)

	return nil
}

func (c *PackageConstants) FindFunction(identifier VariableName) *Constant {
	for _, constant := range c.functions {
		if constant.str == string(identifier) {
			return constant
		}
	}

	return c.FindExternalFunction(identifier)
}

func (c *PackageConstants) FindExternalFunction(identifier VariableName) *Constant {
	for _, constant := range c.externalFunctions {
		if constant.str == string(identifier) {
			return constant
		}
	}

	log.Printf("couldn't find constant external function %v", identifier)
	c.DebugOutput()

	return nil
}

func (c *PackageConstants) FindStringConstant(s string) *Constant {
	for _, constant := range c.strings {
		if constant.str == s {
			return constant
		}
	}
	return nil
}

func (c *PackageConstants) DebugOutput() {
	log.Printf("functions:\n")
	for _, function := range c.functions {
		log.Printf("%v %v\n", function.str, function.debugString)
	}
}
