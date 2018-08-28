# Version : 0.0.1
FROM ubuntu
MAINTAINER IgorBondartsov <bondartsov.igor@gmail.com>

RUN apt-get -qq update
RUN apt-get -qq -y install curl
RUN curl -O https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz
RUN tar -zxvf go1.9.1.linux-amd64.tar.gz

ENV GOPATH=$HOME/go
ENV PATH=$PATH:/go/bin:$GOPATH/bin
ENV SRC_DIR=go/src/github.com/IhorBondartsov/OLX_Parser/website

RUN mkdir -p $SRC_DIR
ADD . $SRC_DIR
WORKDIR $SRC_DIR

EXPOSE 8000

CMD ["sh", "start_script.sh"]
