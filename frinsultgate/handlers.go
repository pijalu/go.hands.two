package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/pijalu/go.hands.two/frinsultproto"
)

func replyWithJSON(w http.ResponseWriter, reply interface{}, status int) {
	b, err := json.Marshal(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Error during marshal: %e", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// p2a converts proto model to api model
func p2a(i *frinsultproto.Frinsult) *insult {
	return &insult{
		ID:    int(i.GetID()),
		Text:  i.GetText(),
		Score: int(i.GetScore()),
	}
}

// a2p converts api to proto
func a2p(i *insult) *frinsultproto.Frinsult {
	return &frinsultproto.Frinsult{
		ID:    int64(i.ID),
		Text:  i.Text,
		Score: int64(i.Score),
	}
}

func getIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		return 0, errors.New("No id provided")
	}

	return strconv.Atoi(id)
}

func getInsults(w http.ResponseWriter, r *http.Request) {
	insults, err := friService.GetFrinsults(r.Context(), &frinsultproto.Void{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to retrieve insults: " + err.Error()))
		return
	}

	reply := make([]insult, 0, len(insults.GetInsults()))
	for _, i := range insults.GetInsults() {
		reply = append(reply, *p2a(i))
	}
	replyWithJSON(w, reply, http.StatusOK)
}

func deleteInsultByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	if _, err := friService.DeleteFrinsultByID(r.Context(),
		&frinsultproto.ByIDRequest{ID: int64(id)}); err != nil {
		return
	}

	log.Printf("Deleted insults %d", id)
	w.WriteHeader(http.StatusOK)
}

func getInsultByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	rec, err := friService.GetFrinsultByID(
		r.Context(),
		&frinsultproto.ByIDRequest{ID: int64(id)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to retrieve ID %d: %s", id, err.Error())))
		return
	}

	replyWithJSON(w, rec, http.StatusOK)
}

func putInsult(w http.ResponseWriter, r *http.Request) {
	log.Printf("Added insults !")
	decoder := json.NewDecoder(r.Body)

	entity := insult{}
	if err := decoder.Decode(&entity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request:" + err.Error()))
		return
	}

	log.Printf("Putting %s", entity.String())

	inserted, err := friService.InsertFrinsult(r.Context(),
		a2p(&entity))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode request:" + err.Error()))
		return
	}
	replyWithJSON(w, inserted, http.StatusOK)
}

func updateInsultByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to retrieve ID: " + err.Error()))
		return
	}

	decoder := json.NewDecoder(r.Body)
	entity := insult{}
	if err := decoder.Decode(&entity); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request:" + err.Error()))
		return
	}
	entity.ID = id

	log.Printf("Updating entity %T", entity)
	if _, err := friService.UpdateFrinsult(r.Context(), a2p(&entity)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to update request:" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func voteInsultByID(w http.ResponseWriter, r *http.Request, vote int) {
	id, err := getIDFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, err := friService.VoteFrinsultByID(r.Context(), &frinsultproto.VoteRequest{
		ID:   int64(id),
		Vote: int64(vote),
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to vote for id %d: %s", id, err.Error())))
		return
	}

	log.Printf("Voted %d for insult %d ", vote, id)
	w.WriteHeader(http.StatusOK)
}

func upvoteInsultByID(w http.ResponseWriter, r *http.Request) {
	voteInsultByID(w, r, 1)
}

// TODO: Downvote handler !
