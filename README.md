# Directory Structure

- The `bin` directory will contain our compiled application binaries, ready for deployment to a production server.
- The `cmd/api` directory will contain the application-specific code for our Greenlight API application. This will include the code for running the server, reading and writing HTTP requests, and managing authentication.
- The `internal` directory will contain various ancillary packages used by our API. It will contain the code for interacting with our database, doing data validation, sending emails and so on. Basically, any code which isn’t application-specific and can potentially be reused will live in here. Our Go code under cmd/api will import the packages in the internal directory (but never the other way around).
- The `migrations` directory will contain the SQL migration files for our database.
- The `remote` directory will contain the configuration files and setup scripts for our production server.
- The `go.mod` file will declare our project dependencies, versions and module path.
- The `Makefile` will contain recipes for automating common administrative tasks — like auditing our Go code, building binaries, and executing database migrations.


# HTTP Methods
|Method	  | Usage       |
|:-------:|:------------|
|GET	    | Use for actions that retrieve information only and don’t change the state of your application or any data.|
|POST	    | Use for non-idempotent actions that modify state. In the context of a REST API, POST is generally used for actions that create a new resource.|
|PUT	    | Use for idempotent actions that modify the state of a resource at a specific URL. In the context of a REST API, PUT is generally used for actions that replace or update an existing resource.|
|PATCH	  | Use for actions that partially update a resource at a specific URL. It’s OK for the action to be either idempotent or non-idempotent.|
|DELETE	  | Use for actions that delete a resource at a specific URL.|

# API Endopoints
|Method|URLPattern|Handler|Action|
|------|----------|-------|------|
|GET|/v1/healthcheck|healthcheckHandler|Show application information|
|GET|/v1/movies|listMoviesHandler|Show the details of all movies|
|POST|/v1/movies|createMovieHandler|Create a new movie|
|GET|/v1/movies/:id|showMovieHandler|Show the details of a specific movie|
|PUT|/v1/movies/:id|editMovieHandler|Update the details of a specific movie|
|DELETE|/v1/movies/:id|deleteMovieHandler|Delete a specific movie|
