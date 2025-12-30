# TaskTracker (task-cli)

TaskTracker is a simple cross-platform command-line task manager written in Go.  
It stores tasks locally on your machine in a human-readable JSON format and provides basic task management commands such as add, list, update, delete, and mark.

The project is designed with a clean separation between:
- command handling
- file I/O
- data models
- output formatting

This makes the codebase easy to extend and reason about.

---

## Features

- Add tasks with descriptions
- List tasks in a clean, column-based layout
- Filter tasks by status (todo, done, in-progress)
- Update task descriptions
- Mark tasks as todo, done, or in-progress
- Delete tasks by ID
- Human-readable timestamps (Today, Yesterday, etc.)
- Colored status output in supported terminals
- Cross-platform (Windows, Linux, macOS)

---

## Requirements

- Go 1.20 or newer (recommended)
- A terminal that supports ANSI colors

Check Go installation:

```sh
go version
```

---

## Project Structure (High Level)

```
TaskTracker/
├── go.mod
├── task-cli.go          // main entry point
├── models/              // Task and status definitions
├── tasks/               // add, list, delete, update, mark commands
├── utility/
│   ├── fileio/          // storage and persistence
│   └── taskprinter/     // formatted output
```

---

## How Data Is Stored

All tasks are stored in a single JSON file called `data.json` inside a hidden directory in the user’s home directory.

### Linux / macOS

```
~/.task-tracker/data.json
```

### Windows

```
C:\Users\Username\.task-tracker\data.json
```

Notes:
- The directory is created automatically on first run
- The file contains a JSON array of task objects
- No external database is used

---

## Running the Project (Development)

From the project root (where `go.mod` exists):

```sh
go run .
```

---

## Compiling the Project (Build Executable)

```sh
go build -o task-cli
```

Run:

```sh
./task-cli list
```

On Windows:

```sh
task-cli.exe list
```

---

## Available Commands

### Add
```sh
task-cli add <task description>
```

### List
```sh
task-cli list
task-cli list todo
task-cli list done
task-cli list in-progress
```

### Update
```sh
task-cli update <task-id> <new description>
```

### Mark
```sh
task-cli mark <task-id> <status>
```

Valid statuses: todo, done, in-progress

### Delete
```sh
task-cli delete <task-id>
```

---

## How It Works

- Load tasks from disk
- Modify tasks in memory
- Save tasks back to disk atomically

This design favors correctness and simplicity.

---
