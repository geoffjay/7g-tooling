# Backend
FROM 7geese-dev-base
MAINTAINER Geoff Johnson <geoff.jay@gmail.com>

ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

RUN pip install --upgrade pip

RUN mkdir 7geese
COPY $SG_PATH/repository 7geese
WORKDIR 7geese
RUN pip install -r requirements.txt \
    && ./manage.py makemigrations \
    && ./manage.py migrate

EXPOSE 8080
CMD ["./manage.py runserver"]
