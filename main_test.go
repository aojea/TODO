// main_test.go

package main

import (
	"testing"

	baloo "gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("http://localhost:8080")

func TestWrongEndpoint_ShouldReturn404(t *testing.T) {
	test.Get("/test_wrong").
		Expect(t).
		Status(404).
		Done()
}

// Lists API

const schemaLists = `{
  "listId": "1",
  "title": "object",
  "username": "username"
}`

func TestCreateList(t *testing.T) {
	test.Post("/api/v1/user/me/lists").
		JSON(map[string]string{"title": "test", "username": "antonio"}).
		Expect(t).
		Status(201).
		Type("application/json").
		JSONSchema(schemaLists).
		Done()
}

func TestUpdateList(t *testing.T) {
	test.Put("/api/v1/user/me/list/1").
		JSON(map[string]string{"title": "test"}).
		Expect(t).
		Status(200).
		Type("application/json").
		JSONSchema(schemaLists).
		Done()
}

func TestGetLists(t *testing.T) {
	test.Get("/api/v1/user/me/lists").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Done()
}

func TestGetList(t *testing.T) {
	test.Get("/api/v1/user/me/list/2").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Done()
}

func TestDeleteList(t *testing.T) {
	test.Delete("/api/v1/user/me/list/2").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Type("json").
		JSON(`{"result":"success"}`).
		Done()
}

// Tasks API

const schemaTasks = `{
	"taskId": "1",
	"listId": "1",
	"position": "1",
  "title": "task title",
	"description": "task description",
	"completed": "false",
	"tags": "tag1" 
}`

func TestCreateTask(t *testing.T) {
	test.Post("/api/v1/lists/1/tasks").
		JSON(map[string]string{"title": "test", "description": "antonio task"}).
		Expect(t).
		Status(201).
		Type("application/json").
		JSONSchema(schemaTasks).
		Done()
}

func TestUpdateTask(t *testing.T) {
	test.Put("/api/v1/lists/1/task/1").
		JSON(map[string]string{"title": "test", "description": "antonio task"}).
		Expect(t).
		Status(200).
		Type("application/json").
		JSONSchema(schemaTasks).
		Done()
}

func TestGetTasks(t *testing.T) {
	test.Get("/api/v1/lists/1/tasks").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Done()
}

func TestGetTask(t *testing.T) {
	test.Get("/api/v1/lists/1/task/1").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Done()
}

func TestDeleteTask(t *testing.T) {
	test.Delete("/api/v1/lists/1/task/1").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Type("json").
		JSON(`{"result":"success"}`).
		Done()
}
