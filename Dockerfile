FROM alpine:3.20
WORKDIR /app

COPY app/dist app/dist
COPY babyfood-linux .

ENV PORT "8080"
ENV PLAYGROUND_ENABLED "false"
EXPOSE 8080

USER nobody

ENTRYPOINT [ "/app/babyfood-linux" ]