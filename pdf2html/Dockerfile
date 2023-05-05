FROM alpine:3.16 as builder

ENV UNATTENDED="--assume-yes"
ENV DEBIAN_FRONTEND=noninteractive

RUN apk add --no-cache cmake g++ gcc gettext git jq libjpeg-turbo make minify pkgconfig ruby tar wget xz
RUN apk add --no-cache boost-dev cairo-dev fontconfig-dev freetype-dev gettext-dev glib-dev gnu-libiconv-dev libjpeg-turbo-dev libpng-dev libxml2-dev

WORKDIR /home/pdf2html

# copy and build libs
COPY ./libs ./libs
RUN cd ./libs && \
    cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local && \
    make -j $(nproc)

# copy and build binary
COPY ./share ./share
COPY ./src ./src
COPY ./CMakeLists.txt ./CMakeLists.txt
RUN cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local && \
    make -j $(nproc) && \
    make -j $(nproc) install

# determine all of the lib dependencies so they can be
# copied over to the final image
RUN mkdir -p /tmp/pdf2html/lib
RUN ldd /usr/local/bin/pdf2html | awk '{ if ($2 == "=>") print $3; else print $1; }' > deps.txt
RUN xargs -I '{}' cp '{}' /tmp/pdf2html/lib < deps.txt


FROM alpine:3.16

RUN apk add --no-cache fontconfig

COPY --from=builder /usr/local/share/poppler /usr/local/share/poppler
COPY --from=builder /usr/local/share/pdf2html /usr/local/share/pdf2html
COPY --from=builder /usr/local/bin/pdf2html /usr/local/bin
COPY --from=builder /tmp/pdf2html/lib /usr/local/lib

# This is for debugging
# CMD [ "/bin/bash" ]
