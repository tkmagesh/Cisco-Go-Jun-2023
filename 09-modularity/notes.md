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
    
## Package ##
- Internal organization of a module
