package explore

func HashVariants(explorationMode int, variantsWithStat []Variant) (int, Variant) {
	currentExplorationCount := explorationMode + 1
	hash := uint64(currentExplorationCount)
	// simple murmur finalizer
	hash *= 0xff51afd7ed558ccd
	hash ^= hash >> 33
	// get value by hash
	numCells := len(variantsWithStat)
	selectRandomVariant := variantsWithStat[hash%uint64(numCells)]
	return numCells, selectRandomVariant
}
