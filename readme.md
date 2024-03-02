This is a simple URL shortener API built using the Gin web framework for Go.

### API Endpoints

- `GET /{id}`: Redirects to the original URL based on the provided ID.
- `POST /shorten`: Returns an ID that refers to the original URL.

### Swagger Documentation

Swagger documentation is available at `/swagger/index.html`.

### Run

1. Build docker Image.
2. Start container with following ENV vars : 
#### SERVER_PORT=
#### DB_USERNAME=
#### DB_PASSWORD=
#### COLLECTION_NAME=