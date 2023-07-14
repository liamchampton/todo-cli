# todo-cli

A simple command line todo list manager written in Go.

## Getting Started

### Running locally

To run this locally, clone the repository and run the commands below:

```bash
go build
./todo-cli
```

Before running the application, you need to run the command below to create the todo list `.yaml` file:

```bash
touch $HOME/.todo-list.yaml 
```
This will create a `.yaml` file in your home directory. This is where the todo list will be stored.

### Adding a task

To add a task run the command below:

```bash
./todo-cli add task-1 --description "finish the cli project" --deadline "30-07-2023"
```

### Listing tasks

To list all tasks run the command below:

```bash
./todo-cli list
```

### Removing a task

To remove a task run the command below:

```bash
./todo-cli remove task-1
```

### Updating a task

To update a task run the command below:

```bash
./todo-cli update task-1 --description "finish the cli project before the end of next month" --deadline "30-08-2023"
```