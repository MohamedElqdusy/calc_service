# Calc Service

# Installation
You can install Calc Service by running the following command:
```sh
$ cd calc_service
$ make
```
After running 'make' command you will find the client and the server binaries in
'bin/' directory.

# System requirements
- Go-lang: go1.9.4 linux/amd64
- GNU Make 4.1
- Protoc: libprotoc 3.5.1

# Usage
Runing the server
```sh
$ cd bin
$ ./server
```
Runing the client
```sh
$ cd bin
$ ./client
```
After running both the server and the client, you would use the client to calculate expressions like 
```
5555+840
```
and you will get the response from the server.

# Notes
You would write the expression with or without spaces as
```
    68423   -   840  
```
or
```
12*10
```
