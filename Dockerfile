FROM golang:latest

LABEL authors="mouss"
RUN mkdir /src
COPY . /src
WORKDIR /src 
RUN go build -o ascii-art-web-dockerize
CMD ["./ascii-art-web-dockerize"]
EXPOSE 8080