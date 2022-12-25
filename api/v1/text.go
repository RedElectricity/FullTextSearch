package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type RecordTextReq struct {
	g.Meta `path:"/record" tags:"Record" method:"post" summary:"Record a Text"`
}

type SearchTextReq struct {
	g.Meta `path:"/search" tags:"Search" method:"post" summary:"Search A Text"`
}

type GetTextReq struct {
	g.Meta `path:"/getText" tags:"Search" method:"post" summary:"Search A Text"`
}

type JsonRes struct {
	g.Meta `mine:"application/json"`
}
