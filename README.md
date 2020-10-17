# go-googleapis
A collection of google API's, a Golang practice project and not production ready

### Introduction
This library is based on an older less complete private python library I made years ago, I just added some more features.
The services in this library where reversed from their respective Android clients.

Do not use this project in production environments unless you know what you are doing.

### Future of this project
I intend to add more API's in the future, whenever I get around to reversing them.
These are on my list immediately:
* Speech-to-text
* Better Text-to-speech

### Usage
You can install this library by using:
```
go get github.com/BRUHItsABunny/go-googleapis
```
This library depends on my HTTP client abstraction (also not production ready yet) :
```
go get github.com/BRUHItsABunny/gOkHttp
```

You can find examples in the `_examples` folder.

You may need to find your own API keys, this library is based on reversed Android clients so most likely you won't want to use this library and settle with the standard Golang SDK for these API's.
