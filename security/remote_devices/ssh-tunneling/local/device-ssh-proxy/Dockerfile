# Copyright (C) 2020-2022 Intel Corporation
# 
# SPDX-License-Identifier: Apache-2.0

FROM alpine:latest

RUN apk add --no-cache --update dumb-init openssh-client && rm -rf /var/cache/apk/*

COPY docker-entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT [ "docker-entrypoint.sh" ]
