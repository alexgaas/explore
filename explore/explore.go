package explore

import (
	"explore/config"
	"math"
)

const (
	ProbabilityDistributionBuckets = 256
	BestScore                      = int64(1e9)
)

var ExplorationProbabilityThreshold = 0.0
var ExploitationMode = 0
var ExplorationMode = 0

// simple calculation of threshold for Bayesian Bandit Algorithm
func calculateExplorationProbabilityThreshold(explorationProbabilityThreshold float64) {
	var explorationProbability = float64(1.0-explorationProbabilityThreshold/ProbabilityDistributionBuckets) / 1.5
	ExplorationProbabilityThreshold = math.Min(ProbabilityDistributionBuckets-1, ProbabilityDistributionBuckets*(1-explorationProbability))
}

func inExplorationMode(variantsWithStat []Variant, selectedVariant *Variant, config *config.Config) (*Variant, *Variant) {
	calculateExplorationProbabilityThreshold(ExplorationProbabilityThreshold)

	numCells, selectRandomVariant := HashVariants(ExplorationMode, variantsWithStat)

	_, selectedScore := selectRandomVariant.Fn(config)

	variantsWithStat[selectRandomVariant.Idx] = Variant{
		Idx:   selectRandomVariant.Idx,
		Name:  selectRandomVariant.Name,
		Score: selectedScore,
		Fn:    selectRandomVariant.Fn}

	bestVariant, bestScore := GetBestVariant(numCells, variantsWithStat)

	selectedVariant = &Variant{
		Idx:   variantsWithStat[bestVariant].Idx,
		Name:  variantsWithStat[bestVariant].Name,
		Score: bestScore,
		Fn:    variantsWithStat[bestVariant].Fn}

	ExplorationMode++
	return selectedVariant, &selectRandomVariant
}

func inExploitationMode(selectedVariant *Variant, variantsWithStat []Variant, config *config.Config) {
	if selectedVariant == nil {
		selectedVariant = &variantsWithStat[0]
	}

	_, _ = selectedVariant.Fn(config)

	ExploitationMode++
}

// ExploitOrExplore
// example of distribution:
/*
+-------------------+-------+
| MODE              | COUNT |
+-------------------+-------+
| Exploitation mode | 16181 |
| Exploration mode  |    75 |
+-------------------+-------+
For number of iterations 16256
*/
func ExploitOrExplore(currentCount int,
	selectedVariant *Variant,
	exploredVariant *Variant,
	variantsWithStat []Variant,
	config *config.Config) (*Variant, *Variant) {
	if float64(currentCount%ProbabilityDistributionBuckets) < ExplorationProbabilityThreshold {
		// Exploitation mode.
		inExploitationMode(selectedVariant, variantsWithStat, config)
	} else {
		// Exploration mode.
		selectedVariant, exploredVariant = inExplorationMode(variantsWithStat, selectedVariant, config)
	}
	return selectedVariant, exploredVariant
}
