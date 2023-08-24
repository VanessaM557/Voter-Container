package main
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handleGetVoterByID(w http.ResponseWriter, r *http.Request, voterList *VoterList) {
	voterID := getPathVariable(r, "id")
	voter, err := GetVoterByID(voterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, voter)
}

func handleUpdateVoterByID(w http.ResponseWriter, r *http.Request, voterList *VoterList) {
	voterID := getPathVariable(r, "id")
	var updatedVoter Voter
	err := json.NewDecoder(r.Body).Decode(&updatedVoter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = UpdateVoterByID(voterID, updatedVoter.FirstName, updatedVoter.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleDeleteVoterByID(w http.ResponseWriter, r *http.Request, voterList *VoterList) {
	voterID := getPathVariable(r, "id")
	err := DeleteVoterByID(voterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getPathVariable(r *http.Request, paramName string) string {
	parts := strings.Split(r.URL.Path, "/")
	for i, part := range parts {
		if part == paramName && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
