## Local Development

To run the application:

```bash
make up
```

we have a volume set up between our project code and the container. Any change to the code is reflected in the container, and we use CompileDaemon to watch our changes.

## API Testing
API test requests are located in `test/api/api.http`. Use the VS Code REST Client extension to execute these tests.