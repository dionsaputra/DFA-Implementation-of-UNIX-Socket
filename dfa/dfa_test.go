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
	assert.True(t, dfa.Transitions != nil && len(dfa.Transitions) == 0)
}

func TestDFA_AddState(t *testing.T) {
	t.Run("When state not exist then state added", func(t *testing.T) {
		dfa := NewDFA().AddState("state 1", true)

		assert.Equal(t, 1, len(dfa.States))
		val, ok := dfa.States["state 1"]
		assert.True(t, ok && val)
	})

	t.Run("When state exist then state updated", func(t *testing.T) {
		dfa := NewDFA().AddState("state 1", true).AddState("state 1", false)

		assert.Equal(t, 1, len(dfa.States))
		val, ok := dfa.States["state 1"]
		assert.True(t, ok && !val)
	})
}

func TestDFA_SetStart(t *testing.T) {
	t.Run("When start is not element of state then start not updated", func(t *testing.T) {
		dfa := NewDFA().AddState("state 1", true).SetStart("state 2")
		assert.Equal(t, "", dfa.Start)
	})

	t.Run("When start is element of state then start updated", func(t *testing.T) {
		dfa := NewDFA().AddState("state 1", true).SetStart("state 1")
		assert.Equal(t, "state 1", dfa.Start)
	})
}

func TestDFA_SetSymbols(t *testing.T) {
	dfa := NewDFA().SetSymbols("a", "b", "a", "b", "c", "d")
	assert.Equal(t, map[string]bool{"a": true, "b": true, "c": true, "d": true}, dfa.Symbols)
}

func TestDFA_AddTransition(t *testing.T) {
	t.Run("When symbol is not element of symbols then transitions not updated", func(t *testing.T) {
		dfa := NewDFA().AddState("s1", true).
			AddState("s2", true).
			SetSymbols("a", "b").
			AddTransition("c", "s1", "s2")
		assert.Equal(t, 0, len(dfa.Transitions))
	})

	t.Run("When source is not element of states then transitions not updated", func(t *testing.T) {
		dfa := NewDFA().AddState("s1", true).
			SetSymbols("a").
			AddTransition("a", "s2", "s1")
		assert.Equal(t, 0, len(dfa.Transitions))
	})

	t.Run("When destination is not element of states then transitions not updated", func(t *testing.T) {
		dfa := NewDFA().AddState("s1", true).
			SetSymbols("a").
			AddTransition("a", "s1", "s2")
		assert.Equal(t, 0, len(dfa.Transitions))
	})

	t.Run("When valid symbol, source, and destination then transitions updated", func(t *testing.T) {
		dfa := NewDFA().AddState("s1", false).
			AddState("s2", true).
			SetSymbols("a").
			AddTransition("a", "s1", "s2")

		assert.Equal(t, 1, len(dfa.Transitions))
		assert.Equal(t, map[TransitionInput]string{TransitionInput{Source: "s1", Symbol: "a"}: "s2"}, dfa.Transitions)
	})
}

func TestDFA_Verify(t *testing.T) {
	t.Run("When inputs terminate at final state return true", func(t *testing.T) {
		dfa := NewDFA().AddState("s1", false).
			AddState("s2", true).
			SetSymbols("a").
			SetStart("s1").
			AddTransition("a", "s1", "s2")

		assert.True(t, dfa.Verify("a"))
	})

	t.Run("When inputs terminate at not final state return false", func(t *testing.T) {
		dfa := NewDFA().AddState("s1", false).
			AddState("s2", false).
			SetSymbols("a").
			SetStart("s1").
			AddTransition("a", "s1", "s2")

		assert.False(t, dfa.Verify("a"))
	})
}
