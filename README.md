# Learning Golang

Welcome to the Golang concepts and notes reference! This documentation covers the fundamental concepts, features, and best practices of Go (Golang) as developed by Google.

## What is Go?

Go is a **statically typed** programming language developed by Google. It captures errors at compile time, leading to more robust software.
> **Philosophy:** Go = C + strings + garbage collection + concurrency

### Key Features
- **Fast Compilation:** No heavy dependency checking.
- **Efficient String Processing:** Designed to process and analyze large amounts of text data (like web pages).
- **Garbage Collection:** Built-in memory management.
- **Concurrency (not just parallelism):** Uses lightweight threads managed by the Go runtime called *Goroutines*.
- **Static Typing**
- **Lightweight & Zero Dependencies**
- **Built-in Support for Testing**
- **Strong Community**

Go is heavily used in world-class scalable systems like **Docker** and **Kubernetes**.

---

## Getting Started

### Important Resources
1. [go.dev](https://go.dev/)
2. [gobyexample.com](https://gobyexample.com/)
3. [A Tour of Go](https://go.dev/tour/welcome/1)
4. [Effective Go](https://go.dev/doc/effective_go)

### Go Workspace Structure
Traditionally, a Go workspace has 3 main folders:
- `bin/`: Executable binaries.
- `pkg/`: Compiled package files.
- `src/`: Source code.

### Starting a New Project
Initialize a Go module in the root folder to manage dependencies:

```bash
go mod init <projectname>
```

This creates a `go.mod` file which represents the root of your Go module. It is responsible for bringing in all the resources required for `main.go`.

To run a program:
```bash
go run main.go
```

**Note on Packages:**
- The concept of a package is closely tied to the directory structure. All `.go` files in a single directory belong to the same package.
- The `main` package is special; it is required for executable programs. The `go run` command looks for the `main()` function in the `main` package to execute everything.

---

## Core Language Concepts

### Slices
Slices are a flexible and dynamic data structure providing a modern alternative to traditional arrays arrays.
- **Dynamic size:** Length can be changed during the program's execution.

### Control Flow
- Go only features **one** looping mechanism: the `for` loop.

### `defer` Keyword
- A code block with the `defer` keyword executes at the end of the surrounding function's lifecycle. It is extremely useful for cleanup (e.g., closing open files or connections).

---

## File Handling (I/O)

We generally read files in 2 ways:
1. **Buffer (`bufio`)**: Reads files in structured chunks of data. Good for large files.
2. **`os.ReadFile` (formerly `ioutil`)**: Loads the *entire file* into memory. Use carefully—if the file size is too large, it can cause memory overflow errors.

---

## Web and HTTP Interfaces

| Feature | `url.URL` | `http.Response` |
| --- | --- | --- |
| **Package** | `net/url` | `net/http` |
| **Represents** | An internet address/identifier | A complete message received from a server |
| **Usage** | Parsing, constructing, and manipulating addresses | Reading the outcome and content of a web request |
| **Key Fields** | `Scheme`, `Host`, `Path`, `RawQuery` | `StatusCode`, `Header`, `Body` |

---

## Concurrency: Goroutines & Sync

Go achieves concurrency native using **Goroutines**.

### The Party Analogy (`sync.WaitGroup`)
- **Goroutines** are like individual tasks that need to be completed concurrently (e.g., setting the table, preparing food, decorating the room).
- **`sync.WaitGroup`** is like your setup checklist for the party. You use it to keep track of tasks that have been completed and to know when everything is done.
  - **`Add()`**: Adding a task to your checklist.
  - **`Done()`**: Checking off a task when it's completed.
  - **`Wait()`**: Waiting for all tasks to be checked off before starting the party.

*Note: While `sync.WaitGroup` is highly recommended, it is not mandatory. You can also use **Channels** to synchronize goroutines.*

---

## Standard Output (`fmt`)

By default, printing standard output goes to the terminal:
```go
fmt.Print("Hello, World!")
```

**Writing to destinations other than terminal:**
You can inject a different `io.Writer` (such as a file `os.File` or a memory buffer `bytes.Buffer`) using `fmt.Fprintf`:
```go
fmt.Fprintf(file, "Hello to file!")
```