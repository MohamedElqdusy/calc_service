package api

import (
	"golang.org/x/net/context"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// Generating a response to a summation request
func (s *Server) Sum(ctx context.Context, in *Args) (*Result, error) {
	var resultValue int64
	resultValue = in.Arg1 + in.Arg2
	log.Printf("Summation %d", resultValue)
	return &Result{resultValue}, nil
}

// Generating a response to a division request
func (s *Server) Div(ctx context.Context, in *Args) (*Result, error) {
	var resultValue int64
	resultValue = in.Arg1 / in.Arg2
	log.Printf("Division %d", resultValue)
	return &Result{resultValue}, nil
}

// Generating a response to a multiplication request
func (s *Server) Mul(ctx context.Context, in *Args) (*Result, error) {
	var resultValue int64
	resultValue = in.Arg1 * in.Arg2
	log.Printf("Multiplication %d", resultValue)
	return &Result{resultValue}, nil
}

// Generating a response to a subtraction request
func (s *Server) Sub(ctx context.Context, in *Args) (*Result, error) {
	var resultValue int64
	resultValue = in.Arg1 - in.Arg2
	log.Printf("Subtraction %d", resultValue)
	return &Result{resultValue}, nil
}
