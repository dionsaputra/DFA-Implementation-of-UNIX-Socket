package dfa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDFA(t *testing.T) {
	dfa := NewDFA()

	assert.True(t, dfa.States != nil && len(dfa.States) == 0)
	assert.Empty(t, dfa.Start)
	assert.True(t, dfa.Symbols != nil && len(dfa.Symbols) == 0)
}

func TestDFA_AddState(t *testing.T) {
	t.Run("When state not exist then state added", func(t *testing.T) {
		dfa := NewDFA()
		dfa.AddState("state 1", true)

		assert.Equal(t, 1, len(dfa.States))
		val, ok := dfa.States["state 1"]
		assert.True(t, ok && val)
	})

	t.Run("When state exist then state updated", func(t *testing.T) {
		dfa := NewDFA()
		dfa.AddState("state 1", true)
		dfa.AddState("state 1", false)

		assert.Equal(t, 1, len(dfa.States))
		val, ok := dfa.States["state 1"]
		assert.True(t, ok && !val)
	})
}

func TestDFA_SetStart(t *testing.T) {
	t.Run("When start is not element of state then start not updated", func(t *testing.T) {
		dfa := NewDFA()
		dfa.AddState("state 1", true)
		dfa.SetStart("state 2")

		assert.Equal(t, "", dfa.Start)
	})

	t.Run("When start is element of state then start updated", func(t *testing.T) {
		dfa := NewDFA()
		dfa.AddState("state 1", true)
		dfa.SetStart("state 1")

		assert.Equal(t, "state 1", dfa.Start)
	})
}

func TestDFA_SetSymbols(t *testing.T) {
	dfa := NewDFA()
	dfa.SetSymbols("a", "b", "a", "b", "c", "d")

	assert.Equal(t, map[string]bool{"a": true, "b": true, "c": true, "d": true}, dfa.Symbols)
}
