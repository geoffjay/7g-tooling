# Frontend
FROM 7geese-dev-base
MAINTAINER Geoff Johnson <geoff.jay@gmail.com>

RUN mkdir 7geese
COPY $SG_PATH/repository 7geese
WORKDIR 7geese

RUN yarn install

EXPOSE 8000
CMD ["yarn run"]
