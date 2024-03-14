FROM scratch

COPY target/ipstack-cli /usr/local/bin/ipstack-cli

ENTRYPOINT [ "/usr/local/bin/ipstack-cli" ]
