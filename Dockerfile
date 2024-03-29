# this version of frotz doesn't compile on the node image we are using, it still works on node:16
# newer versions of frotz have issues with idoregesz...
# https://gitlab.com/DavidGriffith/frotz/-/issues/268
FROM node:16.2 as frotzbuilder 
RUN apt update
RUN apt-get install -y wget curl git make gcc sudo zip
WORKDIR /tmp
ENV VERSION 2e7406ade90ee59ada54f8851a5d1d675a222044
RUN git clone https://gitlab.com/DavidGriffith/frotz.git
# ENV VERSION 2.50
RUN cd frotz && git checkout $VERSION && make dfrotz 

FROM golang:latest AS gobuilder
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd ./cmd
COPY lib ./lib
COPY ent ./ent
COPY schema ./schema
RUN go install -ldflags="-s -w" ./cmd/...

FROM golang:latest AS base
RUN apt update
RUN apt-get install -y libgraph-easy-perl

FROM base AS app
ENV APP_ROOT=/usr/app
ENV DFROTZ /usr/bin/dfrotz 
ENV GRAPH_EASY /usr/bin/graph-easy 
ENV GOBIN ${APP_ROOT}
ARG TWITTER_ACCESS_TOKEN
ENV TWITTER_ACCESS_TOKEN=$TWITTER_ACCESS_TOKEN
RUN env
COPY --from=frotzbuilder /tmp/frotz/dfrotz $DFROTZ
COPY --from=gobuilder /go/bin/ ${APP_ROOT}
COPY config.yml ${APP_ROOT}

WORKDIR ${APP_ROOT}
CMD ./app

