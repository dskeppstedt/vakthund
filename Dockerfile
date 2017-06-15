FROM ubuntu:16.04

#Install docker and compose
RUN apt update
RUN apt install -y \
apt-transport-https \
ca-certificates \
curl \
software-properties-common
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) \
stable"

RUN apt update
RUN apt-get install -y docker-ce

RUN curl -L https://github.com/docker/compose/releases/download/1.13.0/docker-compose-`uname -s`-`uname -m` > docker-compose
RUN chmod +x docker-compose
RUN mv docker-compose /usr/local/bin

#Install go
RUN curl https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz > go.tar.gz
RUN tar -C /usr/local -xzf go.tar.gz


ENV GOPATH /usr/local/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


#For vakthund

RUN mkdir -p /go/src/app/
COPY scripts /scripts/
RUN chmod +x /scripts/*.sh
COPY src /go/src/app/
WORKDIR /go/src/app/
RUN go build -o app main.go
CMD ["./app"]
