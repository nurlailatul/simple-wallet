package contextprop

import "context"

type ContextKey string

const (
	KeyActorID    ContextKey = "actorId"
	KeyActorEmail ContextKey = "actorEmail"
)

func GetContextValue(ctx context.Context, key ContextKey) string {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return value
}
