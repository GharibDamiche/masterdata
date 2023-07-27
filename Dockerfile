FROM alpine
RUN apk --no-cache add tzdata ca-certificates
ADD server /server
ADD config/kore.yaml /kore.yaml
ADD config/kore_dev.yaml /kore_dev.yaml
ENTRYPOINT [ "/server" ]