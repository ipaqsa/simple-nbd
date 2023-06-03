new-disk:
	qemu-img create -f raw ${path}/${disk}.img ${size}G

clean:
	rm -rf build

server-build: clean
	go build -o build/server cmd/server/main.go

client-build: clean
	go build -o build/client cmd/client/main.go

server: server-build
	build/server -c ./config.yml