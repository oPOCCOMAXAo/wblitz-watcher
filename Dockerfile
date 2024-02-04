FROM alpine:3.18

WORKDIR /

COPY /wbwatcher /wbwatcher

RUN chmod +x /wbwatcher

EXPOSE 8080

ENTRYPOINT [ "/wbwatcher"  ]
