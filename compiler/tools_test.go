/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package asmcompile_test

import (
	"fmt"
	"strings"
	"testing"

	asmcompile "github.com/swamp/assembler/compiler"
	swampdisasm "github.com/swamp/disassembler/lib"
)

func internalCompile(asm string) (string, error) {
	const verbose = true

	octets, context, compileErr := asmcompile.Compile(asm, verbose)
	if compileErr != nil {
		return "", compileErr
	}
	contextOutput := context.String()
	lines := swampdisasm.Disassemble(octets)

	for _, line := range lines {
		fmt.Printf("line: %v\n", line)
	}

	output := strings.Join(lines, "\n")
	return contextOutput + "\n\n" + output, nil
}

func CompileTester(t *testing.T, asm string, expected string) {
	output, err := internalCompile(asm)
	if err != nil {
		t.Fatal(err)
	}
	output = strings.TrimSpace(output)

	expected = strings.TrimSpace(expected)
	if output != expected {
		t.Errorf("expected\n%v\nbut received\n%v\n", expected, output)
	}
}
