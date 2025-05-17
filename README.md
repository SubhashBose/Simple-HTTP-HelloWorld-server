# Simeple HTTP Hello-World webserver
A simple HTTP Hello-World webserver in Go lang, with 
* no dependencies
* small binary
* single commnad line run
* fully safe and secure to expose the port to public

This is aimed to make a simple webserver for network testing, which is secure enough to safely expose the server port to public as it only serves a static (but server set) content to client.  

Binaries are compiled for almost every possible OS and architecture combinations. Pre-compiled binaries can be downloaded from [releases](https://github.com/SubhashBose/Simple-HTTP-HelloWorld-server/releases).

Example: (addr and text arguments are optional)
```
./http-helloworld -addr 8000 -text "Hello World!!"
```
