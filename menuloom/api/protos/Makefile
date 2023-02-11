bufgen:
	cd item && buf mod update && buf generate && cd ..
	cd restaurant && buf mod update && buf generate && cd ..
	cd user && buf mod update && buf generate && cd ..
	cd menu && buf mod update && buf generate && cd ..
	cd /Users/surendrakandel/projects/packages/auther/packages/proto_gen && protoc-go-inject-tag -input="*.pb.go" || true
run:
	go run main.go
makezod:
	ts-to-zod  ../web/src/types/user/v1.ts ../web/src/types/user/userZod.ts