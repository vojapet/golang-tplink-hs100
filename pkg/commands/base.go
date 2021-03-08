package commands

import (
)

type EmptyValue struct {
}

type pair struct {
	key string
	value interface{}
}

type Command pair
type Query map[string]interface{}
type Response map[string]interface{}

func BuildQuery(subQueries ...pair) Query {
	query := make(Query)
	for _, subQuery := range subQueries {
		query[subQuery.key] = subQuery.value
	}
	return query
}
