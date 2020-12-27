# go-world

Notes and code examples from learning Go language.

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
- Network aware and concurent apps
- Out of the box experience
- Cross-platform
- Backward compatibility


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

## Building the program

Once the code is written in golang, you need to use **go** tool with the **build** command to produce a binary file which you can execute. You can also build and run at the same time using **run** command.

```bash
go build helloworld.go
./helloworld

or 

go run helloworld.go
```

