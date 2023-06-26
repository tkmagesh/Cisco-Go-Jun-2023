# Go Foundations #

## Instructor ##
- Magesh Kuppan
- tkmagesh77@gmail.com
- 99019-11221
- https://linkedin.com/in/tkmagesh

## Schedule ##
- Commence      : 09:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hr)
- Tea Break     : 03:00 PM (20 mins)
- Wind up       : 05:00 PM

## Software Requirements ##
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Chrome Browser

### Software Requirements - Verification ###
- \> go version

## Methodology ##
- No powerpoints
- 100% code driven class
- Code from the class IS the deliverable

## Repository ##
- https://github.com/tkmagesh/cisco-go-jun-2023

# Who Go? #
1. Ease of use
    - Simpler than C & C++

2. Performance
    - Compiled to NATIVE CODE
    - Supports cross platform compilation
    - Performance comparable to C++

3. Less resources
    - No need for Runtime (JVM / CLR)
    - No need for Web/App servers
    
4. Concurrency Support
    - Built in Scheduler (N:M scheduler where N = # of goroutines and M = # of OS threads)
    - A concurrent operation is represented as a goroutine
    - A goroutine will ONLY need a min of 4 KB of memory
    - The builtin scheduler may schedule multiple goroutines to be scheduled for execution using the same OS thread
    - Concurrency support is built in the language
        - go keyword, channel data type (chan), channel operator (<-), range, select-case constructs
    - SDK API support
        - sync package
        - sync/atomic package

