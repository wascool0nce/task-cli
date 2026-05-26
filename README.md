# task-cli

A small command-line task tracker written in Go.

The application stores tasks in a local `tasks.json` file in the current working directory. If the file does not exist, it is created automatically on the first command that reads or writes tasks.

Project idea: https://roadmap.sh/projects/task-tracker

## Features

- Add new tasks
- List all tasks
- Filter tasks by status
- Update task descriptions
- Delete tasks
- Mark tasks as in progress
- Mark tasks as done
- Persist tasks in JSON format

## Requirements

- Go 1.25.7 or newer

## Run

Run the CLI directly from the repository:

```bash
go run ./cmd/app <command> [arguments]
```

Example:

```bash
go run ./cmd/app add "Buy groceries"
```

## Build

Build a local binary:

```bash
go build -o task-cli ./cmd/app
```

Then run it:

```bash
./task-cli <command> [arguments]
```

## Commands

### Add a Task

```bash
task-cli add "Buy groceries"
```

Output:

```text
Task added successfully (ID: 1)
```

New tasks are created with the `todo` status.

### List Tasks

```bash
task-cli list
```

Output:

```text
Tasks: 1
----------------------------------------
Task ID: 1
Description: Buy groceries
Status: todo
Created: 2026-05-26 09:00:00 +0300 MSK
Updated: 2026-05-26 09:00:00 +0300 MSK
----------------------------------------
```

### Filter Tasks

```bash
task-cli list todo
task-cli list in-progress
task-cli list done
```

Supported statuses:

- `todo`
- `in-progress`
- `done`

### Update a Task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

Output:

```text
Task with ID 1 was updated
```

Updating a task changes its `updatedAt` value.

### Delete a Task

```bash
task-cli delete 1
```

Output:

```text
Task with ID 1 was deleted
```

### Mark a Task as In Progress

```bash
task-cli mark-in-progress 1
```

Output:

```text
Task with ID 1 status was updated
```

### Mark a Task as Done

```bash
task-cli mark-done 1
```

Output:

```text
Task with ID 1 status was updated
```

## Data Format

Tasks are stored in `tasks.json` as an array of objects:

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "createdAt": "2026-05-26T09:00:00+03:00",
    "updatedAt": "2026-05-26T09:00:00+03:00"
  }
]
```

Each task has:

- `id`: unique task identifier
- `description`: short task description
- `status`: one of `todo`, `in-progress`, or `done`
- `createdAt`: task creation time
- `updatedAt`: last task update time

## Errors

The CLI returns a non-zero exit code and prints an error message to stderr when a command is invalid.

Examples:

```text
no command provided
unknown command
task description is required
task ID must be a number
unknown task status: blocked
```

## Development

Run all packages:

```bash
go test ./...
```
