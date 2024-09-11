
# A simple todo cli-app
  This is a small  cli based todo task tracker written in go!!
Installation guide is as follows:
## For UNIX based operating systems:
## Clone the repository:
```bash
git clone https://github.com/amogh-dongre/go-todo-cli.git
```
## CD into the build directory

```bash
cd go-todo-cli/cmd/todo/
```
##Build the binary
``` go
go build -o task
```
##Add the binary file for ease of use(OPTIONAL)
```bash
sudo mv -f task /usr/bin
```
- Note: If you did not follow the above step for safety concerns then you can still execute the command using ./task

##To list all the tasks
```bash
task -list
```

## To add tasks

```bash
task -add <task to be added without brackets>
```
# to mark as complete

```bash
task -complete=<serial number of task to be marked as complete>
```
## to delete task

```bash
task -del=<serial number of task to be deleted>
```
