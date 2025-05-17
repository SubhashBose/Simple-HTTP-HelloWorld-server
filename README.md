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
👉 Visit the [Releases Page](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases) to download the appropriate version for your system.

## Example Usage

```bash
./http-helloworld -addr 8000 -text "Hello, World!!"
````

* `-addr`: (optional) Port number to bind the server to (default: `8080`)
* `-text`: (optional) Response text to serve (default: `"Hello, World!"`)
