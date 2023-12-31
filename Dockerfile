FROM ubuntu

COPY fermat_theorem.go /fermat_theorem.go

RUN apt-get -y update && apt-get install -y apt-transport-https
RUN apt-get -y install golang
RUN go run fermat_theorem.go
