// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package algorithm

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type AlgorithType string

const (
	AlgoTypeBin    AlgorithType = "bin"
	AlgoTypePython AlgorithType = "python"
	AlgoTypeKey                 = "algo_type"
)

func AlgorithmTypeToContext(ctx context.Context, algoType string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, AlgoTypeKey, algoType)
}

func AlgorithmTypeFromContext(ctx context.Context) string {
	return metadata.ValueFromIncomingContext(ctx, AlgoTypeKey)[0]
}

// Algorithm is an interface that specifies the API for an algorithm.
type Algorithm interface {
	// Run executes the algorithm and returns the result.
	Run() ([]byte, error)

	// Add dataset to algorithm.
	AddDataset(dataset string)
}
