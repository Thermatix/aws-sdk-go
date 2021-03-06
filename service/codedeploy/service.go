package codedeploy

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// CodeDeploy is a client for CodeDeploy.
type CodeDeploy struct {
	*aws.Service
}

// New returns a new CodeDeploy client.
func New(config *aws.Config) *CodeDeploy {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "codedeploy",
		APIVersion:   "2014-10-06",
		JSONVersion:  "1.1",
		TargetPrefix: "CodeDeploy_20141006",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &CodeDeploy{service}
}
