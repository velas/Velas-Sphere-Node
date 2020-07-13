package server

import (
	"io/ioutil"
	"net/http"
)

func SignalingServer() {
	http.HandleFunc("/offer", func(w http.ResponseWriter, r *http.Request) {
		// TODO: parse the offer?
		// TODO: node selection?
		answerRequest, err := http.NewRequest("POST", "http://127.0.0.1:8081/answer", r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := http.DefaultClient.Do(answerRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		answerBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()

		w.Write(answerBytes)
	})

	err := http.ListenAndServe(":"+"8080", nil)
	if err != nil {
		panic(err)
	}
}
