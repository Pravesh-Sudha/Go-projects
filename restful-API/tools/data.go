package tools

import "sync"

type Tool struct {
	Name			string `json:"name"`
	Ddescription	string `json:"description"`
	Category		string `json:"category"` 
} 

var toolStore = []Tool{
	{"Docker", "A Platform to build, run and share Containerized applications.", "Containerization"},
	{"Kubernetes", "An orchestration tool for managing containerized workloads.", "Orchestration"},
	{"Jenkins", "An open-source automation server for CICD", "CI/CD"},
}

var mu sync.Mutex