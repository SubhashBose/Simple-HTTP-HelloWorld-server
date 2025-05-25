# Simple HTTP Hello World Web Server

A minimal, dependency-free HTTP server written in Go. This utility is designed for easy network testing and is safe to expose to the public internet.

## Features

- 🚀 **No external dependencies**
- 📦 **Small binary size (~5 MiB)**
- 🧠 **Small CPU and memory footprint**
- 🛠 **One-line command to run - no configuration file**
- 🔒 **Secure by design** – serves only static, predefined content
- 🌍 **Cross-compiled for 15 OS and architecture combinations**

## Use Case

This tool is ideal for simple network testing. It serves a static message (set at server startup) and does not expose any dynamic or client-controlled content, making it safe to use in public or production environments. It can be put in Docker containers and used for health checks. 

## Download

Precompiled binaries are available for a wide range of platforms.  


| OS       | Architecture   | Download Link |
|----------|--------|---------------|
| Linux    | AMD 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-amd64) |
| Linux    | i386 32-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-386) |
| Linux    | ARM 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-arm64) |
| Linux    | ARM 32-bit   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-arm) |
| Linux    | RISC-V 64-bit   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-riscv64) |
| Windows  | AMD 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-amd64.exe) |
| Windows  | i386 32-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-386.exe) |
| Windows  | ARM 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-arm64.exe) |
| Windows  | ARM 32-bit   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-arm.exe) |
| MacOS    | Apple Silicon 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_darwin-arm64) |
| MacOS    | Intel 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_darwin-amd64) |
| FreeBSD  | AMD 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-amd64) |
| FreeBSD  | i386 32-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-386) |
| FreeBSD  | ARM 64-bit  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-arm64) |
| FreeBSD  | ARM 32-bit   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-arm) |


👉 You may also visit the [Releases Page](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases) to download the binaries and M5D checksums.


## Example Usage

```bash
./http-helloworld -addr 8000 -text "Hello, World!!"
````

* `-addr`: (optional) IP:Port address or Port number to bind the server to (default: `8080`)
* `-text`: (optional) Response text to serve (default: `"Hello, World!"`)
* `-log`:  (optional) Turn on console logging all requests received. 
