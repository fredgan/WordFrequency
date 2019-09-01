package main

import (
	"github.com/fredgan/WordFrequency/wordfrequency"
)

const (
	max_line = 0      // 用于测试
	min_rune_len = 2  // 只输出这个长度字符或者以上的字符串
)

const (
	inputFileName = "example_data/input.txt"
	outputFileName = "example_data/output.txt"
)

func main() {
	wordfrequency.WordFrequency(inputFileName, outputFileName, min_rune_len)
}
