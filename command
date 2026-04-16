go mod init go-mvc-crud
go mod tidy

go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/stretchr/testify
go get github.com/glebarez/sqlite

go run main.go
