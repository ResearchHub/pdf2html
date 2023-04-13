FROM alpine:3.16 AS builder
RUN apk add --no-cache libcurl libexecinfo-dev libstdc++ libtool
RUN apk add --no-cache autoconf automake build-base cmake curl make python3
# install nodejs + npm
RUN apk add --no-cache nodejs npm
WORKDIR /var/task
# install aws lambda runtime client for nodejs
RUN npm install aws-lambda-ric
# install aws lambda runtime emulator (for local testing)
RUN mkdir -p ~/.aws-lambda-rie
RUN curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
RUN chmod +x ~/.aws-lambda-rie/aws-lambda-rie
RUN cp ~/.aws-lambda-rie/aws-lambda-rie /usr/local/bin/aws-lambda-rie


FROM alpine:3.16
RUN apk add --no-cache nodejs npm

ENV LANG=en_US.UTF-8
ENV TZ=:/etc/localtime
ENV PATH=/var/lang/bin:/usr/local/bin:/usr/bin/:/bin:/opt/bin:/sbin
ENV LD_LIBRARY_PATH=/var/lang/lib:/lib64:/usr/lib64:/var/runtime:/var/runtime/lib:/var/task:/var/task/lib:/opt/lib
ENV LAMBDA_TASK_ROOT=/var/task
ENV LAMBDA_RUNTIME_DIR=/var/runtime

COPY --from=builder /var/task /var/task
COPY --from=builder /usr/local/bin/aws-lambda-rie /usr/local/bin
# COPY --from=builder /usr/local/lib/node_modules /usr/local/lib
# COPY --from=builder /usr/local/bin/aws-lambda-ric /usr/local/bin

COPY entrypoint.sh  /
RUN chmod +x /entrypoint.sh

WORKDIR /var/task

EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]
