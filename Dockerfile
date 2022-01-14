FROM golang:alpine AS masterapibuilds

WORKDIR /appbuilds

COPY . .

RUN go mod tidy
RUN go build -o binary


FROM alpine:latest as masterapirelease
WORKDIR /app
RUN apk add tzdata
COPY --from=masterapibuilds /appbuilds/binary .
COPY --from=masterapibuilds /appbuilds/.env /app/.env
ENV PORT=7073
ENV DB_USER="sperma"
ENV DB_PASS="asdQWE123!@#"
ENV DB_HOST="128.199.124.131"
ENV DB_PORT="3306"
ENV DB_NAME="db_tot"
ENV DB_DRIVER="mysql"
ENV DB_REDIS_HOST="128.199.124.131:6379"
ENV DB_REDIS_PASSWORD="asdQWE123!@#"
ENV DB_REDIS_NAME="0"
ENV JWT_SECRET_KEY="secrettotomaster"
ENV JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=1800
ENV TZ=Asia/Jakarta



RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT [ "./binary" ]