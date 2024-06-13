FROM golang:1.22.4-alpine3.20 AS builder
WORKDIR /app

RUN apk add --no-cache make git


RUN git version


COPY . .
RUN go build -buildvcs=false
RUN ls -al
RUN pwd
FROM alpine:3.14

RUN apk add --update tzdata
ENV TZ=Asia/Jakarta

RUN ls -al
RUN pwd

ENV USER=be-go
ENV UID=101
ENV GID=101
ENV HOME=/home/$USER
ENV APP_NAME=be-fismed

RUN set -x ; addgroup -g "$GID" -S "$USER"
RUN adduser \
      --disabled-password \
      -g "$GID" \
      -D \
      -s "/bin/bash" \
      -h "/home/$USER" \
      -u "$UID" \
      -G "$USER" "$USER" && exit 0 ; exit 1

WORKDIR ${HOME}

RUN chown -R -v ${UID}.${GID} ${HOME}

USER ${USER}

RUN mkdir log

COPY --from=builder /app/backend_project_fismed /${HOME}/${APP_NAME}

RUN  pwd
RUN  ls -al

USER root
COPY docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh
RUN chown -R -v ${UID}.${GID} ${HOME}
RUN chmod 777 -R .
RUN ls -al ${HOME}
USER ${USER}
EXPOSE 8080
WORKDIR ${HOME}

CMD ["./docker-entrypoint.sh","./be-fismed"]