package postgres

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "github.com/Salikhov079/military-ai/genprotos/ai"
	"github.com/google/uuid"
)

type Response struct {
	Candidates    []Candidate   `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
}

type Candidate struct {
	Content       Content        `json:"content"`
	FinishReason  string         `json:"finishReason"`
	Index         int            `json:"index"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Part struct {
	Text string `json:"text"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

type AiStorage struct {
	db *sql.DB
}

func NewAiStorage(db *sql.DB) *AiStorage {
	return &AiStorage{db: db}
}
func (ai AiStorage) CHat(text *pb.AiCHat) (*pb.AiCHat, error) {
	apiKey := "AIzaSyAqYyftXSGM9ImO8AJmLxhswmJ8BNmp6iE"
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey
	
	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text": text.Text,
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return nil, err
	}

	 //response.Candidates[0].Content.Parts[0]
	 chat:=&pb.AiCHat{Text: response.Candidates[0].Content.Parts[0].Text}
	 id:=uuid.NewString()
	 _,err=ai.db.Query("insert into aistorage(id, user_id, text, reqtext) values ($1, $2, $3, $4)", 
	 id, text.UserId, chat.Text, text.Text)
	 if err!=nil{
		return nil, err
	 }
	 return chat, nil
}


func (ai AiStorage) GetHistory(text *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error){
	var  request pb.GetHistoryResponse
	query:=`select text, reqtext
		    from aistorage
			where user_id=$1`
	rows, err:=ai.db.Query(query, text.Id)
	if err!=nil{
		fmt.Println(err)
		return nil, err
	}
	for rows.Next(){
		arr:=pb.GetAllAi{}
		err=rows.Scan(&arr.ResponseText, &arr.RequestText)
		if err!=nil{
			fmt.Println(err)
			return nil, err

		}
		request.Requests=append(request.Requests, &arr)

	}
	return &request, nil
}