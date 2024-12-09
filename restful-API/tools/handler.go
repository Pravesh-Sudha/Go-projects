package tools

import (
	"encoding/json"
	"net/http"
)

func ListTools(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Request Method",http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toolStore)
}

func AddTool(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method Type", http.StatusMethodNotAllowed)
		return
	}

	var newTool Tool
	if err := json.NewDecoder(r.Body).Decode(&newTool); err != nil {
		http.Error(w, "Invalid Requset Body", http.StatusBadRequest)
		return
	}
	mu.Lock()
	toolStore = append(toolStore, newTool)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTool)

}

func ToolDetails(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method Type", http.StatusMethodNotAllowed)
		return
	}

	toolName := r.URL.Query().Get("name")
	if toolName == "" {
		http.Error(w, "Missing Tool name", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, tool := range toolStore {
		if tool.Name == toolName {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tool)
			return
		}
	} 
	http.Error(w, "Tool Not found", http.StatusNotFound)	
}