.PHONY: build.pdf2html
build.pdf2html:
	docker build -o - pdf2html > .build/pdf2html.docker.tar -t pdf2html:latest

.PHONY: build.pdf2html_x86
build.pdf2html_x86:
	docker buildx build --platform linux/amd64 -o - pdf2html > .build/pdf2html.docker.tar -t pdf2html:latest

.PHONY: build
build: build.pdf2html
	sam build

.PHONY: start.backend
start.backend:
	COMPOSE_PROFILES=backend docker compose up -d

.PHONY: start.lambda
start.lambda: build
	sam local start-lambda

.PHONY: invoke
invoke: build
	echo '{ \
		"s3_input": { "bucket_name": "test", "object_key": "2021.10.25.465725v1.full.pdf" }, \
		"s3_output": { "bucket_name": "test", "object_key": "test.html" } \
	}' | sam local invoke --env-vars locals.json Pdf2HtmlFunction --event -

#echo '{"s3_input": { "bucket_name": "test", "object_key": "2021.10.25.465725v1.full.pdf" }, "s3_output": { "bucket_name": "test", "object_key": "test.html" }}'
