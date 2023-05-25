package main

import (
	"bufio"
	compression "explore/compress"
	cfg "explore/config"
	"explore/console"
	"explore/explore"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	appConfig, err := cfg.GetNewConfig(cfg.GetAppFlags())
	if err != nil {
		log.Fatalln(err)
	}

	currentCount := 0
	counts := appConfig.Count

	// setup compression functions to explore
	variantsWithStat := setupExplorationFunctions(appConfig)

	var exploredVariant *explore.Variant
	var selectedVariant *explore.Variant

	pw := console.InitProgress()
	console.TrackProgress(pw, appConfig)

	for currentCount < counts {
		variant, exploredVariant := explore.ExploitOrExplore(currentCount, selectedVariant, exploredVariant, variantsWithStat, appConfig)

		// show stat for selected variants
		if variant != selectedVariant {
			selectedVariant = variant
			console.BuildSelectedVariantsTable(selectedVariant, exploredVariant)
		}

		currentCount++
	}

	console.StopProgress(pw)

	console.ShowModeStatTable(appConfig, explore.ExploitationMode, explore.ExplorationMode)
	console.ShowVariantsWithStatTable(appConfig, variantsWithStat)
	console.ShowSelectedVariantsTable(appConfig)
}

func setupExplorationFunctions(config *cfg.Config) []explore.Variant {
	var variantsWithStat []explore.Variant

	for idx, codec := range config.Codecs {
		switch codec {
		case "snappy":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDSnappy.String(),
				Score: explore.BestScore,
				Fn:    addSnappyWithStat,
			})
		case "lz4":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDLz4.String(),
				Score: explore.BestScore,
				Fn:    addLz4WithStat,
			})
		case "lz4_high_compression":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDLz4HighCompression.String(),
				Score: explore.BestScore,
				Fn:    addLz4HighCompressionWithStat,
			})
		case "brotli_1":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli1.String(),
				Score: explore.BestScore,
				Fn:    addBrotli1WithStat,
			})
		case "brotli_2":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli2.String(),
				Score: explore.BestScore,
				Fn:    addBrotli2WithStat,
			})
		case "brotli_3":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli3.String(),
				Score: explore.BestScore,
				Fn:    addBrotli3WithStat,
			})
		case "brotli_4":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli4.String(),
				Score: explore.BestScore,
				Fn:    addBrotli4WithStat,
			})
		case "brotli_5":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli5.String(),
				Score: explore.BestScore,
				Fn:    addBrotli5WithStat,
			})
		case "brotli_6":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli6.String(),
				Score: explore.BestScore,
				Fn:    addBrotli6WithStat,
			})
		case "brotli_7":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli7.String(),
				Score: explore.BestScore,
				Fn:    addBrotli7WithStat,
			})
		case "brotli_8":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli8.String(),
				Score: explore.BestScore,
				Fn:    addBrotli8WithStat,
			})
		case "brotli_9":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli9.String(),
				Score: explore.BestScore,
				Fn:    addBrotli9WithStat,
			})
		case "brotli_10":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli10.String(),
				Score: explore.BestScore,
				Fn:    addBrotli10WithStat,
			})
		case "brotli_11":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDBrotli11.String(),
				Score: explore.BestScore,
				Fn:    addBrotli11WithStat,
			})
		case "zlib_1":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib1.String(),
				Score: explore.BestScore,
				Fn:    addZlib1WithStat,
			})
		case "zlib_2":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib2.String(),
				Score: explore.BestScore,
				Fn:    addZlib2WithStat,
			})
		case "zlib_3":
			{
				variantsWithStat = append(variantsWithStat, explore.Variant{
					Idx:   idx,
					Name:  compression.CodecIDZlib3.String(),
					Score: explore.BestScore,
					Fn:    addZlib3WithStat,
				})
			}
		case "zlib_4":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib4.String(),
				Score: explore.BestScore,
				Fn:    addZlib4WithStat,
			})
		case "zlib_5":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib5.String(),
				Score: explore.BestScore,
				Fn:    addZlib5WithStat,
			})
		case "zlib_6":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib6.String(),
				Score: explore.BestScore,
				Fn:    addZlib6WithStat,
			})
		case "zlib_7":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib7.String(),
				Score: explore.BestScore,
				Fn:    addZlib7WithStat,
			})
		case "zlib_8":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib8.String(),
				Score: explore.BestScore,
				Fn:    addZlib8WithStat,
			})
		case "zlib_9":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZlib9.String(),
				Score: explore.BestScore,
				Fn:    addZlib9WithStat,
			})
		case "zstd_1":
			{
				variantsWithStat = append(variantsWithStat, explore.Variant{
					Idx:   idx,
					Name:  compression.CodecIDZstd1.String(),
					Score: explore.BestScore,
					Fn:    addZstd1WithStat,
				})
			}
		case "zstd_3":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZstd3.String(),
				Score: explore.BestScore,
				Fn:    addZstd3WithStat,
			})
		case "zstd_7":
			variantsWithStat = append(variantsWithStat, explore.Variant{
				Idx:   idx,
				Name:  compression.CodecIDZstd7.String(),
				Score: explore.BestScore,
				Fn:    addZstd7WithStat,
			})
		default:
			log.Fatalln("no function to explore")
		}
	}

	return variantsWithStat
}

func addFunctionToExploreOrExploit(filePath string, codecId compression.CodecID) int {
	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	codec := compression.NewCodec(codecId)
	c, compressErr := codec.Compress(buf)
	if compressErr != nil {
		log.Fatal(compressErr)
	}

	// return compress factor - simply how many bytes we have on output
	return len(c)
}

func addWithStat(config *cfg.Config, id compression.CodecID) (int, int64) {
	start := time.Now()
	result := addFunctionToExploreOrExploit(config.FilePath, id)
	elapsed := time.Since(start)

	var score int64
	if cfg.RUNTIME == config.ScoreType {
		score = int64(elapsed*time.Nanosecond) / int64(time.Microsecond)
	} else if cfg.COMPRESS_FACTOR == config.ScoreType {
		// factor in KB
		score = int64(result)
	}

	return result, score
}

// generate function stubs for explorations
// snappy
func addSnappyWithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDSnappy)
}

// lz4
func addLz4WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDLz4)
}
func addLz4HighCompressionWithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDLz4HighCompression)
}

// brotly
func addBrotli1WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli1)
}
func addBrotli2WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli2)
}
func addBrotli3WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli3)
}
func addBrotli4WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli4)
}
func addBrotli5WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli5)
}
func addBrotli6WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli6)
}
func addBrotli7WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli7)
}
func addBrotli8WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli8)
}
func addBrotli9WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli9)
}
func addBrotli10WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli10)
}
func addBrotli11WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDBrotli11)
}

// zlib
func addZlib1WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib1)
}
func addZlib2WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib2)
}
func addZlib3WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib3)
}
func addZlib4WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib4)
}
func addZlib5WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib5)
}
func addZlib6WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib6)
}
func addZlib7WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib7)
}
func addZlib8WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib8)
}
func addZlib9WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZlib9)
}

// zstd
func addZstd1WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZstd1)
}
func addZstd3WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZstd3)
}
func addZstd7WithStat(config *cfg.Config) (int, int64) {
	return addWithStat(config, compression.CodecIDZstd7)
}
