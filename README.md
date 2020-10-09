# syncutil

A collection of utility functions for Go synchronization.


## LoggingRWMutex

The `LoggingRWMutex` implements the same methods as `sync.RWMutex` but logs
the stack trace before each operation.