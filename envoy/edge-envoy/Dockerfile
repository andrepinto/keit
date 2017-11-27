FROM envoyproxy/envoy:latest

RUN apt-get update --fix-missing
RUN apt-get install -y ca-certificates
RUN apt-get install curl --yes

RUN mkdir /keit

WORKDIR /keit

COPY envoy.json .

EXPOSE 8081

CMD /usr/local/bin/envoy -c /keit/envoy.json --service-cluster edge