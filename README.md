# Simple HTTP Hello World Web Server

A minimal, dependency-free HTTP server written in Go. This utility is designed for easy network testing and is safe to expose to the public internet.

## Features

- 🚀 **No external dependencies**
- 📦 **Small binary size**
- 🧠 **Small CPU and memory footprint**
- 🛠 **Single command to run**
- 🔒 **Secure by design** – serves only static, predefined content
- 🌍 **Cross-compiled for 14 OS and architecture combinations**

## Use Case

This tool is ideal for simple network testing. It serves a static message (set at server startup) and does not expose any dynamic or client-controlled content, making it safe to use in public or production environments.

## Download

Precompiled binaries are available for a wide range of platforms.  


| OS       | Arch   | Download Link |
|----------|--------|---------------|
| Linux    | AMD64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-amd64) |
| Linux    | ARM64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-arm64) |
| Linux    | ARM    | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-arm) |
| Linux    | i386   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_linux-386) |
| Windows  | AMD64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-amd64.exe) |
| Windows  | ARM64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-arm64.exe) |
| Windows  | ARM    | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-arm.exe) |
| Windows  | i386   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_windows-386.exe) |
| MacOS    | AMD64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_darwin-amd64) |
| MacOS    | ARM64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_darwin-arm64) |
| FreeBSD  | AMD64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-amd64) |
| FreeBSD  | ARM64  | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-arm64) |
| FreeBSD  | ARM    | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-arm) |
| FreeBSD  | i386   | [Download](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases/latest/download/http-helloworld_freebsd-386) |


👉 You may also visit the [Releases Page](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases) to download the binaries and M5D checksums.


## Example Usage

```bash
./http-helloworld -addr 8000 -text "Hello, World!!"
````

* `-addr`: (optional) IP:Port address or Port number to bind the server to (default: `8080`)
* `-text`: (optional) Response text to serve (default: `"Hello, World!"`)
* `-log`:  (optional) Turn on console logging all requests received. 
