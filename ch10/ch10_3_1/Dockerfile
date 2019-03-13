FROM python:3.5-alpine3.4

RUN apk add --no-cache --virtual=build-deps build-base && \
    apk add --no-cache g++ && \
    pip3 install locustio pyzmq && \
    apk del --no-cache build-deps

WORKDIR /locust
COPY senario.py /locust/

ENTRYPOINT [ "/usr/local/bin/locust" ]

EXPOSE 8089 5557 5558
