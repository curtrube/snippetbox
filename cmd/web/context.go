package main

// contextKey our own custom type for context key.
type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")
