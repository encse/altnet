# this version of frotz doesn't compile on the node image we are using, it still works on node:16
# newer versions of frotz have issues with idoregesz...
# https://gitlab.com/DavidGriffith/frotz/-/issues/268
FROM node:16.2 as frotzbuilder 
RUN apt update
RUN apt-get install -y wget curl git make gcc sudo zip
WORKDIR /tmp
RUN git clone https://gitlab.com/DavidGriffith/frotz.git
RUN cd frotz && git checkout 2.50 && make dfrotz 

FROM golang:latest AS gobuilder
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY cmd ./cmd
COPY lib ./lib
RUN go install ./cmd/...

FROM golang:latest AS base
RUN apt update
RUN apt-get install -y libgraph-easy-perl

FROM base AS app
ENV APP_ROOT=/usr/app
ENV DFROTZ /usr/bin/dfrotz 
ENV GRAPH_EASY /usr/bin/graph-easy 
ARG TWITTER_ACCESS_TOKEN
ENV TWITTER_ACCESS_TOKEN=$TWITTER_ACCESS_TOKEN
RUN env
COPY --from=frotzbuilder /tmp/frotz/dfrotz $DFROTZ
COPY --from=gobuilder /go/bin/ ${APP_ROOT}
COPY data ${APP_ROOT}/data
COPY config.yml ${APP_ROOT}

WORKDIR ${APP_ROOT}
CMD ./app

