# README

## Database 

### _Migration Create_
command: `migrate create -ext sql -dir ./migrations {FILE_NAME}`

### _Migration Up_
command: `migrate -path ./migrations -database postgres:"//localhost:5432/merchant_app?sslmode=disable" up` 

## _Obs:_
  "Start struct attribute names with capital letter if not go will consider private."