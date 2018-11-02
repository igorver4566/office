FROM golang:onbuild

ADD . /home
        
WORKDIR /home

EXPOSE 8080