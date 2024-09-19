### Golang Web Server and Parsing Static File

This project demonstrate a basic golang server creation. Parsing and Routing between mutiple static pages using Golang's core package(net/http).

## Main commands


#### creating server :
```
http.ListenAndServer(":PORT",nil)
```

#### Handling routes
```
http.HandleFunc("/path",routeHandlerFunctionName)
```

#### Parsing Static Html pages
```
template.ParseFiles("/dir/path/file.html")
template.ParseGlob("/dir/path/*.html)
```

#### Render HTML pages and pass data to them
```
Execute(responseWriter, nil)
ExecuteTemplate(responseWriter,"index.html","pass vale to static page")
```