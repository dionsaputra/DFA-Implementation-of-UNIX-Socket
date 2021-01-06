package dfa

type DFA struct {
	States  map[string]bool
	Start   string
	Symbols map[string]bool
}

func NewDFA() DFA {
	return DFA{
		States:  make(map[string]bool),
		Start:   "",
		Symbols: make(map[string]bool),
	}
}

func (dfa *DFA) AddState(state string, isFinal bool) *DFA {
	dfa.States[state] = isFinal
	return dfa
}

func (dfa *DFA) SetStart(state string) *DFA {
	if _, ok := dfa.States[state]; ok {
		dfa.Start = state
	}
	return dfa
}

func (dfa *DFA) SetSymbols(symbols ...string) *DFA {
	for _, symbol := range symbols {
		dfa.Symbols[symbol] = true
	}
	return dfa
}