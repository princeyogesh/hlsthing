FROM ubuntu

WORKDIR /home/app

COPY . .

RUN apt update -y
RUN apt upgrade -y
RUN apt install golang-go -y


