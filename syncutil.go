package syncutil

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"sync"
)

// LoggingRWMutex wraps a sync.RWMutex and logs to STDERR on every lock/unlock call.
type LoggingRWMutex struct {
	mu sync.RWMutex

	// Optional name used during print.
	Name string
}

// Lock logs to STDERR that a lock is being obtained and then obtains the lock.
func (rw *LoggingRWMutex) Lock() {
	rw.printStack("obtaining lock", debug.Stack())
	rw.mu.Lock()
	rw.printStack("obtained lock", debug.Stack())
}

// Unlock logs to STDERR that a lock is being released and then releases the lock.
func (rw *LoggingRWMutex) Unlock() {
	rw.printStack("releasing lock", debug.Stack())
	rw.mu.Unlock()
	rw.printStack("released lock", debug.Stack())
}

// RLock logs to STDERR that a read lock is being obtained and then obtains the read lock.
func (rw *LoggingRWMutex) RLock() {
	rw.printStack("obtaining read lock", debug.Stack())
	rw.mu.RLock()
	rw.printStack("obtained read lock", debug.Stack())
}

// RUnlock logs to STDERR that a read lock is being released and then releases the read lock.
func (rw *LoggingRWMutex) RUnlock() {
	rw.printStack("releasing read lock", debug.Stack())
	rw.mu.RUnlock()
	rw.printStack("released read lock", debug.Stack())
}

// RLocker returns the underlying mutex's locker.
func (rw *LoggingRWMutex) RLocker() sync.Locker {
	return rw.mu.RLocker()
}

func (rw *LoggingRWMutex) printStack(action string, stack []byte) {
	lines := strings.Split(string(stack), "\n")
	if rw.Name == "" {
		lines[0] = fmt.Sprintf("%s %s", lines[0], action)
	} else {
		lines[0] = fmt.Sprintf("%s %s for %q", lines[0], action, rw.Name)
	}
	fmt.Fprintln(os.Stderr, strings.Join(lines, "\n"))
}
