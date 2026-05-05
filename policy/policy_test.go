package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 90, Capacity: 88, Latency: 13, Risk: 10, Weight: 9}
	if got := Score(signal); got != 187 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 63, Capacity: 79, Latency: 11, Risk: 8, Weight: 7}
	if got := Score(signal); got != 138 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 86, Capacity: 90, Latency: 23, Risk: 5, Weight: 12}
	if got := Score(signal); got != 187 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
