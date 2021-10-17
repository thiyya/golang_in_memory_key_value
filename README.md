# golang_in_memory_key_value

In memory key-value store olarak çalışan bir REST-API servisidir.

### Usage

```go
go mod init golang_in_memory_key_value
go mod vendor
go build
go run main.go 

// key set edilir 
POST _ http://localhost:8080/keys
Body : 
{
  "key": "erhan",
  "value": "xxxx",
  "ttl": 915
}

// Verilen ttl süresi içinde sorgulanır
GET _ http://localhost:8080/keys/erhan

// Eğer servis down olursa aldıgı backup sayesinde key-value degerlerini reload edebilir.
const CLEANING_INTERVAL : TTL i biten keyleri silme intervalidir.
const BACKUP_INTERVAL : Backup almak için belirlenen intervaldir.
const DEFAULT_TTL : Eger TTL verilmediyse default olarak verilen TTL dir.

```