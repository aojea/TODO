package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type list struct {
	ID       int    `json:"listId"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

func (l *list) getList(db *sql.DB) error {
	return db.QueryRow("SELECT title FROM lists WHERE id=$1",
		l.ID).Scan(&l.Title)
}

func (l *list) updateList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (l *list) deleteList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (l *list) createList(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO lists(title, user_id) VALUES('%s', (SELECT id FROM users WHERE username = '%s'))", l.Title, l.Username)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&l.ID)
	if err != nil {
		return err
	}
	return nil
}
func getLists(db *sql.DB) ([]list, error) {
	statement := "SELECT l.id, l.title, u.username FROM lists l INNER JOIN users u ON u.id = l.user_id"
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lists := []list{}
	for rows.Next() {
		var l list
		if err := rows.Scan(&l.ID, &l.Title, &l.Username); err != nil {
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
		// GET
		case http.MethodGet:
			lists, err := getLists(db)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, lists)
		// POST
		case http.MethodPost:
			var l list
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&l); err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
			defer r.Body.Close()

			if err := l.createList(db); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusCreated, l)

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
