package domain

import (
	"fmt"
	"strconv"
	"strings"
)

// UssdAction just for now we hace 3 actions: doble_carga, recarga_tarjeta, transferir_credito
type UssdAction struct {
	UssdSteps []string `json:"ussdSteps"`
	Action    string   `json:"action"`
}

func (a *UssdAction) GetUSSD() string {

	if len(a.UssdSteps) == 0 {
		return ""
	}

	if len(a.UssdSteps) == 1 {
		return fmt.Sprint("*", a.UssdSteps[0], "#")
	}

	result := strings.Join(a.UssdSteps, "*")
	return fmt.Sprint("*", result, "#")
}

func (a *UssdAction) Replace(symbol string, num int) error {
	if !contains(a.UssdSteps, symbol) {
		return fmt.Errorf("%v symbol not found", symbol)
	}

	for i, v := range a.UssdSteps {
		if v == symbol {
			a.UssdSteps[i] = strconv.Itoa(num)
			return nil
		}
	}
	return nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}
