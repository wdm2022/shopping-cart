package utils

import (
	"context"
	"time"
)

func ContextWithTimeOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func ContextWithTimeOut10() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
