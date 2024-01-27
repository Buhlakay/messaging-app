module github.com/Buhlakay/messaging-app/msg-send

go 1.21

toolchain go1.21.6

require github.com/jackc/pgx/v5 v5.5.2 // indirect

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

require github.com/Buhlakay/messaging-app/msg-send/database v1.0.0

replace github.com/Buhlakay/messaging-app/msg-send/database => ./database
