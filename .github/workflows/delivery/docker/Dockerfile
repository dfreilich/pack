FROM alpine:3.10

ARG PACK_VERSION
ENV PACK_URL https://github.com/buildpack/pack/releases/download/v${PACK_VERSION}/pack-v${PACK_VERSION}-linux.tgz
RUN echo "${PACK_URL}"
RUN env
RUN echo "hell world ${PACK_URL}"
RUN wget -O- "${PACK_URL}" | tar -C /usr/local/bin -xz pack
ENTRYPOINT ["/usr/local/bin/pack"]
