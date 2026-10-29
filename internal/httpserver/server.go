package httpserver

import (
	"blockchain/internal/chain"
	shared "blockchain/internal/shared"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

type Message struct {
	BPM int
}

func Run(gen chain.Generator, valid chain.Validator) error {
	mux := makeMuxRouter(gen, valid)
	httpAddr := os.Getenv("ADDR")
	log.Printf("Listening on %v", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter(gen chain.Generator, valid chain.Validator) http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock(gen, valid)).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(chain.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func handleWriteBlock(gen chain.Generator, valid chain.Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var m Message

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&m); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		shared.Mutex.Lock()
		//ensure atomicity when creating new block
		newBlock := gen(chain.Blockchain[len(chain.Blockchain)-1], m.BPM)
		shared.Mutex.Unlock()

		if valid(newBlock, chain.Blockchain[len(chain.Blockchain)-1]) {
			chain.Blockchain = append(chain.Blockchain, newBlock)
			spew.Dump(chain.Blockchain)
		}

		respondWithJSON(w, r, http.StatusCreated, newBlock)

	}
}
