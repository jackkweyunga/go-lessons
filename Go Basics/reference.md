Using the commands : 

creating/initializing a new module in go
```cmd
 go mod init example/hello
```

after creating the hello.go file and filling it with some code, you will need to run it.
```cmd
go run .
```

In case you added another external module in your code, you need to make sure it is installed by running the code below.
```cmd
go mod tidy
```

wow that was basic.

Online reference: https://go.dev/doc/tutorial/getting-started