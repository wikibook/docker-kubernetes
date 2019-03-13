FROM alpine:3.7

RUN apk add --no-cache --virtual=build-deps wget && \
    wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 && \
    mv jq-linux64 /usr/local/bin/jq && \
    chmod +x /usr/local/bin/jq && \
    apk del build-deps

COPY show-attr.sh /usr/local/bin/

ENTRYPOINT ["sh", "/usr/local/bin/show-attr.sh"]
CMD [""]
