FROM golang:1.17-alpine3.13

ARG USER=hiroki
# install sudo as root
RUN apk add --update sudo

# add new user
RUN adduser -D $USER \
        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER

USER $USER

WORKDIR /go/src

# Install gcc and Go modules.
COPY ./movieshare_api/go.mod ./
RUN sudo apk add --no-cache alpine-sdk build-base \ 
        && go mod download 
