FROM    golang:1.17.0-alpine3.14
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --update alpine-sdk
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17
RUN go build -o bin/main .
EXPOSE 8080
CMD [ "bin/main" ]
