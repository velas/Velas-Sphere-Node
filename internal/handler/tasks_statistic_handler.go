package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/enum"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/spf13/cast"
	"github.com/syndtr/goleveldb/leveldb"
)

const TaskStatisticPageLimit = 10

func NewGetTasksStatisticHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.DB == nil {
			log.Println("no db provided")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		db := config.DB
		cursorStr, err := getTaskStatisticCursor(r, db)
		if err != nil {
			log.Println("failed to get cursor")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cursor, err := cast.ToInt64E(cursorStr)
		if err != nil {
			log.Println("invalid cursor")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		response := entity.TaskStatisticListResponse{
			Data: make([]entity.TaskStatistic, 0, TaskStatisticPageLimit),
		}

		for i := 0; i < TaskStatisticPageLimit; i++ {
			entry, err := db.Get([]byte(entity.TaskIDFromInt(cursor)), nil)
			if err != nil {
				if err == leveldb.ErrNotFound {
					break
				}
				log.Println("failed to get the task statistic entry")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			var taskStatistic entity.TaskStatistic
			err = json.Unmarshal(entry, &taskStatistic)
			if err != nil {
				log.Println("failed to unmarshal task statistic entry")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			response.Data = append(response.Data, taskStatistic)
			//to get previous task
			cursor--
		}

		respBB, err := json.Marshal(&response)
		if err != nil {
			log.Println("failed to marshal response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(respBB)
	}
}

func getTaskStatisticCursor(r *http.Request, db *leveldb.DB) (string, error) {
	raw := r.URL.Query().Get("cursor")
	if raw == "" {
		lastTask, err := db.Get([]byte(enum.LastTaskKey), nil)
		if err != nil {
			if err == leveldb.ErrNotFound {
				return "", nil
			}
			return "", errors.Wrap(err, "failed to get last task entry")
		}

		var last resources.TaskExecutionRequestResponse
		err = json.Unmarshal(lastTask, &last)
		if err != nil {
			return "", errors.Wrap(err, "failed to unmarshal last task entry")
		}
		return last.Id, nil
	}
	return raw, nil
}
