#start from base image
FROM golang:1.16
#configure repo url
ENV REPO_URL=github.com/eremitic/bookstore_oauth_api

ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL
#Copy entire source code from current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o oauth-api .

EXPOSE 8080

CMD ["./oauth-api"]





