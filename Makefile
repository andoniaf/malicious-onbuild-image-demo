.PHONY: build-malicious
## Build local malicious image
build-malicious:
	cd containers/malicious && \
	docker build -f Dockerfile.bad -t onbuild_malicious . \
	&& cd ../..

.PHONY: build-innocent
## Build client image from local malicious image
build-innocent:
	cd containers/inno && \
	docker build --network=host -f Dockerfile.inno -t onbuild_innocent --no-cache . \
	&& cd ../..

.PHONY: build-innocent-2
## Build client from local malicious image + use.dockerignore
build-innocent-2:
	cd containers/inno-2 && \
	docker build --network=host -f Dockerfile.inno -t onbuild_innocent_2 --no-cache . \
	&& cd ../..

.PHONY: build-innocent-3
## Build client from Dockerhub malicious image + use.dockerignore
build-innocent-3:
	cd containers/inno-3 && \
	docker build --network=host -f Dockerfile.inno -t onbuild_innocent_3 --no-cache . \
	&& cd ../..

.PHONY: server-up
## Start SFTP local server
server-up:
	docker run -d --name=russian_sftp_server \
		-p 25478:25478 \
		-v $(PWD)/containers/sftp/upload:/var/root \
		mayth/simple-upload-server -token patata /var/root

.PHONY: open-upload
## Shows SFTP content
open-upload:
	cd containers/sftp && \
	rm -r app && \
	tar -xvzf upload/upload.tar.gz && \
	ls -la app

.PHONY: build-go
## Build go binary
build-go:
	mv malicious-onbuild-demo malicious-onbuild-demo.old ; \
	go build .

.PHONY: run
## Run all demo steps
run: build-malicious
	clear && $(PWD)/malicious-onbuild-demo --all

.PHONY: run-fast
## Run all demo steps (without typewritter animation)
run-fast: build-malicious
	clear && $(PWD)/malicious-onbuild-demo -i --all


# Plonk the following at the end of your Makefile
.DEFAULT_GOAL := show-help

# From https://gist.github.com/klmr/575726c7e05d8780505a
.PHONY: show-help
show-help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)"
	@echo
	@sed -n -e "/^## / { \
		h; \
		s/.*//; \
		:doc" \
		-e "H; \
		n; \
		s/^## //; \
		t doc" \
		-e "s/:.*//; \
		G; \
		s/\\n## /---/; \
		s/\\n/ /g; \
		p; \
	}" ${MAKEFILE_LIST} \
	| LC_ALL='C' sort --ignore-case \
	| awk -F '---' \
		-v ncol=$$(tput cols) \
		-v indent=19 \
		-v col_on="$$(tput setaf 6)" \
		-v col_off="$$(tput sgr0)" \
	'{ \
		printf "%s%*s%s ", col_on, -indent, $$1, col_off; \
		n = split($$2, words, " "); \
		line_length = ncol - indent; \
		for (i = 1; i <= n; i++) { \
			line_length -= length(words[i]) + 1; \
			if (line_length <= 0) { \
				line_length = ncol - indent - length(words[i]) - 1; \
				printf "\n%*s ", -indent, " "; \
			} \
			printf "%s ", words[i]; \
		} \
		printf "\n"; \
	}' \
	| more $(shell test $(shell uname) == Darwin && echo '--no-init --raw-control-chars')
