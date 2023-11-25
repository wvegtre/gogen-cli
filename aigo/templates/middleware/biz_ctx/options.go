package biz_ctx

import (
	"net/http"

	"gen-templates/internal/common"
	"gen-templates/middleware/log"

	"go.uber.org/zap"
)

type RequestCtxOption func(args CtxOptionArgs)

type CtxOptionArgs map[string]any

const (
	bizHeaderKeyAccountID = "biz_account_id"
	bizHeaderKeyStoreID   = "biz_store_id"
	bizHeaderKeyOperator  = "operator"
)

const (
	bizAccessToken = "biz_access_token"
	bizOperator    = "biz_operator"
)

func SetOperatorByRequestHeader(header http.Header) RequestCtxOption {
	return func(args CtxOptionArgs) {
		if args == nil {
			args = make(map[string]any)
		}
		args[bizOperator] = header.Get(bizHeaderKeyOperator)
	}
}

func SetAccountInfoByRequestHeader(header http.Header) RequestCtxOption {
	return func(args CtxOptionArgs) {
		if args == nil {
			args = make(map[string]any)
		}
		at := common.AccessToken{
			AccountID: header.Get(bizHeaderKeyAccountID),
			StoreID:   header.Get(bizHeaderKeyStoreID),
		}
		args[bizAccessToken] = at
		args[log.GetBizLoggerKey()] = append(args[log.GetBizLoggerKey()].([]zap.Field),
			zap.String(bizHeaderKeyAccountID, at.AccountID),
			zap.String(bizHeaderKeyStoreID, at.StoreID),
		)
	}
}
