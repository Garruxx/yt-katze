package utils

import "github.com/graphql-go/graphql"

func GqlResolver[T any](
	params map[string]any,
	resolver func(graphql.ResolveParams) (any, error),
) (T, error) {
	param := graphql.ResolveParams{
		Args: params,
	}
	result, err := resolver(param)
	if err != nil {
		return result.(T), err
	}
	return result.(T), nil
}
