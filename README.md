## Playing with Go Chi

This repository is a playground for experimenting with Go.


## Local Development

You must have go, docker and make installed.\
All the dev is done inside containers.


To run the application:

```bash
make up
```

we have a volume set up between our project code and the container. Any change to the code is reflected in the container, and we use CompileDaemon to watch our changes.

### generate ent
we use ent framework as a our ORM. \
Before running the app for the first time, we need to:
```bash
go generate ./repositories
```
everytime, we add a new file to repositories/schema, we should run this command.

## API Testing
API test requests are located in `test/api/api.http`. Use the VS Code REST Client extension to execute these tests.