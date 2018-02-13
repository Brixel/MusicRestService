# MusicRestService
A rest service that serves music files from a database. This app is built and explained in the talk on the 
Workshopday2018 of Brixel

We start with a simple "Hello world" example to test our develpment environment.

```Go
fmt.Println("Hello Brixel!")
```

Next we will build a simple http server that serves the hello world message when a user enters a URL. Therefore we 
need a package that can be downloaded with the <b>"go get"</b> command.

```bash
$ go get github.com/julienschmidt/httprouter
```

This will download the httprouter of Julien Schmidt from github to your GOPATH.
