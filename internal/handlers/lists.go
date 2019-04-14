package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type list struct {
	ID     int    `json:"listId"`
	Title  string `json:"title"`
	UserID int    `json:"userId"`
}

func (l *list) getList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (l *list) updateList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (l *list) deleteList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (l *list) createList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getLists(db *sql.DB) ([]list, error) {
	statement := "SELECT id, title FROM lists"
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lists := []list{}
	for rows.Next() {
		var l list
		if err := rows.Scan(&l.ID, &l.Title); err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}
	return lists, nil
}

// ListsHandler operates over the TODO lists
func ListsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			lists, err := getLists(db)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, lists)
		case http.MethodPost:
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

// ListIDHandler operater on a specific list
func ListIDHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			vars := mux.Vars(r)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "ListId: %v\n", vars["listId"])
		case http.MethodPost:
			w.WriteHeader(http.StatusOK)
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
