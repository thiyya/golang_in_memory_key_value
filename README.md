# golang_in_memory_key_value

REST-API service works as an in memory key-value store (like redis)

### Usage

```go
go mod init golang_in_memory_key_value
go mod vendor
go build
go run main.go 

docker : 
docker image build -t thiyya/golang_in_memory_key_value .
docker container run --name letsGo -p 8080:8080 thiyya/golang_in_memory_key_value go run main.go
or 
docker-compose up

// set key
POST _ http://localhost:8080/keys
Body : 
{
  "key": "erhan",
  "value": "xxxx",
  "ttl": 915
}

// Query within the given ttl time
GET _ http://localhost:8080/keys/erhan

// If the service is down, it can reload the key-values thanks to the backup it takes.
// Docker volume added
const CLEANING_INTERVAL : It is the interval for deleting keys whose TTL has expired.
const BACKUP_INTERVAL : It is the interval determined to take a backup.
const DEFAULT_TTL : If TTL is not given, it is TTL given by default.

```
