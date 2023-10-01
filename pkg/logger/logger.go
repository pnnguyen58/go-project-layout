package logger

import (
	"context"
	"go.uber.org/zap"
)

func New(ctx context.Context) (*zap.Logger, error) {
	return zap.NewDevelopment()
}
