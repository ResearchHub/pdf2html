.PHONY: build.pdf2html
build.pdf2html:
	docker build -o - pdf2html > .build/pdf2html.docker.tar

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
		"s3_input": { "bucket_name": "pdf-uploads", "object_key": "s41551-022-00989-w.pdf" }, \
		"s3_output": { "bucket_name": "pdf-uploads", "object_key": "s41551-022-00989-w.html" } \
	}' | sam local invoke Pdf2HtmlFunction --event -
