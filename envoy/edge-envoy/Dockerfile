FROM envoyproxy/envoy:latest

RUN apt-get update --fix-missing
RUN apt-get install -y ca-certificates
RUN apt-get install curl --yes

COPY envoy.json /etc/envoy/

EXPOSE 8081

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.json --service-cluster edge