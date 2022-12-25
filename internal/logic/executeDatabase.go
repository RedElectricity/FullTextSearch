package logic

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/guid"
)

func recordAText(Text string) string {
	TextID := guid.S()
	ctx := gctx.New()

	_, err := g.Redis("TextRecord").Do(ctx, "SET", TextID, Text)
	if err != nil {
		panic(err)
	}

	return TextID
}

func recordIndex(wordIndex *garray.StrArray, TextID string) {
	ctx := gctx.New()

	wordIndex.Iterator(func(k int, v string) bool {
		result, err := g.Redis().Do(ctx, "GET", v)
		if err != nil {
			panic(err)
			return false
		}

		if result != nil {
			_, err = g.Redis("IndexRecord").Do(ctx, "SET", v, append(result.Strings(), TextID))
			if err != nil {
				panic(err)
				return false
			}
		} else {
			_, err = g.Redis("IndexRecord").Do(ctx, "SET", v, []string{TextID})
			if err != nil {
				panic(err)
				return false
			}
		}

		return true
	})
}

func getIndexRecord(words *garray.StrArray) *gmap.StrAnyMap {
	resultIndex := gmap.NewStrAnyMap()
	ctx := gctx.New()

	words.Iterator(func(k int, v string) bool {
		val, err := g.Redis("IndexRecord").Do(ctx, "GET", v)
		if err != nil {
			panic(err)
			return false
		}

		if val != nil {
			resultIndex.Set(v, nil)
		} else {
			resultIndex.Set(v, val.Strings())
		}

		return true
	})

	return resultIndex
}

func GetText(TextID string) string {
	ctx := gctx.New()

	val, err := g.Redis("TextRecord").Do(ctx, "GET", TextID)
	if err != nil {
		panic(err)
	}

	if val != nil {
		return ""
	}
	return val.String()
}
