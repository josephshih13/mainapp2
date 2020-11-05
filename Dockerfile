FROM golang:1.15.3-alpine
WORKDIR /mainapp2
ADD . /mainapp2
RUN cd /mainapp2 && go build -o mainapp2
EXPOSE 9936
ENTRYPOINT ./mainapp2