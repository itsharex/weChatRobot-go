FROM golang:1.17.2 as build

WORKDIR /src/weChatRobot-go

COPY . ./

RUN rm -rf bin/ || true && \
    make weChatRobot

FROM alpine:3.15

USER root

WORKDIR /weChatRobot-go

COPY --from=build /src/weChatRobot-go/bin/weChatRobot_* ./weChatRobot

EXPOSE 8080

ENTRYPOINT ["/weChatRobot-go/weChatRobot"]
CMD ["-p", "8080"]
