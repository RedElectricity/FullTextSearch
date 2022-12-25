package logic

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/shabbyrobe/go-porter2"
	"github.com/yanyiwu/gojieba"
	"unicode"
)

func splitText(Text string) []string {
	var ctx = gctx.New()
	var useHMM, _ = g.Cfg().Get(ctx, "Jieba.useHMM")

	spliter := gojieba.NewJieba()
	defer spliter.Free()

	words := spliter.CutForSearch(Text, useHMM.Bool())
	for i := 0; i < len(words); i++ {
		nowWord := words[i]
		var isEnglish = true

		for _, i3 := range nowWord {
			if unicode.IsLetter(i3) == false {
				isEnglish = false
				break
			}
		}

		if isEnglish {
			words[i] = porter2.Stem(nowWord, porter2.StemFlag(1))
		}
	}

	return words
}

func analyzeText(words []string) *garray.StrArray {
	wordMap := garray.NewStrArray()

	for _, word := range words {
		if !wordMap.ContainsI(word) {
			wordMap.Append(word)
		}
	}

	return wordMap
}

func MakeIndex(text string) string {
	index := analyzeText(splitText(text))

	textID := recordAText(text)
	recordIndex(index, textID)

	return textID
}
