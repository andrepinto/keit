FROM envoyproxy/envoy:latest

RUN apt-get update --fix-missing
RUN apt-get install -y ca-certificates
RUN apt-get install curl --yes

RUN mkdir /keit

COPY main /usr/local/bin/main
RUN chmod +x /usr/local/bin/main

WORKDIR /keit

COPY envoy.json .

EXPOSE 3000

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]