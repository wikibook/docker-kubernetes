FROM alpine:3.7

WORKDIR /
ENV GOPATH /go

# ① 빌드 시에만 필요한 라이브러리 및 도구 설치
RUN apk add --no-cache --virtual=build-deps go git gcc g++

# ② 실행 시에도 필요한 라이브러리 및 도구 설치
RUN apk add --no-cache ca-certificates

# ③ todoapi를 빌드해 실행 파일을 만듬
COPY . /go/src/github.com/gihyodocker/todoapi
RUN go get github.com/go-sql-driver/mysql
RUN go get gopkg.in/gorp.v1
RUN cd /go/src/github.com/gihyodocker/todoapi && go build -o bin/todoapi cmd/main.go
RUN cd /go/src/github.com/gihyodocker/todoapi && cp bin/todoapi /usr/local/bin/

# ④ 빌드 시에만 필요한 라이브러리 및 도구 제거
RUN apk del --no-cache build-deps

CMD ["todoapi"]
