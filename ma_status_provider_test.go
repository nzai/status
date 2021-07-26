package status

import (
	"testing"
)

func TestMAStatusProvicedr_GetStatus(t *testing.T) {
	closes := [][]float32{{10.23, 9.17, 9, 8, 11, 8, 11.45}}

	provider := NewMAStatusProvider(3)
	for _, group := range closes {
		quotes := make([]*Quote, len(group))
		for index, _close := range group {
			quotes[index] = &Quote{Close: _close}
		}

		status, err := provider.GetStatus(quotes)
		if err != nil {
			t.Errorf("MAStatusProvicedr.GetStatus() error = %v", err)
		}

		t.Logf("status: %s", status)
	}
}
