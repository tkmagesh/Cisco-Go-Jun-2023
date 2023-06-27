# Panic & Recovery #

## Panic ##
A state of the application where the application execution cannot proceed further

- Use the **panic()** function to programmatically raise a panic
- Use the **recover()** function to gain access to the error that resulted in a panic
- deferred functions are the appropriate place to use the **recover()** function