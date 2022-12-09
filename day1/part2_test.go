package adventofcode22

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	testCases := []struct {
		name             string
		givenFile        string
		expectedCalories int
	}{
		{
			name:             "example",
			givenFile:        "input_data_test",
			expectedCalories: 45000,
		},
		{
			name:             "input",
			givenFile:        "input_data",
			expectedCalories: 205805,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			observed, err := part2(tc.givenFile)
			require.NoError(t, err)

			assert.Equal(t, tc.expectedCalories, observed)
		})
	}
}
