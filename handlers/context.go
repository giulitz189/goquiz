package handlers

import (
	"context"

	"github.com/nu7hatch/gouuid"
	"fmt"
)

type contextKey string

const requestIDKey = contextKey("requestID")

var requestID *uuid.UUID

func getRequestID (ctx context.Context) (id string, ok bool) {
	// Verifichiamo la validità del valore assunto da counterKey
	v := ctx.Value(requestIDKey)
	if v == nil {
		return
	}
	id, ok = fmt.Sprint(v), true
	return
}

func SetRequestID (ctx context.Context) context.Context {
	requestID, _ = uuid.NewV4()
	return context.WithValue(ctx, requestIDKey, requestID)
}