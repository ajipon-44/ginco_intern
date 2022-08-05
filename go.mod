module github.com/minguu42/myapp

go 1.18

require (
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	local.packages/auth v0.0.0-00010101000000-000000000000
)

require github.com/jinzhu/inflection v1.0.0 // indirect

replace local.packages/auth => ./auth
