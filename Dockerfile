FROM alpine:edge AS build
RUN apk update
RUN apk upgrade
RUN apk add --no-cache go gcc g++
WORKDIR /
COPY . .
RUN CGO_ENABLED=1 go build /


FROM alpine:edge
WORKDIR /
RUN wget https://download.oracle.com/otn_software/linux/instantclient/185000/instantclient-basic-linux.x64-18.5.0.0.0dbru.zip && \
    unzip instantclient-basic-linux.x64-18.5.0.0.0dbru.zip && \
    apk add --no-cache libaio libnsl libc6-compat gcc && ln -s /usr/lib/* /instantclient_18_5 && \
    cd /instantclient_18_5 && \
    ln -s libnsl.so.2 /usr/lib/libnsl.so.1 && \
    ln -s /lib/libc.so.6 /usr/lib/libresolv.so.2 && \
    ln -s /lib/libc.musl-x86_64.so.1 /usr/lib/ld-linux-x86-64.so.2
ENV LD_LIBRARY_PATH=/instantclient_18_5
COPY --from=build /apiserver /apiserver
ENTRYPOINT [ "/apiserver" ]
