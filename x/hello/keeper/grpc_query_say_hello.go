// package keeper

// import (
// 	"context"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// 	"hello/x/hello/types"
// )

// func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	// TODO: Process the query
// 	_ = ctx

//		return &types.QuerySayHelloResponse{}, nil
//	}
package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"hello/x/hello/types"
)

func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    // Validation and Context unwrapping
    ctx := sdk.UnwrapSDKContext(goCtx)

    _ = ctx
    // Custom Response
    return &types.QuerySayHelloResponse{Name: fmt.Sprintf("Dai che e' Venerdi Dio Cane %s!", req.Name)}, nil
}