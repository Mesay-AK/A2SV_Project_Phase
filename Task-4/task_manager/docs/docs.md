# Task Manager Documentation

## Table of Contents

1. [Introduction](#introduction)
2. [installation](#installation)
3. [Features](#features)
4. [Usage](#usage)
    - [Starting the Server](#starting-the-server)
    - [API Endpoints](#api-endpoints)

5. [Dependencies](#dependencies)


## Introduction

The **Task Manager** is a simple web-based application designed to manage tasks. It allows users to create, update, delete, and view tasks through a RESTful API. The application is built using the Go programming language and the Gin web framework.

## Installation

To install and run the Task Manager locally, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Mesay-AK/A2SV_Project_Phase/tree/main/Task-4/task_manager.git
   ```

2. **Navigate to the Project Directory**:
    ```bash
    cd task_manager
    ```
## Features

- **Create a Task**: Add a new task with a title, description, due date, and status.
- **View Tasks**: Retrieve a list of all tasks or view details of a specific task by ID.
- **Update a Task**: Modify the details of an existing task.
- **Delete a Task**: Remove a task from the list by its ID.

## Usage

### Starting the server Use :
```bash 
    ./task_manager
```

### Starting the server Use :

Here is the list of available API endpoints:

- **GET /tasks** : Retrieve all tasks.
- **GET /tasks/:id** : Retrieve a specific task by ID.
- **POST /tasks** : Create a new task.
- **PUT /tasks/** : Update an existing task by ID.
- **DELETE /tasks/** :Delete a task by ID.

### Dependencies
- ***Go***: The Go programming language.
- ***Gin***: A web framework for Go.
- ***log***: Standard logging package.
