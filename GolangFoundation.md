# Go Foundations

## Content 
1. [Introduction](##Introduction)
2. [Setting Up Golang](##Setting-Up-Golang)
3. [Types](##Types)

***

## 1. Introduction
Go(Golang) was created by Google in November of 2009 originally to be an amazing procedural language for scalable cloud software. It was built to be simple and fast. 

### Use Cases:
* System Admin scripts
* Web app that serves HTML pages
* File server
* Scripts that deploy code to environments

### Why Go?

* It's simple. The overarching reason behind every feature is readability. 

* It compiles down to binaries so deploying is simple.

* It comes with only the basic structures and methods. Add packages only as needed. This allows the source code to remain light weight. 

* It comes with an auto formatter, so spacing is not an argument on teams. 
* Basic testing is built-in

* Concurrency is made easy for the developer. No thread management needed. The design of communication eliminates race conditions.

* It comes with all the things you need to make a server.

* Version 1 is completely backwards compativle, as promised. 

* Interfaces allow for abstraction, easy testability (through Dependency Injection), and the Adapter Pattern.

* Consisten documentation found at [golang.org/doc](https://golang.org/doc/)

***

## 2. Setting Up Golang
### Learning Objectives:
* Installation of Golang
* Steps to set up Golang
* Go Modes: Go Modules and GOPATH
* Parts of a Go file
* Package, Import and Declaration sections of a Go file
* Building and running a Go file

### Installation 
#### Windows Installation
1. To install Golang on a Windows system, download and install the latest 64-bit Go set for Microsoft Windows OS: [https://golang.org/install/](https://golang.org/doc/install).
2. Follow the instructions on the Go installation program.
3. Run the Command Prompt on your computer by searching for "cmd". Open the command line and type: `"go version"`.
4. One example of the output of `go version` is:
```go
go version go1.14.2 windows/amd64
```

### Linux Installation
#### Step One: Install Go Language
Upagrade to apply the latest security updates on Ubuntu.
```linux
sudo apt-get update
sudo apt-get -y upgrade
```

You need to download the Go binary file. Check out [https://golang.org/install/](https://golang.org/doc/install) for the latest version of go.

Then run: 

```linux
cd /tmp 

wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
```

Download the archive and install it to the desired location on the system such as /usr/local 

```linux
sudo tar -xvf go1.14.2.linux-amd64.tar.gz

sudo mv go /user/local
```

#### Step 2: Set up Go environment
Set up the Go language environment variable so that Go can be used outside the `$GOPATH`.

Add /user/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile(for a system-wide installation) or $HOME/.profile:

```linux
export PATH=$PATH:/user/local/go/bin
```
NOTE: changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

***

## 3. Types*
### Concepts
* What is a built-in type
* The number types
* The string type
* Boolean types

### Intro
Types are how we identify a data type in Go. There are many different types in Go, such as `struct` and `interface`. 

#### Numbers
In go there are two major number types: **Integers** and **Floating point** numbers.

##### Integers
