package logic

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/util/gconv"
)

func SearchText(text string) *garray.StrArray {
	resultTextID := garray.NewStrArray()

	index := analyzeText(splitText(text))

	resultTextIndex := getIndexRecord(index)

	resultTextIndex.Iterator(func(k string, v interface{}) bool {
		if v != nil {
			for _, s := range gconv.Strings(v) {
				if !resultTextID.ContainsI(s) {
					resultTextID.Append(s)
				}
			}
		}

		return true
	})

	return resultTextID
}
