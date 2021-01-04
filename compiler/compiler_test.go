/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package asmcompile_test

import (
	"testing"
)

func TestCallExternal(t *testing.T) {
	CompileTester(t, "callexternal 0 list_map 2 3 4", `
[constant1 funcExternal:list_map #1]

00: callexternal 0 1 ([2 3 4])
	`)
}
