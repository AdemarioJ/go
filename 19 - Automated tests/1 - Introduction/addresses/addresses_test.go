package addresses_test

import (
	. "test-introduction/addresses"
	"testing"
)

type testScenario struct {
	addressInsert string
	expectedType  string
}

func TestTypeOfAddress(t *testing.T) {
	t.Parallel()

	scenarioTest := []testScenario{
		{"rua das flores", "Rua"},
		{"avenida das flores", "Avenida"},
		{"estrada das flores", "Estrada"},
		{"rodovia das flores", "Rodovia"},
		{"Praça da Liberdade", "Invalid type"},
		{"RUA das FLORES", "Rua"},
		{"", "Invalid type"},
	}

	for _, scenario := range scenarioTest {
		result := TypeOfAddress(scenario.addressInsert)
		if result != scenario.expectedType {
			t.Errorf("The type received is not the expected. Expected: %s, Received: %s",
				scenario.expectedType,
				result)
		}
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("The test should not pass")
	}
}
