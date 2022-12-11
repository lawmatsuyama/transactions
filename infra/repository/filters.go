package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// IsZeroFunc represents a generic type func to check if value is zero value
type IsZeroFunc[T any] func(value T) bool

func isZeroComparable[T comparable](value T) bool {
	var zero T
	return value == zero
}

func isZeroTime(value time.Time) bool {
	return value.IsZero()
}

func filterSimple[T any](filter bson.D, key string, value T, isZero IsZeroFunc[T]) bson.D {
	if isZero(value) {
		return filter
	}

	return append(filter, bsonE(key, value))
}

func filterRange[T any](filter bson.D, key string, valueGreater, valueLess T, isZero IsZeroFunc[T]) bson.D {
	d := bson.D{}
	if !isZero(valueGreater) {
		d = append(d, bsonE("$gte", valueGreater))
	}

	if !isZero(valueLess) {
		d = append(d, bsonE("$lte", valueLess))
	}

	if len(d) == 0 {
		return filter
	}

	return append(filter, bsonE(key, d))

}

func bsonE(key string, value any) bson.E {
	return bson.E{Key: key, Value: value}
}
