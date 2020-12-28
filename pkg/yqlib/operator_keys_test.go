package yqlib

import (
	"testing"
)

var keyOperatorScenarios = []expressionScenario{
	{
		description: "Get Keys of map",
		document:    `{a: {b: cat, c: dog, d: frog}}`,
		expression:  `.a | keys`,
		expected: []string{
			"D0, P[a], (!!seq)::- b\n- c\n- d\n",
		},
	},
	{
		skipDoc:    true,
		document:   `{a1: {b: cat, c: dog, d: frog}, a2: {e: cat, f: dog, g: frog}}`,
		expression: `(.a1, .a2) | keys`,
		expected: []string{
			"D0, P[a1], (!!seq)::- b\n- c\n- d\n",
			"D0, P[a2], (!!seq)::- e\n- f\n- g\n",
		},
	},
	{
		description: "Get Keys of array",
		document:    `{a: [0,1,2]}`,
		expression:  `.a | keys`,
		expected: []string{
			"D0, P[a], (!!seq)::0\n",
			"D0, P[a], (!!int)::1\n",
			"D0, P[a], (!!int)::2\n",
		},
	},
	{
		description: "Set key style",
		document:    `{a: {b: cat, c: dog, d: frog}}`,
		expression:  `(.a | keys | ..) style = 'single'`,
		expected: []string{
			"D0, P[], (!!doc)::{a: {'b': cat, 'c': dog, 'd': frog}}\n",
		},
	},
	{
		description: "Set key alias",
		document:    `{a: {b: cat, c: dog, d: frog}}`,
		expression:  `(.a | keys | .. | select(.=="b")) alias = 'boo'`,
		expected: []string{
			"D0, P[], (!!doc)::{a: {*meow: cat, c: dog, d: frog}}\n",
		},
	},
	{
		description: "Set key alias",
		document:    `{a: {b: cat, c: dog, d: frog}}`,
		expression:  `(.a | key("b")) alias = 'boo'`,
		expected: []string{
			"D0, P[], (!!doc)::{a: {*meow: cat, c: dog, d: frog}}\n",
		},
	},
	{
		skipDoc:    true,
		document:   `{a: [0,1,2]}`,
		expression: `(.a | keys) style = 'single'`,
		expected: []string{
			"D0, P[a], (!!int)::0\n",
			"D0, P[a], (!!int)::1\n",
			"D0, P[a], (!!int)::2\n",
		},
	},
	{
		skipDoc:    true,
		document:   `a: cat`,
		expression: `.a | keys`,
		expected:   []string{},
	},
	{
		skipDoc:    true,
		document:   `{}`,
		expression: `.a | keys`,
		expected:   []string{},
	},
}

func TestKeyOperatorScenarios(t *testing.T) {
	for _, tt := range keyOperatorScenarios {
		testScenario(t, &tt)
	}
	documentScenarios(t, "Keys", keyOperatorScenarios)
}
