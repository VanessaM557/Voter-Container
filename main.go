package main
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/go-redis/redis/v8" 
)

var context = context.Background()
var rdb *redis.Client

//Vote struct 

type Voter struct {
	ID   string
	Name string
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	http.HandleFunc("/voters", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetAllVoters(w, r)
		case http.MethodPost:
			handleCreateNewVoter(w, r)
		default:
			http.Error(w, "Method is not permmitted", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/voters/", func(w http.ResponseWriter, r *http.Request) {
		handleGetVoterByID(w, r)
	})

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handleGetAllVoters(w http.ResponseWriter, r *http.Request) {
	voterKeys, err := rdb.Keys(context, "voter:*").Result()
	if err != nil {
		http.Error(w, "Error retrieving voters", http.StatusInternalServerError)
		return
	}

	voters := []Voter{}
	for _, key := range voterKeys {
		voterJSON, err := rdb.Get(context, key).Result()
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

func handleCreateNewVoter(w http.ResponseWriter, r *http.Request) {
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

	err = rdb.Set(context, "voter:"+voter.ID, voterJSON, 0).Err()
	if err != nil {
		http.Error(w, "Error saving voter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(voterJSON)
}

func handleGetVoterByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/voters/")

	voterJSON, err := rdb.Get(context, "voter:"+id).Result()
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
