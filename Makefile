build:
	cd ../firstshipper_backend && git add . && git commit -m "update" || true && git push origin main
	cd ../menuloom_backend && git add . && git commit -m "update" || true && git push origin main
	cd ../configs && git add . && git commit -m "update" || true && git push origin main
	cd ../foundations && git add . && git commit -m "update"|| true && git push origin main
	cd ../services && git add . && git commit -m "update" || true && git push origin main
	go get  github.com/ramsfords/backend/firstshipper_backend@latest
	go get 	github.com/ramsfords/backend/menuloom_backend@latest
	go get  github.com/ramsfords/backend/configs@latest
	go get  github.com/ramsfords/backend/foundations@latest
	go get  github.com/ramsfords/services@latest
	go mod tidy
	go mod vendor
run:
	go run main.go
makezod:
	ts-to-zod  ../web/src/types/user/v1.ts ../web/src/types/user/userZod.ts