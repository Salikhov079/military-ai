package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/military-ai/genprotos/ai"
	s "github.com/Salikhov079/military-ai/storage"
)

type AIService struct {
	stg s.InitRoot
	pb.UnimplementedAiServiceServer
}

func NewAIService(stg s.InitRoot) *AIService {
	return &AIService{stg: stg}
}

func (c *AIService) CHat(ctx context.Context, ai *pb.AiCHat) (*pb.AiCHat, error) {
	void, err := c.stg.Ai().CHat(ai)
	if err != nil {
		log.Print(err)
	}
	return void, err
}


func (c *AIService) GetHistory(ctx context.Context,text *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error) {
	void, err := c.stg.Ai().GetHistory(text)
	if err != nil {
		log.Print(err)
	}
	return void, err
}