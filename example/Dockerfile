FROM golang:1.8

MAINTAINER Leonardo Esparis <ljesparis@gmail.com>

WORKDIR "/opt"

ADD .docker_build/gostagram_webapp_example /opt/bin/gostagram_webapp_example
ADD ./templates /opt/templates
ADD ./static /opt/static
ADD .env /opt/bin

ENV HOST gostagram.herokuapp.com
ENV SCHEME https

CMD ["/opt/bin/gostagram_webapp_example"]

