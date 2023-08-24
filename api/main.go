package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/go-redis/redis/v8"
)

// structs

type AppContext struct {
	Ctx context.Context
	Rdb *redis.Client
}

type Voter struct {
	ID        string
	FirstName string
	LastName  string
}

type HealthData struct {
	TotalAPICalls           int
	TotalAPICallsWithError  int
	BootTime                time.Time
	Uptime                  string
}

type VoterList struct {
	Voters     []Voter
	HealthData HealthData
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	appContext := &AppContext{
		Ctx: ctx,
		Rdb: rdb,
	}

	http.HandleFunc("/voters", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetAllVoters(appContext, w, r)
		case http.MethodPost:
			handleCreateNewVoter(appContext, w, r)
		}
	})

	http.HandleFunc("/voters/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/voters/")
		switch r.Method {
		case http.MethodGet:
			handleGetVoterByID(appContext, w, r, id)
		case http.MethodPut:
			handleUpdateVoterByID(appContext, w, r, id)
		case http.MethodDelete:
			handleDeleteVoterByID(appContext, w, r, id)
		}
	})

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handleGetAllVoters(appContext *AppContext, w http.ResponseWriter, r *http.Request) {
	voterKeys, err := appContext.Rdb.Keys(appContext.Ctx, "voter:*").Result() // Fix here
	if err != nil {
		http.Error(w, "Error retrieving voters", http.StatusInternalServerError)
		return
	}
	voters := []Voter{}
	for _, key := range voterKeys {
		voterJSON, err := appContext.Rdb.Get(appContext.Ctx, key).Result() // Fix here
		if err != nil {
			http.Error(w, "Error retrieving voter", http.StatusInternalServerError)
			return
		}
		var voter Voter
		err = json.Unmarshal([]byte(voterJSON), &voter)
		if err != nil {
			http.Error(w, "Error decoding voter", http.StatusInternalServerError)
			return
		}
		voters = append(voters, voter)
	}
	responseJSON, err := json.Marshal(voters)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func handleCreateNewVoter(appContext *AppContext, w http.ResponseWriter, r *http.Request) {
	var voter Voter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&voter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	voterJSON, err := json.Marshal(voter)
	if err != nil {
		http.Error(w, "Error encoding voter", http.StatusInternalServerError)
		return
	}
	err = appContext.Rdb.Set(appContext.Ctx, "voter:"+voter.ID, voterJSON, 0).Err() // Fix here
	if err != nil {
		http.Error(w, "Error saving voter", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(voterJSON)
}


