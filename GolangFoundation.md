# Go Foundations

## Content 
1. [Introduction](##Introduction)
2. [Setting Up Golang](##Setting-Up-Golang)
3. [Types](##Types)

***

## __1. Introduction__
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
go version go1.14.4 windows/amd64
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

## __3. Types__
### Concepts
* What is a built-in type
* The number types
* The string type
* Boolean types

### Intro
Types are how we identify a data type in Go. There are many different types in Go, such as `struct` and `interface`. 

***

#### Numbers
In Go there are two major number types: **Integers** and **Floating point** numbers.

#### Integers
Integers are non-decimal based numbers that can be represented using one of the following types:

* The `int` types can represent any positive or negative non-decimal number
* The `uint` types can only be positive non-decimal numbers

*Sized Types:*

* `int8`, `int16`, `int32`, and `int64`
* `uint8`, `uint16`, `uint32`, and `uint64`

The 8, 16, 32, and 64 by each type represents the maximum size of the type in bits. `int8` has a max size of 8 bits and `uint64` has a max size of 64 bits. 

Examples:

* Given a 64 bit processors `int` will default to `int64`
* Given a 32 bit processors `int` will default to `int32`

#### Float

Floating point numbers are decimal numbers. Floating point numbers are represented by either `float32` or `float64`. Like `int`, the 32 and 64 represent the bit capacity of the floating point number.

There are also `complex64` and `complex128` types that fall under floating point that we will not cover here.

#### Using Number Literals

A literal is a value written as is, and not assigned to any variables.

It is very common to use the following operators with numbers when doing various calculations. 

|Operator    | Description        |
|:----------:|:-------------------|
| +          | addition           |
| -          | subtraction        |
| *          | multiplication     |
| /          | division           |
| %          | remainder/Modulus  |

Example of using number literals and operators

```Golang
package main 

import "fmt"

func main() {
    fmt.Println(200 - 1 + 3 - 4)
    fmt.Println(-10 + 4)
    fmt.Println(1.2 + 3.4)
}
```

#### Rune and Byte

The `rune` and `byte` types are also technically numbers as well and are alias to types covered above. Their actual use will be discussed in following sections. For now its good to note that `rune` is the same as an `int32` in type. It is used to represent UTF-8 values of string literals. 

Similarly `byte` is an alias for `unint8`. Byte is most commonly seen in when working with strings, which in Golang, are an array of bytes. 

#### __Strings__

String types are used to represent text. They are made up of a series of characters. More specifically, strings are made up of `bytes`, where there is typically one byte per character. 

*__Creating Byte Literals__*

There are two ways to create string literals. A string literal can be created by surrounding the characters with double quites `"` or "back ticks" `.

*__Double Quotes__*

* Example: `"Get Going"`
* Will translate escape escape characters. For eample `\n` is translated to a newline.

*__Back Ticks__*

* **Example:** ``Get Going``
* Will not translate escape characters.
* Handy when creating literals that contain double quotes, such as when working with JSON strings. 

*__Concatenated Strings__*

You often need to combine one more or more strings at some point. The quickest, most common way to do this is to use the `+` operator to combine the two strings. 

**Example:** ``` "Go on " + with your bad self." ```

There are other methods of concating strings that may be more efficient when working with large strings, or gorups of strings; however, using the `+` is perfectly fine in most cases. The other mehods will not be covered in this course, but if you are interested in some of the other methods [this is a good blog presenting](https://www.hermanschaaf.com/efficient-string-concatenation-in-go/) some benchmarking and techniques. 

