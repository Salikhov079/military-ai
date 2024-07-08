package postgres

import (
	pbs "github.com/Salikhov079/military-ai/genprotos/ai"
)

type InitRoot interface {
	Ai() Ai
}

type Ai interface {
	CHat(text *pbs.AiCHat) (*pbs.AiCHat, error)
	GetHistory(text *pbs.GetHistoryRequest) (*pbs.GetHistoryResponse, error)
}