package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

type IndexedVideo struct {
	Parts [][]string
	Title string
}

func upload(storage map[string]IndexedVideo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// 1000MB
		err := r.ParseMultipartForm(1000 << 20)
		if err != nil {
			// failed to parse the form
			fmt.Println(err)
			return
		}

		file, _, err := r.FormFile("video")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		parts := [][]string{}

		for {
			tempFile, err := ioutil.TempFile("tmp", "upload-*.mp4")
			if err != nil {
				fmt.Println(err)
			}
			defer tempFile.Close()

			buf := make([]byte, 10<<20)

			n, err := io.ReadFull(file, buf)

			tempFile.Write(buf[:n])

			nodeName := ""

			parts = append(parts, []string{nodeName + tempFile.Name()})

			if err == io.ErrUnexpectedEOF {
				break
			}
			if err != nil {
				fmt.Println("error happened:", err)
				return
			}

		}

		id := "1"

		storage[id] = IndexedVideo{
			Parts: parts,
			Title: "My Awesome Video",
		}

		responseJSONBytes, err := json.Marshal(
			map[string]string{
				"id": id,
			},
		)
		if err != nil {
			fmt.Println("failed to marshal json: ", err)
			return
		}

		w.Write(responseJSONBytes)
	}
}

func download(storage map[string]IndexedVideo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		readers := []io.ReadSeeker{}
		for _, part := range storage["1"].Parts {
			partFile, err := os.Open(part[0])
			if err != nil {
				fmt.Println("failed to open part file", err)
				return
			}

			readers = append(readers, partFile)
		}

		multiReadSeeker := MultiReadSeeker(readers...)

		http.ServeContent(w, r, "1", time.Now(), multiReadSeeker)

	}
}

func main() {
	storage := map[string]IndexedVideo{}
	r := chi.NewRouter()
	r.Post("/upload", upload(storage))
	r.Get("/{id}", download(storage))

	http.ListenAndServe(":8080", r)
}
