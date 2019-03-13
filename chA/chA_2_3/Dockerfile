FROM golang:1.9

RUN mkdir /echo
COPY main.go /echo

RUN useradd gihyo 
USER gihyo

CMD ["go", "run", "/echo/main.go"]
