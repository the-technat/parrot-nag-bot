############################
# STEP 1 build executable binary
############################
# Using go version 1.16.5 to build application
FROM golang:1.16.5 as builder

# /go is WORKDIR of golang, we specify that a little bit 
WORKDIR ${GOPATH}/src/git.technat.ch/technat/parrot-nag-bot

# Get Application in build container
COPY . . 

# Build time! (with zero dependencies)
RUN set -x && \ 
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -o ./parrot-bot-nag 

############################
# STEP 2 build a small image
############################
# Compose actuall image
FROM scratch

# Some Labels
LABEL maintainer "Technat technat@technat.ch"
LABEL src "https://git.technat.ch/technat/parrot-nag-bot"

# Don't run as root
USER 1000

# Get a workdir
WORKDIR /app

# Get static binary in the container
COPY --from=builder ${GOPATH}/src/git.technat.ch/parrot-nag-bot . 

# Expose REST API
EXPOSE 80

# Simply run the bot without any args provided
ENTRYPOINT ["/app/parrot-nag-bot"]

