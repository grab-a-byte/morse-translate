package main

import (
	"encoding/json"
	"fmt"
	"io"
	"morse/morse"
	"net/http"

	"github.com/google/uuid"
)

type request struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type response struct {
	Message     string `json:"message"`
	RequestId   string `json:"requestId"`
	Translation string `json:"translation"`
	Type        string `json:"type"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/Translate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "/api/Translate requires a post request", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request from body", http.StatusBadRequest)
			return
		}
		var req request
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Unable to read body of request", http.StatusBadRequest)
			return
		}

		if req.Type == "MorseCode" {
			result, err := morse.DecryptMorse(morse.MorseCode(req.Message))
			if err != nil {
				errMsg := fmt.Sprintf("Failed to complete request with error %s", err)
				http.Error(w, errMsg, http.StatusBadRequest)
				return
			}

			reqId := uuid.New()
			response := response{
				Message:     req.Message,
				RequestId:   reqId.String(),
				Translation: result,
				Type:        req.Type,
			}

			resJson, err := json.MarshalIndent(response, "", " ")
			if err != nil {
				panic("I give up")
			}
			w.Write(resJson)
			return
		} else if req.Type == "Text" {
			result, err := morse.EncryptMorse(req.Message)
			if err != nil {
				errMsg := fmt.Sprintf("Failed to complete request with error %s", err)
				http.Error(w, errMsg, http.StatusBadRequest)
			}

			reqId := uuid.New()
			response := response{
				Message:     req.Message,
				RequestId:   reqId.String(),
				Translation: string(result),
				Type:        req.Type,
			}

			resJson, err := json.MarshalIndent(response, "", " ")
			if err != nil {
				panic("I give up")
			}
			w.Write(resJson)
			return
		}

		http.Error(w, "Unable to handle request", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", mux)
}
