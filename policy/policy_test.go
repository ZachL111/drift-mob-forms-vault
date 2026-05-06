package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 90, Capacity: 88, Latency: 13, Risk: 10, Weight: 9}, wantScore: 187, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 63, Capacity: 79, Latency: 11, Risk: 8, Weight: 7}, wantScore: 138, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 86, Capacity: 90, Latency: 23, Risk: 5, Weight: 12}, wantScore: 187, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
