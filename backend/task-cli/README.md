# Task Tracker CLI
A simple command-line task manager written in Go that helps you track your tasks, their status, and progress.

## Features

- Add, update, and delete tasks
- Mark tasks as "todo", "in-progress", or "done"
- List all tasks or filter by status
- Persistent storage in JSON format
- Simple command-line interface

## Installation

1. Ensure you have Go installed (version 1.16 or higher recommended)
2. Clone this repository or download the source code
3. Build the application:
```bash
go build -o task-cli
```
4. Move the executable to a directory in your PATH (optional but recommended):
```bash
sudo mv task-cli /usr/local/bin/
```

## Usage

### Basic Commands

```bash
task-cli [command] [arguments]
```
### Commands

#### Add a new task:

```bash
task-cli add "Task description"
```
#### Update a task:

```bash
task-cli update [task-id] "New description"
```

#### Delete a task:

```bash
task-cli delete [task-id]
```
#### Change task status:

```bash
task-cli mark-todo [task-id]        # Mark as todo
task-cli mark-in-progress [task-id] # Mark as in-progress
task-cli mark-done [task-id]        # Mark as done
```

#### List tasks:

```bash
task-cli list               # List all tasks
task-cli list todo          # List todo tasks
task-cli list in-progress   # List in-progress tasks
task-cli list done          # List completed tasks
```

## Data Storage

Tasks are stored in a JSON file named ```tasks.json``` in the current working directory. The file is automatically created when you run the application for the first time.

## Task Structure

Each task has the following properties:

```id```: Unique identifier (auto-incremented)
```description```: Task description
```status```: Current status (todo/in-progress/done)
```createdAt```: Creation timestamp
```updatedAt```: Last update timestamp

## Examples

1. Add a new task:
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```
2. Mark task as in progress:
```bash
task-cli mark-in-progress 1
# Output: Task 1 marked as in-progress
```
3. List all in-progress tasks
```bash
task-cli list in-progress
# Output: [list of in-progress tasks]
```
## Error Handling

The application provides descriptive error messages for:

- Invalid commands
- Missing or incorrect arguments
- Non-existent task IDs
- File system errors

## Building from Source

1. Clone the repository:
```bash
git clone https://github.com/your-repo/task-tracker-cli.git
cd task-tracker-cli
```
2. Build the executable:
```bash
go build -o task-cli
```
3. Run tests (if available):
```bash
go test
```
## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is open-source and available under the MIT License.
