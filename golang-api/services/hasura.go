package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type HasuraService struct {
	endpoint    string
	adminSecret string
}

func NewHasuraService() *HasuraService {
	return &HasuraService{
		endpoint:    os.Getenv("HASURA_ENDPOINT"),
		adminSecret: os.Getenv("HASURA_ADMIN_SECRET"),
	}
}

type HasuraEvent struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

func (s *HasuraService) TriggerEvent(eventType string, data map[string]interface{}) error {
	event := HasuraEvent{
		Type: eventType,
		Data: data,
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.endpoint+"/v1/metadata", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hasura-Admin-Secret", s.adminSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to trigger event: %s", resp.Status)
	}

	return nil
}

func (s *HasuraService) ExecuteGraphQL(query string, variables map[string]interface{}) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hasura-Admin-Secret", s.adminSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
