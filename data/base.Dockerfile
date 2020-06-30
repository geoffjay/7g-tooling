FROM debian:buster
MAINTAINER Geoff Johnson <geoff.jay@gmail.com>

# Python setup
RUN apt-get update -qq \
    && apt-get install --no-install-recommends -qq -y \
        python3 \
        python3-pip

# Javascript setup
RUN apt-get update -yq \
    && apt-get install curl gnupg -yq \
    && curl -sL https://deb.nodesource.com/setup_12.x | bash \
    && apt-get install nodejs -yq

RUN npm install -g yarn

RUN mkdir src
WORKDIR src