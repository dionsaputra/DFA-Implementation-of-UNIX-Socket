package dfa

type DFA struct {
	States      map[string]bool
	Start       string
	Symbols     map[string]bool
	Transitions map[TransitionInput]string
}

type TransitionInput struct {
	Source string
	Symbol string
}

func NewDFA() *DFA {
	return &DFA{
		States:      make(map[string]bool),
		Start:       "",
		Symbols:     make(map[string]bool),
		Transitions: make(map[TransitionInput]string),
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

func (dfa *DFA) AddTransition(symbol string, source string, destination string) *DFA {
	if _, ok := dfa.Symbols[symbol]; !ok {
		return dfa
	}
	if _, ok := dfa.States[source]; !ok {
		return dfa
	}
	if _, ok := dfa.States[destination]; !ok {
		return dfa
	}
	dfa.Transitions[TransitionInput{Source: source, Symbol: symbol}] = destination
	return dfa
}

func (dfa *DFA) Verify(inputs ...string) bool {
	lastState := dfa.Start
	for _, input := range inputs {
		transitionInput := TransitionInput{
			Source: lastState,
			Symbol: input,
		}
		if val, ok := dfa.Transitions[transitionInput]; ok {
			lastState = val
		} else {
			return false
		}
	}

	return dfa.States[lastState]
}
