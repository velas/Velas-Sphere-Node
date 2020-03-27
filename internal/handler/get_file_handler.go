package handler

// func NewGetFileHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if config.DB == nil {
// 			log.Println("no db provided")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		db := config.DB

// 		entry, err := db.Get([]byte(enum.NodeInfoKey), nil)
// 		if err != nil {
// 			log.Println("failed to get the entry")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(entry)
// 	}
// }
