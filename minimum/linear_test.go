package minimum

import "testing"

func TestLinearLossMinimum(t *testing.T) {
	// min := &Linear{
	// 	step:        0.2,
	// 	minimumStep: 0.015,
	// 	limit:       100,
	// }
	accuracy := 0.015
	min := Must(NewLinear(
		WithAccuracy(accuracy),
		WithMaximumSearchDepthOf(100),
	))

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			x2 := min.FindMinimum(0.1, testCase.Function)
			if !closeEnoughTo(accuracy, x2, testCase.ExpectedMinimum) {
				t.Fatal("expected minimum value was not found", x2, "!=", testCase.ExpectedMinimum)
			}
		})
	}

	// t.Fatal("dump")
}
