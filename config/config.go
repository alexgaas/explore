package config

import (
	compression "explore/compress"
	"flag"
	"strings"
)

type ScoreType int

const (
	RUNTIME         ScoreType = 0
	COMPRESS_FACTOR ScoreType = 1
)

type Flags struct {
	count     int
	scoreType int
	codec     string
	filePath  string
}

type Config struct {
	Count     int
	ScoreType ScoreType
	Codecs    []string
	FilePath  string
}

func GetAppFlags() Flags {
	flags := Flags{}
	flag.IntVar(&flags.count, "c", 16256, "Number of operations per each defined codec/file. Default value - 16256")

	flag.IntVar(&flags.scoreType, "s", int(RUNTIME), "Score type: 0 - Runtime, 1 - Compress factor. Default value - 0")

	defaultBrotliCodec := compression.CodecIDBrotli6.String()
	defaultLzZCodec := compression.CodecIDLz4HighCompression.String()
	defaultSnappyCodec := compression.CodecIDSnappy.String()
	defaultZlibCodec := compression.CodecIDZlib5.String()
	defaultZstdCodec := compression.CodecIDZstd3.String()
	var codecs = defaultBrotliCodec + " " + defaultLzZCodec + " " + defaultSnappyCodec + " " + defaultZlibCodec + " " + defaultZstdCodec
	flag.StringVar(&flags.codec, "d", codecs, "Define codecs.\n"+
		"Options - "+
		"snappy, lz4, lz4_high_compression, brotli_1, brotli_2, brotli_3, brotli_4, brotli_5, brotli_6, brotli_7, "+
		"brotli_8, brotli_9, brotli_10, brotli_11, zlib_1, zlib_2, zlib_3, zlib_4, zlib_5, zlib_6, zlib_7, zlib_8, "+
		"zlib_9, zstd_1, zstd_3, zstd_7\n"+
		"Default value - brotli_6 lz4_high_compression snappy")

	flag.StringVar(&flags.filePath, "f", "test/book1", "File path to test score (Runtime / Compress factor. "+
		"Default - folder [test] in the root of source code folder")

	flag.Parse()
	return flags
}

func GetNewConfig(flags Flags) (*Config, error) {
	var err error
	var config Config

	config.Count = flags.count

	config.ScoreType = ScoreType(flags.scoreType)

	config.Codecs = strings.Fields(flags.codec)

	config.FilePath = flags.filePath

	return &config, err
}
