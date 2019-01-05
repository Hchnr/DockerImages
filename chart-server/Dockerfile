FROM registry.cn-beijing.aliyuncs.com/shannonai/golang:v2.0.0 AS builder

ENV PROJ_NAME="jiweil/chart-server"
ARG SSH_KEY

RUN mkdir -p /root/go/src/$PROJ_NAME/
COPY . src/$PROJ_NAME
RUN $GOPATH/bin/download_dependency.sh $SSH_KEY

# ---
FROM registry.cn-beijing.aliyuncs.com/shannonai/debian
LABEL maintainer="hechenrui123@gmail.com"

ARG HELM_VERSION="v2.12.1"

RUN apt-get update -y \
    && apt-get install -y wget \
    && apt-get install -y git \
    && rm -rf /var/lib/apt/lists/*

RUN set -ex \
    && cd /usr/local \
    && wget https://code.aliyun.com/k9kdqvbb/files/raw/master/helm-$HELM_VERSION-linux-amd64.tar.gz \
    && tar -zxvf helm-$HELM_VERSION-linux-amd64.tar.gz \
    && cp linux-amd64/helm /usr/local/bin/helm \
    && rm helm-$HELM_VERSION-linux-amd64.tar.gz && rm -rf linux-amd64

WORKDIR /root
COPY --from=builder /root/go/bin/chart-server chart-server
COPY entrypoint.sh entrypoint.sh
COPY update.sh update.sh
COPY conf conf

RUN mkdir logs \
    && chmod 755 -R logs \ 
    && chmod u+x chart-server \
    && chmod u+x entrypoint.sh \
    && chmod u+x update.sh \
    && chown -R root: *

EXPOSE 9421

ENTRYPOINT ["./entrypoint.sh"]
CMD ["chart-server"] # set default arg for entrypoint
