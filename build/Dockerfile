FROM golang:alpine3.17 AS build
RUN mkdir /root/ssh
RUN apk --update-cache add git openssh-client

WORKDIR /app
COPY ./ms-dashboards .
RUN echo "StrictHostKeyChecking no" > /etc/ssh/ssh_config
RUN apk add git ca-certificates tzdata gcc libc-dev openssh-client && update-ca-certificates
ENV GOPRIVATE="gitlab.com/kowiste"
RUN GIT_TERMINAL_PROMPT=1
COPY /keys/kowiste /root/.ssh/id_rsa
RUN chmod 400 /root/.ssh/id_rsa
RUN mkdir -p -m 0700 ~/.ssh && eval `ssh-agent` && ssh-add /root/.ssh/id_rsa
RUN git config --global url."git@gitlab.com:kowiste".insteadOf "https://gitlab.com/kowiste"
RUN go env -w GOPROXY=direct; go env -w GOSUMDB=off
RUN go env -w GOPRIVATE="gitlab.com/kowiste/*"
RUN unset GOPATH
RUN go build

FROM alpine:3.17
# copy the app form the previous image
COPY --from=build /app/ms-dashboards /go/app/ms-dashboards
EXPOSE 5030 5030
WORKDIR /go/app
CMD ["/go/app/ms-dashboards"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]
