# pdf2html

## Overview

This repo contains two packages:

-   **pdf2html:** The pdf2html command line utility, written primarily in C++.
-   **pdf2html-lambda:** A lambda function that calls the pdf2html CLI to convert PDF files from S3, written in Go.

## Requirements

The following tools are required in order to develop and run the CLI and lambda function:

-   [homebrew](https://brew.sh) (Mac only)
-   [docker](https://www.docker.com)

Additionally, `gcc` and `make` are required. This can be installed on a Mac by installing the xcode command line tools, via:

```bash
xcode-select --install
```

The following are also required, but can be installed via homebrew. This can be easily done usingg the `Brewfile` at the root of the repo. If you don't have homebrew or are not on a Mac, then follow the links to go through the install process for each.

-   [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)
-   [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)
-   [Go](https://go.dev) (v1.20+)

To install via the `Brewfile`, run the following command at the root of the repo:

```bash
brew bundle install
```

## Building pdf2html-lambda

The lambda function is written in go and is managed using AWS SAM (serverless application model). The lambbda itself is packaged as a docker container, leveraging the docker container support of AWS lambda. This allows us to build the `pdf2html` CLI utility directly into the final container of the `pdf2html-lambda` function.

Most of the build process is managed using `Make`. You can run the following command to build the lambda, which will build the `pdf2html` CLI utility as well.

```bash
make build
```

**_Note:_** Be sure to rebuild the lambda after every code change to ensure the version of the lambda that is run is always up-to-date.

## Running locally

We use [MinIO](https://min.io) to simulate AWS S3 locally. MinIO runs as a docker container, which is configured via the root level `docker-compose.yml` file.

You can run all of the required backend services by running:

```bash
make start.backend
```

Then you can start the lambda by running:

```bash
make start.lambda
```

You can also invoke the lambda once with a test payload by running:

```bash
make invoke
```
