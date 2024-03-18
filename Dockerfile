FROM scratch

ARG TARGETOS
ARG TARGETARCH

COPY target/builds/ipstack-cli-$TARGETOS-$TARGETARCH /usr/local/bin/ipstack-cli

ENTRYPOINT [ "/usr/local/bin/ipstack-cli" ]
