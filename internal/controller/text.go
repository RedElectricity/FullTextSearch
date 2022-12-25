package controller

import (
	"FullTextSearch/internal/logic"
	"FullTextSearch/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"FullTextSearch/api/v1"
)

var Text = cText{}

type cText struct{}

func (r *cText) RecordText(ctx context.Context, req *v1.RecordTextReq) (res *v1.JsonRes, err error) {
	text := g.RequestFromCtx(ctx).Get("text")
	textID := logic.MakeIndex(text.String())

	g.RequestFromCtx(ctx).Response.WriteJson(entity.InputTextRespond{
		Status: 0,
		TextID: textID,
	})
	return
}

func (r *cText) SearchText(ctx context.Context, req *v1.SearchTextReq) (res *v1.JsonRes, err error) {
	text := g.RequestFromCtx(ctx).Get("text")
	textIDs := logic.SearchText(text.String())

	g.RequestFromCtx(ctx).Response.WriteJson(entity.SearchTextRespond{
		Status: 0,
		TextID: textIDs.Slice(),
	})

	return
}

func (r *cText) GetText(ctx context.Context, req *v1.GetTextReq) (res *v1.JsonRes, err error) {
	textID := g.RequestFromCtx(ctx).Get("textId")
	text := logic.GetText(textID.String())

	g.RequestFromCtx(ctx).Response.WriteJson(entity.GetTextRespond{
		Status: 0,
		Text:   text,
	})

	return
}
