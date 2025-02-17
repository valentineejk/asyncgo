db_login:
	psql ${DB_URL}

create-migration:
	migrate create -ext sql -dir database/migrations -seq $(name)

migrate:
	migrate -database ${DB_URL} -path database/migrations up
#	migrate -path database/migrations -database "${DB_URL}" -verbose up


sqlc:
	#generate sql commands
	sqlc generate
	
allow:
	direnv allow

#direnv allow .
#direnv reload
#source .envrc
#echo $DB_URL