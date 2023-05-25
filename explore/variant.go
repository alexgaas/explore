package explore

import "explore/config"

type Variant struct {
	Idx   int
	Name  string
	Score int64
	Fn    func(config *config.Config) (int, int64)
}

func GetBestVariant(numCells int, variantsWithStat []Variant) (int, int64) {
	bestVariant := 0
	bestScore := BestScore
	for i := 0; i < numCells; i++ {
		score := variantsWithStat[i].Score
		if score < bestScore && score > 0 {
			bestScore = score
			bestVariant = i
		}
	}
	return bestVariant, bestScore
}
