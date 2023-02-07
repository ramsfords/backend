build:
	go mod tidy
	go mod vendor
prod:
	git stash
	git pull origin main
	go mod tidy
	go mod vendor
	go build main.go
	screen ./main serve --http="0.0.0.0:8090"
local:
	git add .
	git commit -m "update"
	git push origin main --force

run:
	go run main.go
makezod:
	ts-to-zod  ../web/src/types/user/v1.ts ../web/src/types/user/userZod.ts