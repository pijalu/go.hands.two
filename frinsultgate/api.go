package main

import "fmt"

// insult represent an insult
type insult struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Score int    `json:"score"`
}

// String returns the string version
func (i *insult) String() string {
	return fmt.Sprintf("{id: %d, text: %s, score: %d}",
		i.ID,
		i.Text,
		i.Score)
}
