# go-world

Notes and code examples from learning Go language.

[Introduction](#introduction)
[Installation](#installation)
[Compilation](#compilation)

## Introduction

Go was created to address the following challenges:

1. Efficient compilation (e.g. Python, Java, JavaScript)
2. Efficient execution (e.g. Java, C++)
3. Ease of programming (e.g. Python, JavaScript)

### Language Characteristics

- Strong, static type system
- C-inspired syntax
- Multi-paradigm
  - Procedural
  - Object oriented
- Garbage collection
- Fully compiled
- Rapid compilation
- Single binary output

### Language Values

- Simplicity
  ```go
  // Increment and decrements are statements on Go
  i := 1
  i++
  println(i) // 2
  i++
  println(i) // 3
  // All loops in Go are for-loops
  for i :=0; i < 5; i++ ... // loop with incrementor
  for i < 5 ... // loop till condition
  for ... // infinite loop
  for user := range users ... // loop over collection
  ```
- Network aware and concurent apps
  - **net and net/http packages** - Create web servers using only standard library
  - **goroutines** - Start thousand of concurent tasks with minimal resources
  - **channels** - Safely communicate between concurent tasks
- Out of the box experience
  - **Standard Library** - string manipulation, data compression, file manipulation, network APIs, testing APIs
  - **Go CLI** - project initialization, build, code generation, retrieve dependencies, test, profiling, documentation, report language bugs
- Cross-platform
  - GOOS=windows, GOARCH=amd64
  - GOOS=darwin, GOARCH=amd64
  - GOOS=android GOARCH=arm
- Backward compatibility
  - It is intended that programs written to the Go1 specification will continue to compile and run correctly, unchanged, over the lifetime of that specification
  - There are however exceptions related to security, unspecified behavior, specification errors, bugs

### Use cases

- Web services (backend)
- Web applications (frontend)
- DevOps (Docker, K8s)
- GUI / Thick-client 
- Machine learning


## Installation

1. Download latest binary installation e.g. 1.15.6
   ```bash
   wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
   ```
2. Extract the archive.
   ```bash
   sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
   ```
3. Add /usr/local/go/bin to the PATH evironment variable.
   ```bash
   echo "PATH=$PATH:/usr/local/go/bin >> $HOME/.profile
   ```
4. Apply changes to variable.
   ```bash
   source $HOME/.profile
   ```
5. Verify go installation
   ```bash
   go version
   ```

## Compilation

Once the code is written in golang, you need to use **go** tool with the **build** command to produce a binary file which you can execute. You can also build and run at the same time using **run** command.

```bash
go build helloworld.go
./helloworld

or 

go run helloworld.go
```

