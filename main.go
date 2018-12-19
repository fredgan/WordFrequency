package main

import (
	"bufio"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io"
	"os"
	"sort"
)

const max_line = 0

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, 0)

	for k, v := range m {
		p = append(p, Pair{k, v})
	}
	sort.Sort(p)
	return p
}

func main() {
	var words []string
	x := gojieba.NewJieba()
	defer x.Free()

	wordMap := map[string]int{}
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	count := 1
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		// 分词
		words = x.Cut(line, true)
		for _, w := range words {
			wordMap[w]++
		}
		count++
		if count%10000 == 0 { // 打印进度
			fmt.Println("line=", count)
		}
		if max_line > 0 && count > max_line {
			break
		}
	}

	fmt.Println("len(map)=", len(wordMap))

	line := 0
	sortfileObj, err := os.OpenFile("sort_result.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	sortResult := sortMapByValue(wordMap)
	// 打印500行
	for _, v := range sortResult {
		if len([]rune(v.Key)) < 2 {
			continue
		}
		fmt.Println("v.Key", v.Key, "v.Value", v.Value)
		line++
		if line > 500 {
			break
		}
	}
	// 保存到文件
	for _, v := range sortResult {
		if len([]rune(v.Key)) < 2 {
			continue
		}
		io.WriteString(sortfileObj, fmt.Sprintf("%v,%v\n", v.Key, v.Value))
	}

}
