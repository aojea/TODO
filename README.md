# "TODO" REST API

RESTful API for a TODO application.

## Features

* User authentication
* Tasks priorization
* Tag filtering

## API REFERENCE

Base URL: todo/v1

### Resources types

#### Lists

Each user can create and delete TODO lists

##### Resources

```
{
  id: string,  # TODO list identifier
  title: string, # Title of the TODO list
}
```

##### Methods

list GET /user/lists Returns all the authenticated user's task lists.

get GET /user/lists/{listId} Returns the authenticated user's specified task list.

insert POST /user/lists Creates a new task list for the authenticated user

update PUT /user/lists/{listId} Updates the authenticated user's specified task list.

delete DELETE /users/lists/{listId} Deletes the authenticated user's specified task list.

patch PATCH /users/lists/{listId} Updates the authenticated user's specified task list. This method supports patch semantics.

#### Tasks

Each TODO list is composed by tasks that can be created, updated, deleted and reorganized by users

##### Resources

```
{
  id: string,  # list tasks identifier
  title: string, # Title of the TODO list tasks
  description: string, # Description of the tasks
  position: int, # Position in the list of the tasks
  completed: boolean, # Flag indicated if the task was done
  tag: string, # comma separated list to tag the tasks
}
```

##### Methods

list GET /lists/{listId}/tasks Returns all the tasks in the specified list
    Allow to query tasks by tag
    tasks?query=tag%20eq%20'test'

get GET /lists/{listId}/tasks/{taskId} Returns the specified task

insert POST /lists/{listId}/tasks/ Creates a new tasks in the the specified list

update PUT /lists/{listId}/tasks/{taskId} Updates tthe specified task

delete DELETE /lists/{listId}/tasks/{taskId}  Deletes the specified tasks

patch PATCH /lists/{listId}/tasks/{taskId} Updates the specified task