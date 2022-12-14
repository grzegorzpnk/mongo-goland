FROM nicolaka/netshoot

WORKDIR /mongoclient

USER root

COPY --chown=root ./bin/mongoclient ./mongoclient
#COPY --chown=root ./src/cmd/config.json ./


RUN chmod +x ./mongoclient

ENTRYPOINT ["./mongoclient"]