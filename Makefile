build-malicious:
	cd containers/malicious && \
	docker build -f Dockerfile.bad -t onbuild_malicious . \
	&& cd ../..

build-innocent:
	cd containers/inno && \
	docker build --network=host -f Dockerfile.inno -t onbuild_innocent . \
	&& cd ../..

build-innocent-2:
	cd containers/inno-2 && \
	docker build --network=host -f Dockerfile.inno -t onbuild_innocent_2 . \
	&& cd ../..

# Build from Dockerhub
build-innocent-3:
	cd containers/inno-3 && \
	docker build --network=host -f Dockerfile.inno -t onbuild_innocent_3 . \
	&& cd ../..

server-up:
	docker run --rm --name=upload_server \
		-p 25478:25478 \
		-v $(PWD)/containers/sftp/upload:/var/root \
		mayth/simple-upload-server -token patata /var/root

open-upload:
	cd containers/sftp && \
	rm -r app && \
	tar -xvzf upload/upload.tar.gz && \
	ls -la app

build-go:
	mv malicious-onbuild-demo malicious-onbuild-demo.old ; \
	go build .