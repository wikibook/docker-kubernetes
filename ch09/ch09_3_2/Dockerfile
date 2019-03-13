FROM alpine:3.7 

RUN apk add --no-cache wget 
RUN wget https://github.com/progrium/entrykit/releases/download/v0.4.0/entrykit_0.4.0_Linux_x86_64.tgz 
RUN tar -xvzf entrykit_0.4.0_Linux_x86_64.tgz 
RUN rm entrykit_0.4.0_Linux_x86_64.tgz
RUN mv entrykit /bin/entrykit
RUN chmod +x /bin/entrykit
RUN entrykit --symlink
