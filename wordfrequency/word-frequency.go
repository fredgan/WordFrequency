package wordfrequency

import (
	"bufio"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io"
	"os"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// 排序map
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, 0)

	for k, v := range m {
		p = append(p, Pair{k, v})
	}
	sort.Sort(p)
	return p
}

/**
\param [in] intputFileName. The intput words need to stat word frequency
\param [in,out] outputFileName. The output file for the word frequency by sorted result
\param [in] minRuneLen. The min of the word of the cutting to save in the output file.
         such as, if minRuneLen is set to 2, then '你' or 'a' and so on frequency in the result will be ignored.
 */
func WordFrequency(intputFileName string, outputFileName string, minRuneLen int) {
	var words []string
	x := gojieba.NewJieba()
	defer x.Free()

	wordMap := map[string]int{}
	f, err := os.Open(intputFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 读入文件并分词保存到map中
	count := 1
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //读入一行
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
	}

	// 打印map的大小
	fmt.Println("Count total words of result's lines is ", len(wordMap))


	sortfileObj, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}

	// 对词频降序排列
	sortResult := sortMapByValue(wordMap)
	// Debug print 500 lines in head
	//line := 0
	//for _, v := range sortResult {
	//	if len([]rune(v.Key)) < minRuneLen {
	//		continue
	//	}
	//	fmt.Println("v.Key", v.Key, "v.Value", v.Value)
	//	line++
	//	if line > 500 {
	//		break
	//	}
	//}

	// 保存到文件
	for _, v := range sortResult {
		if len([]rune(v.Key)) < minRuneLen {
			continue
		}
		_,_ = io.WriteString(sortfileObj, fmt.Sprintf("%v : %v\n", v.Key, v.Value))
	}
}