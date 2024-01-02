package models

import "math/big"

type Food struct {
	Name        string   `json:"name"`
	TraceNumber *big.Int `json:"traceNumber"`
	TraceName   string   `json:"traceName"`
	Quality     uint8    `json:"quality"`
}
