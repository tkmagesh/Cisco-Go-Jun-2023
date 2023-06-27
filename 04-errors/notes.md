# Error Handling #

- error => any object that implements the "error" interface
- How to create an error?
    - errors.New()
    - fmt.Errorf()
- Other utilities
    - errors.Join()
    - errors.Unwrap()
    - errors.Is()
    - errors.As()