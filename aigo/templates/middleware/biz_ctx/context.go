package biz_ctx

import (
	"context"

	"gen-templates/internal/common"
)

type BizContext struct {
	common.AccessToken
}

func AppendFieldsToContext(parent context.Context, key, value any) context.Context {
	if parent == nil {
		return nil
	}
	ctx := context.WithValue(parent, key, value)
	return ctx
}

func AppendBizFieldsToContextByRequestHeader(ctx context.Context, ops ...RequestCtxOption) context.Context {
	op := make(CtxOptionArgs)
	for _, o := range ops {
		o(op)
	}

	if op == nil {
		return ctx
	}

	for k, v := range op {
		ctx = AppendFieldsToContext(ctx, k, v)
	}

	return ctx
}

func GetAccountInfoFromContext(ctx context.Context) (common.AccessToken, bool) {
	if ctx == nil {
		return common.AccessToken{}, false
	}
	at, ok := ctx.Value(bizAccessToken).(common.AccessToken)
	return at, ok
}
