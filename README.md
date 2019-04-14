# "TODO" REST API

RESTful API for a TODO application.

## System Design

This applications is meant to be used in a microservices architecture. It´s stateless so it can scale out just adding more application in parallel and load balancing the requests.

It uses a datastore, in this case a MariaDB database, so the data can  persist.

User management is out of the scope of this project, this implies deal with password management and meet security requirements and there are different projects much more mature that can provide it.

TODO app assumes that there is a user table in the database and use the authentication headers to link the TODO lists to the authenticated user.

## Features

* User management is out of the scope of this project, this implies deal with password management and meet security requirements and there are different approaches with a much mature
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

list GET /user/me/lists Returns all the authenticated user's task lists.

get GET /user/me/lists/{listId} Returns the authenticated user's specified task list.

insert POST /user/me/lists Creates a new task list for the authenticated user

update PUT /user/me/lists/{listId} Updates the authenticated user's specified task list.

delete DELETE /user/me/lists/{listId} Deletes the authenticated user's specified task list.

patch PATCH /user/me/lists/{listId} Updates the authenticated user's specified task list. This method supports patch semantics.

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

## References

[Go: Authorization API](https://auth0.com/docs/quickstart/backend/golang/01-authorization)


## Testing

1. Deploy a database

docker run -p 3306:3306 --name todo-db -e MYSQL_ROOT_PASSWORD=password -d mariadb:10.3