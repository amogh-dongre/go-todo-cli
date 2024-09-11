
# A simple todo cli-app
  This is a small  cli based todo task tracker written in go!!
Installation guide is as follows:
## For UNIX based operating systems:
## Clone the repository:
```bash
git clone https://github.com/amogh-dongre/go-todo-cli.git
```
```bash
cd go-todo-cli/cmd/todo/
```

``` go
go build -o task
```

```bash
sudo mv -f task /usr/bin
```
#To list all the tasks
```bash
task -list
```

# TO add tasks

```bash
task -add <task to be added without brackets>
```
# to mark as complete

```bash
task -complete=<serial number of task to be marked as complete>
