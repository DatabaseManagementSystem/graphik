package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	apipb "github.com/autom8ter/graphik/gen/go"
	generated1 "github.com/autom8ter/graphik/gen/gql/generated"
	"github.com/autom8ter/graphik/gen/gql/model"
	"github.com/autom8ter/graphik/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *mutationResolver) CreateDoc(ctx context.Context, input model.DocConstructor) (*model.Doc, error) {
	doc, err := r.client.CreateDoc(ctx, protoDocC(&input))
	if err != nil {
		return nil, err
	}
	return gqlDoc(doc), nil
}

func (r *mutationResolver) EditDoc(ctx context.Context, input model.Edit) (*model.Doc, error) {
	res, err := r.client.EditDoc(ctx, protoEdit(&input))
	if err != nil {
		return nil, err
	}
	return gqlDoc(res), nil
}

func (r *mutationResolver) EditDocs(ctx context.Context, input model.EditFilter) (*model.Docs, error) {
	docs, err := r.client.EditDocs(ctx, protoEditFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlDocs(docs), nil
}

func (r *mutationResolver) CreateConnection(ctx context.Context, input model.ConnectionConstructor) (*model.Connection, error) {
	res, err := r.client.CreateConnection(ctx, protoConnectionC(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnection(res), nil
}

func (r *mutationResolver) EditConnection(ctx context.Context, input model.Edit) (*model.Connection, error) {
	res, err := r.client.EditConnection(ctx, protoEdit(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnection(res), nil
}

func (r *mutationResolver) EditConnections(ctx context.Context, input model.EditFilter) (*model.Connections, error) {
	connections, err := r.client.EditConnections(ctx, protoEditFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnections(connections), nil
}

func (r *mutationResolver) Publish(ctx context.Context, input model.OutboundMessage) (*emptypb.Empty, error) {
	return r.client.Publish(ctx, &apipb.OutboundMessage{
		Channel: input.Channel,
		Data:    apipb.NewStruct(input.Data),
	})
}

func (r *queryResolver) Ping(ctx context.Context, input *emptypb.Empty) (*model.Pong, error) {
	res, err := r.client.Ping(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return &model.Pong{Message: res.GetMessage()}, nil
}

func (r *queryResolver) GetSchema(ctx context.Context, input *emptypb.Empty) (*model.Schema, error) {
	res, err := r.client.GetSchema(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return gqlSchema(res), nil
}

func (r *queryResolver) Me(ctx context.Context, input *emptypb.Empty) (*model.Doc, error) {
	res, err := r.client.Me(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return gqlDoc(res), nil
}

func (r *queryResolver) GetDoc(ctx context.Context, input model.PathInput) (*model.Doc, error) {
	res, err := r.client.GetDoc(ctx, protoIPath(&input))
	if err != nil {
		return nil, err
	}
	return gqlDoc(res), nil
}

func (r *queryResolver) SearchDocs(ctx context.Context, input model.Filter) (*model.Docs, error) {
	res, err := r.client.SearchDocs(ctx, protoFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlDocs(res), nil
}

func (r *queryResolver) DepthSearchDocs(ctx context.Context, input model.DepthFilter) (*model.DocTraversals, error) {
	res, err := r.client.DepthSearchDocs(ctx, protoDepthFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlTraversals(res), nil
}

func (r *queryResolver) GetConnection(ctx context.Context, input model.PathInput) (*model.Connection, error) {
	res, err := r.client.GetConnection(ctx, protoIPath(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnection(res), nil
}

func (r *queryResolver) SearchConnections(ctx context.Context, input model.Filter) (*model.Connections, error) {
	res, err := r.client.SearchConnections(ctx, protoFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnections(res), nil
}

func (r *queryResolver) ConnectionsFrom(ctx context.Context, input model.ConnectionFilter) (*model.Connections, error) {
	res, err := r.client.ConnectionsFrom(ctx, protoConnectionFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnections(res), nil
}

func (r *queryResolver) ConnectionsTo(ctx context.Context, input model.ConnectionFilter) (*model.Connections, error) {
	res, err := r.client.ConnectionsFrom(ctx, protoConnectionFilter(&input))
	if err != nil {
		return nil, err
	}
	return gqlConnections(res), nil
}

func (r *queryResolver) AggregateDocs(ctx context.Context, input model.AggregateFilter) (interface{}, error) {
	res, err := r.client.AggregateDocs(ctx, protoAggFilter(&input))
	if err != nil {
		return nil, err
	}
	return res.GetNumberValue(), nil
}

func (r *queryResolver) AggregateConnections(ctx context.Context, input model.AggregateFilter) (interface{}, error) {
	res, err := r.client.AggregateConnections(ctx, protoAggFilter(&input))
	if err != nil {
		return nil, err
	}
	return res.GetNumberValue(), nil
}

func (r *subscriptionResolver) Subscribe(ctx context.Context, input model.ChannelFilter) (<-chan *model.Message, error) {
	ch := make(chan *model.Message)
	stream, err := r.client.Subscribe(ctx, protoChannelFilter(&input))
	if err != nil {
		return nil, err
	}
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				msg, err := stream.Recv()
				if err != nil {
					logger.Error("failed to receive subsription message", zap.Error(err))
					continue
				}
				ch <- &model.Message{
					Channel:   msg.GetChannel(),
					Data:      msg.GetData().AsMap(),
					Sender:    gqlPath(msg.GetSender()),
					Timestamp: msg.GetTimestamp().AsTime(),
				}
			}
		}
	}()
	return ch, nil
}

func (r *subscriptionResolver) SubscribeChanges(ctx context.Context, input model.ExpressionFilter) (<-chan *model.Change, error) {
	ch := make(chan *model.Change)
	stream, err := r.client.SubscribeChanges(ctx, protoExpressionFilter(&input))
	if err != nil {
		return nil, err
	}
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				msg, err := stream.Recv()
				if err != nil {
					logger.Error("failed to receive change", zap.Error(err))
					continue
				}
				ch <- &model.Change{
					Method:        msg.GetMethod(),
					Identity:      gqlDoc(msg.GetIdentity()),
					Timestamp:     msg.GetTimestamp().AsTime(),
					PathsAffected: gqlPaths(msg.GetPathsAffected()),
				}
			}
		}
	}()
	return ch, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Subscription returns generated1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated1.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
