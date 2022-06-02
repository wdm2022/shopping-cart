FROM alpine:3.15 AS builder
ARG TASK_URL=https://github.com/go-task/task/releases/download/v3.12.1/task_linux_amd64.tar.gz
ARG HELM_URL=https://get.helm.sh/helm-v3.9.0-linux-amd64.tar.gz

RUN apk add --no-cache curl git

RUN curl -o task.tgz -L ${TASK_URL} \
    && tar xvzf task.tgz task \
    && mv task /usr/bin/ \
    && rm task.tgz

RUN curl -o helm.tgz -L ${HELM_URL} \
    && tar xvzf helm.tgz linux-amd64/helm \
    && mv linux-amd64/helm /usr/bin/ \
    && rm -rf helm.tgz linux-amd64