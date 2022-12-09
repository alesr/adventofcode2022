package adventofcode22

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	testCases := []struct {
		name             string
		givenFile        string
		expectedCalories int
	}{
		{
			name:             "example",
			givenFile:        "day1_test_input.txt",
			expectedCalories: 24000,
		},
		{
			name:             "input",
			givenFile:        "day1_input.txt",
			expectedCalories: 70613,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			observed, err := part1(tc.givenFile)
			require.NoError(t, err)

			assert.Equal(t, tc.expectedCalories, observed)
		})
	}
}
