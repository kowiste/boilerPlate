FROM golang:alpine3.17 AS build
WORKDIR /app
RUN mkdir /root/ssh
RUN apk --update-cache add git openssh-client
RUN echo "StrictHostKeyChecking no" > /etc/ssh/ssh_config
RUN apk add git ca-certificates tzdata gcc libc-dev openssh-client && update-ca-certificates
ENV GOPRIVATE="gitlab.techxonn.com/modules"
RUN GIT_TERMINAL_PROMPT=1
COPY ./ms-company /root/.ssh/id_rsa
RUN mkdir -p -m 0700 ~/.ssh && chmod 600 ~/.ssh/id_rsa && eval `ssh-agent` && ssh-add /root/.ssh/id_rsa
RUN git config --global url."git@gitlab.techxonn.com:prudenzia".insteadOf "https://gitlab.techxonn.com/prudenzia"
RUN git config --global url."git@gitlab.techxonn.com:modules".insteadOf "https://gitlab.techxonn.com/modules"
RUN git clone git@gitlab.techxonn.com:prudenzia/ms-company.git
#RUN go clean -modcache
RUN go env -w GOPROXY=direct; go env -w GOSUMDB=off
RUN go env -w GOPRIVATE="gitlab.techxonn.com/prudenzia/*,gitlab.techxonn.com/modules/*"
RUN unset GOPATH
RUN cd ms-company/src && go install gitlab.techxonn.com/modules/kafka-base
RUN cd ms-company/src && go build -o /app/ms-company/src/main


FROM alpine:3.17
WORKDIR /go/app
COPY --from=build /app/ms-company/src/main /go/app/main
