FROM pdf2html:latest AS pdf2html


FROM golang:1.20.2-alpine3.16 AS builder
WORKDIR /opt
COPY go.mod go.sum ./
RUN go mod download
COPY ./pdf2html-lambda ./pdf2html-lambda
RUN GOOS=linux GOARCH=amd64 go build -o main ./pdf2html-lambda/main.go


FROM alpine:3.16
# COPY --from=pdf2html /usr/local/share/poppler /usr/local/share/poppler
# COPY --from=pdf2html /usr/local/share/pdf2html /usr/local/share/pdf2html
# COPY --from=pdf2html /usr/local/bin/pdf2html /usr/local/bin
# COPY --from=pdf2html /usr/local/lib /usr/local/lib
ADD .build/pdf2html.docker.tar .
COPY --from=builder /opt/main /main
# Add Lambda Runtime Interface Emulator and use a script in the ENTRYPOINT for simpler local runs
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
RUN chmod 755 /usr/bin/aws-lambda-rie
COPY ./pdf2html-lambda/entry.sh /
RUN chmod 755 /entry.sh
ENTRYPOINT [ "/entry.sh", "/main" ]
