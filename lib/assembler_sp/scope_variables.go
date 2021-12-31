package assembler_sp

import (
	"fmt"
	"log"
	"strings"

	"github.com/swamp/opcodes/opcode_sp"
)

type ScopeVariables struct {
	parent         *ScopeVariables
	nameToVariable map[string]*VariableImpl
	childScopes []*ScopeVariables
	debugString string
}

func NewFunctionVariables(debugString string) *ScopeVariables {
	return &ScopeVariables{nameToVariable: make(map[string]*VariableImpl), debugString: debugString}
}

func NewFunctionVariablesWithParent(parent *ScopeVariables, debugString string) *ScopeVariables {
	newScope := &ScopeVariables{nameToVariable: make(map[string]*VariableImpl), parent: parent, debugString: debugString}

	parent.addChildScope(newScope)

	return newScope
}

func labelToOpcodePosition(label* Label) opcode_sp.OpcodePosition {
	if label == nil {
		panic("how is this possible")
	}
	if label.OpLabel() == nil {
		panic(fmt.Errorf("label is not defined %v", label))
	}
	return opcode_sp.OpcodePosition(label.OpLabel().DefinedProgramCounter().Value())
}

func VariableInfosDebugOutput(variables []opcode_sp.VariableInfo ) {
	log.Printf("count: %d", len(variables))
	for _, variable := range variables {
		log.Printf("variable:%v", variable)
	}
}

func GenerateVariablesWithScope(c *ScopeVariables, scopeID uint) []opcode_sp.VariableInfo {
	var variableInfos []opcode_sp.VariableInfo

	for _, variable := range c.nameToVariable {
		if !variable.EndIsDefined() {
			panic(fmt.Errorf("variable is not end defined %v", variable))
		}
		variableInfo := opcode_sp.VariableInfo{
			StartOpcodePosition: labelToOpcodePosition(variable.startLabel),
			EndOpcodePosition:   labelToOpcodePosition(variable.endLabel),
			ScopeID:             scopeID,
			TypeID:              uint32(variable.typeID),
			Name: variable.identifier.Name(),
		}
		variableInfos = append(variableInfos, variableInfo)
	}

	for _, scope := range c.childScopes {
		if len(scope.nameToVariable) > 0 {
			variableInfos = append(variableInfos, GenerateVariablesWithScope(scope, scopeID+1)...)
		}
	}

	return variableInfos
}

func (c *ScopeVariables) addChildScope(child *ScopeVariables) {
	c.childScopes = append(c.childScopes, child)
}


func indent(indentCount uint) string {
	return strings.Repeat("..", int(indentCount))
}

func (c *ScopeVariables) DebugOutput(indentCount uint) {
	for _, variable := range c.nameToVariable {

		log.Printf("%v%v %v = %v // %v", indent(indentCount), variable.identifier, variable.typeString, variable.source, variable.debugString)
	}

	for _, scope := range c.childScopes {
		if len(scope.nameToVariable) > 0 {
			log.Printf("%v %v (child scope)",  indent(indentCount), scope.debugString)
			scope.DebugOutput(indentCount+1)
		}
	}
}

func (c *ScopeVariables) DefineVariable(name VariableName, posRange SourceStackPosRange, typeID TypeID, typeString TypeString, label *Label) (*VariableImpl, error) {
	if uint(posRange.Size) == 0 {
		return nil, fmt.Errorf("octet size zero is not allowed for allocate stack memory")
	}

	stringName := string(name)
	_, alreadyHas := c.nameToVariable[stringName]
	if alreadyHas {
		return nil, fmt.Errorf("cannot define variable again '%s'", name)
	}

	v := NewVariable(name, posRange, typeID, typeString, label)

	c.nameToVariable[stringName] = v

	return v, nil
}

func (c *ScopeVariables) StopScope(label* Label) error {
	//name := variable.identifier.String()
	//delete(c.nameToVariable, name)

	for _, variable := range c.nameToVariable {
		variable.EndLabel(label)
	}

	return nil
}

func (c *ScopeVariables) FindVariable(name string) (SourceStackPosRange, error) {
	existingVariable, alreadyHas := c.nameToVariable[name]
	if !alreadyHas {
		if c.parent != nil {
			return c.parent.FindVariable(name)
		}
		return SourceStackPosRange{}, fmt.Errorf("could not find variable %v", name)
	}

	return existingVariable.source, nil
}
