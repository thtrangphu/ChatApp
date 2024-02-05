package apiproto

import (
	"context"
	"encoding/json"

	"github.com/mekanican/chat-backend/internal/config"
	grpc "google.golang.org/grpc"
)

type keyAuth struct {
	key string
}

func (t keyAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "apikey " + t.key,
	}, nil
}

func (t keyAuth) RequireTransportSecurity() bool {
	return false
}

var conn *grpc.ClientConn
var client CentrifugoApiClient

type MessageData struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func Initialize() error {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithPerRPCCredentials(keyAuth{config.GetString("GRPC_TOKEN")}))
	if err != nil {
		return err
	}
	client = NewCentrifugoApiClient(conn)

	return nil
}

func PublishData(channel string, message MessageData) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = client.Publish(context.Background(), &PublishRequest{
		Channel: channel,
		Data:    data,
	})

	if err != nil {
		return err
	}
	return nil
}

func CleanUp() {
	conn.Close()
}
