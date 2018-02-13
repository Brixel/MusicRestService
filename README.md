# MusicRestService
A rest service that serves music files from a database. This app is built and explained in the talk on the 
Workshopday2018 of Brixel.

We start with a simple "Hello world" example to test our development environment.

```Go
fmt.Println("Hello Brixel!")
```

Next we will build a simple http server that serves the hello world message when a user enters a URL. Therefore we 
need a package that can be downloaded with the<b>"go get"</b>command.

```bash
$ go get github.com/julienschmidt/httprouter
```

This will download the httprouter of Julien Schmidt from github to your GOPATH.

If we now do a request to the server, we should get a response.

```bash
$ curl http://localhost:3000/hello
Hello Brixel!
```

Of course this is not the content we want to share. We want to share a music library. To do this, we need to build a 
type of which we can share values, something that represents the music library, a<b>model</b>. We will store our models 
in a separate folder, the models folder.

When we create the model, make sure that the name of the type and the properties of the type are spelled with a capital 
letter. Otherwise they won't be visible for other packages (including the main package). 

````Go
type Track struct {
	    Id     string
	    Title  string
	    Artist string
	    Album  string
	    Genre  string
}
````

