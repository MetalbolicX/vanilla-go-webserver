# vanilla-go-webserver

By José Martínez Santana

## Technologies used


<p align="center">
	<a href="https://go.dev/" target="_blank" rel="noreferrer"><img  alt="Golang" height="50px" style="padding-right:10px;" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg"/></a>
</p>

## Description

A web server build with the standard library. Created for personal or small project API development and static file sharing.

## Features

- [x] Crete, Read, Update, and Delete (**CRUD**) methods to interact with a database.
- [x] Lightweight and efficient web server implementation using **Go's standard library**.
- [x] Enables easy development of custom APIs for personal projects or any small project.
- [x] Supports serving static files for sharing static content like `HTML`, `CSS`, `JavaScript`, etc.
- [x] Supports the routing with regular expressions validation.
- [x] Capability to add middlewares.
- [x] Supports to load simple `.env` file without external libraries.

# Usage

1. Clone the repository:
``` Bash
git clone https://github.com/MetalbolicX/vanilla-go-webserver.git
```
2. Install Go if you haven't already: [https://golang.org/doc/install](https://golang.org/doc/install)
3. Navigate to the project directory:
```Bash
cd your-project-directory
```
4. Change the variables of the `.env` file.
5. Build the server:
```Bash
go build
```
6. Start the server with the `.exe` file created by the `go build` command.
7. Access the server in your browser: `http://localhost:3000` or change the **port** in the `.env` file.

## Directory structure

The folders structures of the project are described in the following image:

```Bash
├── data
│   ├── external
│   │   └── exercises.db
├── database
│   └── relational-database.go
├── go.mod
├── go.sum
├── handlers
│   ├── customers.go
│   ├── handlers.go
│   ├── home.go
│   └── pages.go
├── main.go
├── middlewares
│   ├── checkauthentication.go
│   └── logging.go
├── models
│   └── types.go
├── repository
│   └── repository.go
├── server
│   ├── router.go
│   └── server.go
├── static
│   ├── css
│   │   └── styles.css
│   └── js
│       └── test.js
├── templates
│   └── index.html
├── types
│   └── types.go
└── utils
    ├── endpoint-identifier.go
    └── envfile-loader.go
```

For custom changes I suggest just to modify or add the next folders:

1. database.
2. handlers.
3. middlewares.

To better understanding, let's see an example.

### Change the database

The current database is a [SQLite3](https://www.sqlite.org/index.html), now it is necessary to scale to a [PostgresSQL](https://www.postgresql.org/) database.

1. In the file `relational-database.go` change the following lines:
```Go
package database

import (
	"context"
	"database/sql"
	"log"
	"time"

  // Erase this line
	// _ "github.com/mattn/go-sqlite3"
)

type relationalDBRepo struct {
	db *sql.DB
}
```
2. In `.env` file change:

|Before|After|
|:---|:---|
|`DB_MANAGEMENT_SYSTEM=sqlite3`|`DB_MANAGEMENT_SYSTEM=postgres`|
|`DB_URL=./data/external/exercises.db`|`DATABASE_URL=postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable` (Check the documentation for the correct implementation of the connection string)|

<ins>NOTE</ins>: If somebody wants to add a `NoSQL` database create another file (Ex. `mongodb.go`) and add the logic to interact with the database.

### Handlers addition

1. In the `handlers` folder add file called `home.go` and add the next code:
```Go
package handlers

import (
	"encoding/json"
	"net/http"
)

type homeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(homeResponse{
		Message: "Welcome to my first web app with Go!!",
		Status:  true,
	})
}
```
2. Add the the new handler route to the server in the `main.go` file in the `bindRoutes` function.
```Go
func bindRoutes(s *server.Server) {
	s.Handle(http.MethodGet, "/home", handlers.HomeHandler)
}
```

### Middleware addition

1. In the `middlewares` folders add a new file.
2. In the new file add the logic of the code. For example:
```Go
package middlewares

import (
	"github.com/MetalbolicX/vanilla-go-webserver/types"
)

func Example() types.Middleware {
	return func(handlerLogic http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Add the logic ...
	}
}
```
3. Add the middleware in the `main.go`. For example:
```Go
func bindRoutes(s *server.Server) {
	s.Handle(http.MethodDelete, "/customer/\\d+",
		s.AddMiddleware(handlers.DeleteCustomerHandler,
			middlewares.CheckAuth(),
			middlewares.Logging(),
			middlewares.Example()))
}
```

# Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. I invite you to collaborate directly in this repository: [vanilla-go-webserver](https://github.com/MetalbolicX/vanilla-go-webserver)

# License

vanilla-go-webserver is released under the [MIT License](https://opensource.org/licenses/MIT).
