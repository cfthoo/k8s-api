# k8s-api
>**NOTE**: The is just-for-fun and self learning project

## Prerequisite
* k8s env or minikube



## 
This is an example that illustrates how to use the k0s-api.
You can just replace your own namespace and configmap then run the main file.

There is also some function that can be used under function package.

Have fun.


## API documentation

1. Create task for a todolist
This Create method creates a task
PATH: {url}/todolist
METHOD: POST
REQUEST PAYLOAD: 
    {
        "name": "my first task"
    }
RETURN PAYLOAD:
    {
        "id": 1,
        "name": "my first task",
        "created_by": "105301550950520990207",
        "created_at": "2023-05-01T03:16:57.837083Z",
        "modified_at": "2023-05-01T03:16:57.837083Z"
    }

2. Get tasks for a todolist
This Get method returns all task under a specific user.
PATH: {url}/todolist
METHOD: GET
RETURN PAYLOAD:
    {
        "id": 1,
        "name": "my first task",
        "created_by": "105301550950520990207",
        "created_at": "2023-05-01T03:16:57.837083Z",
        "modified_at": "2023-05-01T03:16:57.837083Z"
    }

3. Get task by id 
This Get method returns a task by id.
PATH {url}/todolist{id}
METHOD: GET
RETURN PAYLOAD:
    {
        "id": 1,
        "name": "my first task",
        "created_by": "105301550950520990207",
        "created_at": "2023-05-01T03:16:57.837083Z",
        "modified_at": "2023-05-01T03:16:57.837083Z"
    }

4. Update task for a todolist
This Update method updates a task
PATH: {url}/todolist
METHOD: PUT
REQUEST PAYLOAD: 
    {
        "id": 1
        "name": "Updating my first task"
    }
RETURN PAYLOAD:
    {
        "id": 1,
        "name": "Updating my first task",
        "created_by": "105301550950520990207",
        "created_at": "2023-05-01T03:16:57.837083Z",
        "modified_at": "2023-05-01T03:20:25.747081Z"
    }

5. Delete task by id 
This Delete method delete a task by id.
PATH {url}/todolist{id}
METHOD: DELETE
RETURN PAYLOAD:
    {}


![Alt text](https://github.com/cfthoo/image/blob/main/up.png)
