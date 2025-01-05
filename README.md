# Cloud FLEx
## Development
### Requirements
- Go
- PostgreSQL
### Running the server locally
- Start PostgreSQL
- Create a `.env` file with the database info and key for signing the JWT tokens like this:
```
DB_HOST=127.0.0.1
DB_USER=postgres
DB_PASSWORD=
DB_NAME=cloud-flex
DB_PORT=5432
DB_TIMEZONE=Asia/Kolkata
API_SECRET=examplesecret
```
- `go run main.go`
