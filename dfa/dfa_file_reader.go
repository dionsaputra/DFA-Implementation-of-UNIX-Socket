package dfa

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonDFA struct {
	Start  string `json:"start"`
	States []struct {
		Name    string `json:"name"`
		IsFinal bool   `json:"is_final"`
	} `json:"states"`
	Symbols     []string `json:"symbols"`
	Transitions []struct {
		Symbol      string `json:"symbol"`
		Source      string `json:"source"`
		Destination string `json:"destination"`
	} `json:"transitions"`
}

func ReadDFA(fileName string) (*DFA, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var jsonDFA JsonDFA
	if err := json.Unmarshal(byteValue, &jsonDFA); err != nil {
		return nil, err
	}

	dfa := NewDFA().SetSymbols(jsonDFA.Symbols...)

	for _, state := range jsonDFA.States {
		dfa.AddState(state.Name, state.IsFinal)
	}
	dfa.SetStart(jsonDFA.Start)

	for _, transition := range jsonDFA.Transitions {
		dfa.AddTransition(transition.Symbol, transition.Source, transition.Destination)
	}
	return dfa, nil
}
