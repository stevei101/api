package main

import (
	"context"
	"fmt"
	"log"

	dialogflow "google.golang.org/api/dialogflow/v3"
	"google.golang.org/api/option"
)

func main() {
	// Replace with your project ID and service account key file path
	projectID := "your-project-id"
	keyFilePath := "path/to/your/key.json"

	// Create a new Dialogflow CX client
	ctx := context.Background()
	client, err := dialogflow.NewService(ctx, option.WithCredentialsFile(keyFilePath))
	if err != nil {
		log.Fatalf("Error creating Dialogflow CX client: %v", err)
	}

	// Create a new agent
	agent := &dialogflow.Agent{
		DisplayName: "My Agent",
	}
	createdAgent, err := client.Agents.Create(projectID, agent).Do()
	if err != nil {
		log.Fatalf("Error creating agent: %v", err)
	}
	fmt.Println("Agent created:", createdAgent.Name)

	// List agents
	agents, err := client.Agents.List(projectID).Do()
	if err != nil {
		log.Fatalf("Error listing agents: %v", err)
	}
	for _, agent := range agents.Agents {
		fmt.Println(agent.DisplayName)
	}
}
