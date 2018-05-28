GCC=go
GCMD=run
GPATH=main.go
migrate-up:
	db-migrate up
migrate-down:
	db-migrate down
run:
	make build
	$(GCC) $(GCMD) $(GPATH)
build:
	make build_db
build_db:
	# rm db_structs.go
	go run pkg/db/main/main.go -json=pkg/db/config.json
	mv db_structs.go pkg/db/