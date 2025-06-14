package resolver

import (
	"context"

	"github.com/ilmsadmin/Zplus-SaaS/apps/backend/gateway/types"
)

// getRequestContext extracts the request context from GraphQL context
func getRequestContext(ctx context.Context) *types.RequestContext {
	if reqCtx := ctx.Value("request_context"); reqCtx != nil {
		return reqCtx.(*types.RequestContext)
	}
	return &types.RequestContext{}
}