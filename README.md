# Description

This is a RESTful API on TV Shows built using Go and Gin framework. It does the basic CRUD operations except the update operation.

# Purpose

I built this project while learning how to build APIs in Golang.

# Usage

This is not really to be used by others but if for some reason you want to use it then;

1. Clone the repo using this command `git clone https://github.com/quamejnr/showtime`
2. Naviagate into the repo and run `go run .` command in the root of the directory.

# Endpoints

| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
| List TV Shows | GET | localhost:8080/series/ | None | 200 OK |
| Get TV Show | GET | localhost:8080/series/{id} | id: int | 200 OK |
| Add TV Show | POST | localhost:8080/series/ |  | 201 Created |
| Delete TV Show | DELETE | localhost:8080/series/{id} | id: int | 204 No Content |

## Sample Data
Sample JSON data to test `Add TV Show` endpoint
```json
{
    "id": 3,
    "title": "Stranger Things",
    "genre": "Adventure",
    "seasons": [
        {
            "season": 1,
            "episodes": [
                "s01201",
                "s01e02",
                "s01e03"
                ]
        }
    ]
}
```
Sample JSON response for `List TV Shows` endpoint
```json
[
  {
    "id": 1,
    "title": "Ted Lasso",
    "genre": "Comedy",
    "seasons": [
      {
        "season": 1,
        "episodes": [
          "s01e01",
          "s01e02",
          "s01e03",
          "s01e04",
          "s01e05"
        ]
      },
      {
        "season": 2,
        "episodes": [
          "s02e01",
          "s02e02",
          "s02e03",
          "s02e04",
          "s02e05"
        ]
      }
    ]
  }
]
```

# Future Work

Hoping to link to a database as data is currently being stored in memory.
