package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Task object definition
type task struct {
	ID          int    `json:"taskId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Position    int    `json:"position"`
	Completed   bool   `json:"completed"`
	ListID      int    `json:"listId"`
}

func (t *task) getTask(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, title, description, tags, position, completed, list_id  FROM tasks WHERE id = '%d'", t.ID)
	rows, err := db.Query(statement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Tags, &t.Position, &t.Completed, &t.ListID); err != nil {
			return err
		}
	}
	return nil
}

func (t *task) updateTask(db *sql.DB) error {
	// Note: It doesn't fail if id is not valid
	completed := 0
	if t.Completed {
		completed = 1
	}
	statement := fmt.Sprintf("UPDATE tasks SET title='%s', description='%s', tags='%s', position='%d', completed='%d', list_id='%d' WHERE id='%d'", t.Title, t.Description, t.Tags, t.Position, completed, t.ListID, t.ID)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func (t *task) deleteTask(db *sql.DB) error {
	// Note: It doesn't fail if id is not valid
	statement := fmt.Sprintf("DELETE FROM tasks WHERE id=%d", t.ID)
	_, err := db.Exec(statement)
	return err
}

func (t *task) createTask(db *sql.DB) error {
	completed := 0
	if t.Completed {
		completed = 1
	}
	statement := fmt.Sprintf("INSERT INTO `tasks` (title, description, tags, position, completed, list_id) VALUES('%s', '%s', '%s', '%d', '%d', '%d')", t.Title, t.Description, t.Tags, t.Position, completed, t.ListID)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil
}
func (l *list) getTasks(db *sql.DB) ([]task, error) {
	statement := fmt.Sprintf("SELECT id, title, description, tags, position, completed, list_id  FROM tasks WHERE list_id = '%d'", l.ID)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []task{}
	for rows.Next() {
		var t task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Tags, &t.Position, &t.Completed, &t.ListID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// TasksHandler operates over the TODO tasks /lists/{listId}/tasks
func TasksHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtain the list id from the request
		vars := mux.Vars(r)
		listID, err := strconv.Atoi(vars["listId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid list ID")
			return
		}
		l := list{ID: listID}

		switch r.Method {
		// GET tasks that belong to a list
		case http.MethodGet:
			lists, err := l.getTasks(db)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, lists)
		case http.MethodPost:
			var t task
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&t); err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
			defer r.Body.Close()
			// TODO: check if id payload is different to the URL one and fail
			t.ListID = listID
			if err := t.createTask(db); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusCreated, t)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

// TaskIDHandler operates on tasks on a specific list /lists/{listId}/task/{taskId}
func TaskIDHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtain the list and task id from the request
		vars := mux.Vars(r)
		listID, err := strconv.Atoi(vars["listId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid list ID")
			return
		}
		id, err := strconv.Atoi(vars["taskId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid task ID")
			return
		}

		t := task{ID: id, ListID: listID}

		switch r.Method {
		// GET task
		case http.MethodGet:
			if err := t.getTask(db); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, t)
		// PUT update task
		case http.MethodPut:
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&t); err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid request payload")
				return
			}
			defer r.Body.Close()
			// TODO: URL defined ids take precedence, we should check against the payload and fail if different
			t.ListID = listID
			t.ID = id
			if err := t.updateTask(db); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, t)
		// DELETE delete task
		case http.MethodDelete:
			if err := t.deleteTask(db); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
