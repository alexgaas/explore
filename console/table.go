package console

import (
	cfg "explore/config"
	"explore/explore"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
)

/*
Exploration / exploitation statistics
*/
var (
	colTitleExploreMode  = "Mode"
	colTitleExploreCount = "Count"
	rowHeader            = table.Row{colTitleExploreMode, colTitleExploreCount}
)

func ShowModeStatTable(config *cfg.Config, exploitationMode int, explorationMode int) {
	t := table.NewWriter()

	t.AppendRow(table.Row{"Exploitation mode", exploitationMode})
	t.AppendRow(table.Row{"Exploration mode", explorationMode})

	t.AppendHeader(rowHeader)
	t.SetCaption(fmt.Sprintf("Number of iterations %d", config.Count))
	fmt.Println(t.Render())
}

//

/*
Show variants with collected statistics
*/
var (
	colTitleAllVariantsCompressionType = "Compression mode"
	colTitleAllVariantsScore           = "Score"
)

func ShowVariantsWithStatTable(config *cfg.Config, variants []explore.Variant) {
	var unit = ""
	if config.ScoreType == cfg.RUNTIME {
		unit = "runtime in µs"
	} else if config.ScoreType == cfg.COMPRESS_FACTOR {
		unit = "compression factor in bytes"
	}
	allVariantsTable := table.NewWriter()
	rowHeader := table.Row{colTitleAllVariantsCompressionType, colTitleAllVariantsScore + " (" + unit + ")"}
	allVariantsTable.AppendHeader(rowHeader)

	for _, variant := range variants {
		allVariantsTable.AppendRow(table.Row{variant.Name, variant.Score})
	}

	fmt.Println(allVariantsTable.Render())
}

//

/*
Show selected variants
*/
var (
	colTitleBestCompressionType     = "Selected (Best)\nCompression mode"
	colTitleBestScore               = "Selected (Best) Score"
	colTitleExploredCompressionType = "Explored\nCompression mode"
	colTitleExploredScore           = "Explored Score"
	selectedVariantsTable           = table.NewWriter()
)

func BuildSelectedVariantsTable(selectedVariant *explore.Variant, exploredVariant *explore.Variant) {
	selectedVariantsTable.AppendRow(table.Row{
		selectedVariant.Name, selectedVariant.Score, exploredVariant.Name, exploredVariant.Score})
}

func ShowSelectedVariantsTable(config *cfg.Config) {
	var unit = ""
	if config.ScoreType == cfg.RUNTIME {
		unit = "runtime in µs"
	} else if config.ScoreType == cfg.COMPRESS_FACTOR {
		unit = "compression factor in bytes"
	}
	rowHeader := table.Row{colTitleBestCompressionType, colTitleBestScore + "\n(" + unit + ")",
		colTitleExploredCompressionType, colTitleExploredScore + "\n(" + unit + ")"}
	selectedVariantsTable.AppendHeader(rowHeader)
	fmt.Println(selectedVariantsTable.Render())
}

//
