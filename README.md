# MusicRestService
A rest service that serves music files from a database. This app is built and explained in the talk on the 
Workshopday2018 of Brixel.

We start with a simple "Hello world" example to test our development environment.

```Go
fmt.Println("Hello Brixel!")
```

Next we will build a simple http server that serves the hello world message when a user enters a URL. Therefore we 
need a package that can be downloaded with the **"go get"** command.

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
type of which we can share values, something that represents the music library, a **model**. We will store our models 
in a separate folder, the models folder.

When we create the model, make sure that the name of the type and the properties of the type are spelled with a capital 
letter. Otherwise they won't be visible for other packages (including the main package). 

```Go
type Track struct {
	Id     string
	Title  string
	Artist string
	Album  string
	Genre  string
}
```

Then we create a new route for the router to serve our model.

```Go
r.GET("/tracks/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := models.Track{
		Title:"Song 2",
		Artist: "Blur",
		Album: "Blur",
		Genre: "Rock",
		Id: "1",
	}
 
 	tj, err := json.Marshal(t)
	if err!=nil {
		fmt.Println("Error while trying to marshal track:", err)
	}
 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", tj)
})
```

We change the json names of the properties of the model because in the javascript conventions, fields are written  with 
lower case.

 ```Go
type Track struct {
	Id     string	`json:"id"`
	Title  string	`json:"title"`
	Artist string	`json:"artist"`
	Album  string	`json:"album"`
	Genre  string	`json:"genre"`
}
 ```
 
Now we add a POST and DELETE route, just like we did the GET route.
