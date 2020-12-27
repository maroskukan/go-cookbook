# go-world
Learning Go

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
