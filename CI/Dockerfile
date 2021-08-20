FROM golang:1.15

RUN apt-get -qq update && apt-get install -y jq && apt-get install -y sudo

WORKDIR /flare
COPY . /flare

RUN export GOPATH=$HOME/go && chmod +x compile.sh && bash ./compile.sh

RUN chmod +x cmd/export_chain_apis.sh && chmod +x cmd/single.sh

EXPOSE 9660
CMD ["/flare/cmd/single.sh"]
