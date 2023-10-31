## Defer Keyword

The `defer` keyword in Go is used to schedule a function call to be executed at the end of the surrounding function's scope, just before the function returns. It's a powerful and useful feature in Go, and there are several common use cases for it:

1. **Resource Cleanup:** One of the most common use cases for `defer` is resource cleanup. You can defer the closing of files, database connections, or network sockets to ensure they are properly closed, even if an error occurs within the function. For example:

   ```go
   func readFile(filename string) ([]byte, error) {
       file, err := os.Open(filename)
       if err != nil {
           return nil, err
       }
       defer file.Close() // Ensure the file is closed when this function returns

       // Read and process the file
       // ...

       return data, nil
   }
   ```

2. **Unlocking Mutexes:** When working with Goroutines and synchronization primitives like `sync.Mutex`, you can use `defer` to unlock mutexes, ensuring that they are released even if an error occurs within the critical section:

   ```go
   var mu sync.Mutex

   func foo() {
       mu.Lock()
       defer mu.Unlock() // Ensure the mutex is unlocked when this function returns
       // Critical section
   }
   ```

3. **Logging and Profiling:** `defer` can be used for logging or profiling, ensuring that a log entry or profiling data is recorded just before the function exits, regardless of how it exits (normal return, panic, or an error). For example:

   ```go
   func someFunction() {
       defer func() {
           if r := recover(); r != nil {
               log.Printf("Recovered from panic: %v", r)
           }
       }()

       // Rest of the function
   }
   ```

4. **Database Transactions:** When working with databases, you can use `defer` to commit or roll back transactions. This ensures that the transaction is correctly handled, whether the function exits normally or due to an error:

   ```go
   func performDatabaseTask(db *sql.DB) error {
       tx, err := db.Begin()
       if err != nil {
           return err
       }
       defer func() {
           if err != nil {
               tx.Rollback()
           } else {
               tx.Commit()
           }
       }()

       // Database operations
       // ...

       return nil
   }
   ```

5. **Clean-Up and Finalization:** `defer` can also be used for general clean-up and finalization tasks. For example, you might want to defer closing a network connection, cleaning up temporary files, or updating statistics just before a function exits.

The `defer` keyword allows you to write more robust and readable code by ensuring that necessary actions are taken before a function returns, regardless of how it exits. It's a valuable feature for managing resources and handling various scenarios in Go programs.