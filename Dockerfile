## We specify the base image we need for our
## go application
FROM golang:1.17-alpine

ENV REPO_URL=github.com/amitabhprasad/bookstore-app/bookstore-items-api

ENV GOPATH=/app

ENV APP_PATH=${GOPATH}/src/${REPO_URL}

ENV WORKPATH=$APP_PATH/src

COPY src ${WORKPATH}

WORKDIR $WORKPATH
RUN go build -o items-api .

# expose port 8084
EXPOSE 8084

CMD ["./items-api"]