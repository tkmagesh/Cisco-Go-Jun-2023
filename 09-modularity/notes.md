# Modules & Packages #

## Module ##
- Any code that need to be versioned and distributed together
- Folder should contain the **go.mod** file
- module file (go.mod)
    - name of the module
    - minimum version of the go runtime targetted

    - Create a module 
        > go mod init <module_name>

- To Run a module
    > go run .

- To build a module
    > go build .
    > go build -o <binary_name> .

- To add a 3rd party package
    > go get <package_name>
    > ex : go get github.com/fatih/color
    - By default downloaded to **$GOPATH$/pkg/mod** folder

- To update the go.mod file
    > go mod tidy

- To download the dependencies
    > go mod download

- To upgrade a dependency
    > go get -u <package_name>

- To upgrade all the dependencies
    > go get -u

- To localize the dependencies
    > go mod vendor
    
- Other useful commands
    > go mod graph
    > go mod why <module_name>

- Reference
```
    https://go.dev/ref/mod
```

## Package ##
- Internal organization of a module
