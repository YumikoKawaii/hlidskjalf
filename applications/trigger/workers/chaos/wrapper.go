package chaos

import "context"

func ToTriggerFunction[Request, Response any](name string, function func(ctx context.Context, request Request) (Response, error)) TriggerFunction {
	return TriggerFunction{
		Name: name,
		Function: func(ctx context.Context) error {
			var req Request
			_, err := function(ctx, req)
			return err
		},
	}
}
