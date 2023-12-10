FROM ubuntu

COPY fermat-theorem.go /fermat-theorem.go

RUN apt-get -y update && apt-get install -y apt-transport-https
RUN apt-get -y install golang

CMD go run fermat-theorem.go
