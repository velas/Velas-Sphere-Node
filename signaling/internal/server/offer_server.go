package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pion/webrtc/v3"
)

func OfferServer(peerConnection *webrtc.PeerConnection) {
	http.HandleFunc("/offer", func(w http.ResponseWriter, r *http.Request) {
		if peerConnection == nil {
			log.Println("no peer connection provided")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		offer := webrtc.SessionDescription{}

		err = json.Unmarshal(body, &offer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = peerConnection.SetRemoteDescription(offer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		answer, err := peerConnection.CreateAnswer(nil)
		if err != nil {
			log.Println("failed to create the answer:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			log.Println("failed to set local description", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		answerBytes, err := json.Marshal(answer)
		if err != nil {
			log.Println("failed to marshal the answer", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(answerBytes)
	})

	err := http.ListenAndServe(":"+"8080", nil)
	if err != nil {
		panic(err)
	}
}
