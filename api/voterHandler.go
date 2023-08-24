package main

import (
	"encoding/json"
	"net/http"
	"github.com/go-redis/redis/v8"
)

func handleGetVoterByID(appContext *AppContext, w http.ResponseWriter, r *http.Request, id string) {
	voterJSON, err := appContext.Rdb.Get(appContext.Ctx, "voter:"+id).Result()
	if err == redis.Nil {
		http.Error(w, "Voter cannot be found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving voter", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(voterJSON))
}

func handleUpdateVoterByID(appContext *AppContext, w http.ResponseWriter, r *http.Request, id string) {
	var updatedVoter Voter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedVoter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = UpdateByID(appContext, id, updatedVoter.FirstName, updatedVoter.LastName)
	if err != nil {
		http.Error(w, "Error updating voter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Voter updated successfully"}`))
}

func handleDeleteVoterByID(appContext *AppContext, w http.ResponseWriter, r *http.Request, id string) {
	err := DeleteVoterByID(appContext, id)
	if err != nil {
		http.Error(w, "Error deleting voter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Voter deleted successfully"}`))
}

