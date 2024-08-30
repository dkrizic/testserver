package graph

import "database/sql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	dB *sql.DB
}

type Opts func(*Resolver)

func DB(dB *sql.DB) Opts {
	return func(r *Resolver) {
		r.dB = dB
	}
}

func NewResolver(opts ...Opts) *Resolver {
	r := &Resolver{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
