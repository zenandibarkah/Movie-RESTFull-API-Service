# Movie-RESTFull-API-Service

## Requirement

1. You have PostgreSQL in your computer or you can use the cloud instead.
2. For this project, we will need some packages. The dependencies that you need:

```go
  go get -u github.com/gin-gonic/gin
  go get github.com/lib/pq
  go get github.com/joho/godotenv
```

- Gin is one of frameworks that most used in Go.
- PQ is package that required to `connect Go to PostgreSQL database`.
- Godotenv is package to read the `.env` file.

## How to run this project

1. You can clone this project.
2. You have to make database `xsis_movies` in the postgres. Run this to terminal:
   ```zsh
    psql -U postgres
    postgres=# CREATE database xsis_movies;
    postgres=# exit
   ```
3. Then you can import this sql file to database with running this command:
   ```zsh
    psql -d <database name> -f postgresql.sql --username=<username postgres>
   ```
4. Or you can run manually to create database, create table, and insert data with `DBeaver`, `PGAdmin`, and so on.
5. You can run:
   ```go
    go mod tidy
   ```
   to install the dependencies.
6. And finally, you can run with
   ```go
    go run main.go
   ```
7. For documentations, you can import `Postman Collections/Movie-RestAPI.postman_collection.json` file to your Postman if you want to test the API.

## The API with following endpoints:

1. List of Movie:
   ```zsh
    Endpoint: /Movies
    Method: GET
    Response: 
            {
                "status": true,
                "message": "OK!",
                "errors": null,
                "data": [
                    {
                        "id": 1,
                        "title": "Pengabdi Setan 2 Comunion",
                        "description": "Sebuah film horor Indonesia tahun 2022 yang disutradarai dan\r\n    ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi\r\n    Setan.",
                        "rating": 7,
                        "image": "",
                        "created_at": "2023-02-17T16:00:00Z",
                        "updated_at": "2023-02-17T16:00:00Z"
                    },
                    {
                        "id": 2,
                        "title": "Pengabdi Setan ",
                        "description": "",
                        "rating": 8,
                        "image": "",
                        "created_at": "2023-02-17T16:00:00Z",
                        "updated_at": "2023-02-18T23:27:43Z"
                    }
                ]
            }
   ```
2. Detail of Movie:
   ```zsh
    Endpoint: /Movies/:id
    Method: GET
    Response: 
            {
                "status": true,
                "message": "OK!",
                "errors": null,
                "data": {
                    "id": 1,
                    "title": "Pengabdi Setan 2 Comunion",
                    "description": "Sebuah film horor Indonesia tahun 2022 yang disutradarai dan\r\n    ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi\r\n    Setan.",
                    "rating": 7,
                    "image": "",
                    "created_at": "2023-02-17T16:00:00Z",
                    "updated_at": "2023-02-17T16:00:00Z"
                }
            }
   ```
3. Add New Movie:
   ```zsh
    Endpoint: /Movies
    Method: POST
    Response: 
            {
                "status": true,
                "message": "OK!",
                "errors": null,
                "data": {
                    "id": 7,
                    "title": "A Man Called Otto",
                    "description": "",
                    "rating": 9,
                    "image": "",
                    "created_at": "2023-02-19T00:27:01Z",
                    "updated_at": "2023-02-19T00:27:01Z"
                }
            }
   ```
4. Update Movie:
   ```zsh
    Endpoint: /Movies/:id
    Method: PATCH
    Response: 
            {
                "status": true,
                "message": "OK!",
                "errors": null,
                "data": {
                    "id": 2,
                    "title": "Pengabdi Setan ",
                    "description": "Chapter 1 of Pengabdi Setan",
                    "rating": 8,
                    "image": "",
                    "created_at": "2023-02-17T16:00:00Z",
                    "updated_at": "2023-02-19T00:28:59Z"
                }
            }
   ```
5. Delete Movie:
   ```zsh
    Endpoint: /Movies/:id
    Method: DELETE
    Response: 
            {
                "status": true,
                "message": "OK!",
                "errors": null,
                "data": {}
            }
   ```