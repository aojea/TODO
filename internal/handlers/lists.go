package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type list struct {
	ID       int    `json:"listId"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

func (l *list) updateList(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (l *list) deleteList(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM lists WHERE id=%d", l.ID)
	_, err := db.Exec(statement)
	return err
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
		// GET: get all lists that belongs to the user
		case http.MethodGet:
			lists, err := getLists(db)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, lists)
		// POST: create a new list for the user
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
		// Obtain the list id from the request
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["listId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}
		l := list{ID: id}

		switch r.Method {
		// PUT: Update list id
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
		// DELETE: Delete list id
		case http.MethodDelete:
			if err := l.deleteList(db); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
