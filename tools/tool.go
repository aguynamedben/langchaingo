package tools

import "context"

// Tool is a tool for the llm agent to interact with different applications.
type Tool interface {
	Name() string
	Description() string
	IsMultiInput() bool
	CallSingle(ctx context.Context, input string) (string, error)
	CallMulti(ctx context.Context, input map[string]any) (string, error)
}
