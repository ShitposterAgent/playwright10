package ai

import (
	"fmt"
	"gocli/types"
)

func QueryOllama(cfg *types.AIConfig, prompt string) (string, error) {
	// TODO: Implement actual HTTP call to Ollama API
	fmt.Printf("Querying Ollama at %s with model %s: %s\n", cfg.Endpoint, cfg.Model, prompt)
	return "AI response (stub)", nil
}
