build:
	go build -o bin/v1/main main.v1.go
	go build -o bin/v2/main main.v2.go

docker-v1:
	env GOOS=linux go build -o bin/v1/main main.v1.go
	docker build -t andrepinto/keit:1.0.1 -f Dockerfile.v1 .

docker-v2:
	env GOOS=linux go build -o bin/v2/main main.v2.go
	docker build -t andrepinto/keit:2.0.1 -f Dockerfile.v2 .

docker-v1-envoy:
	env GOOS=linux go build -o bin/v1/main main.v1.go
	cp bin/v1/main envoy/v1
	docker build -t andrepinto/keit-envoy:1.0.1 -f envoy/v1/Dockerfile envoy/v1/

docker-v2-envoy:
	env GOOS=linux go build -o bin/v2/main main.v2.go
	cp bin/v2/main envoy/v2
	docker build -t andrepinto/keit-envoy:2.0.1 -f envoy/v2/Dockerfile envoy/v2/

docker-edge-envoy:
	docker build -t andrepinto/keit-edge-envoy:1.0.1 -f envoy/edge-envoy/Dockerfile envoy/edge-envoy/

run-edge-envoy:
	docker run -p 8081:8081 andrepinto/keit-edge-envoy:1.0.1 

run-keit-envoy-1:
	docker run -p 80:80 -p 8001:8001 andrepinto/keit-envoy:1.0.1 

run-keit-envoy-2:
	docker run -p 80:80 -p 8001:8001 andrepinto/keit-envoy:2.0.1 
		
.PHONY: build docker run