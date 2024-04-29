FROM ubuntu:latest
LABEL authors="noel"

ENTRYPOINT ["top", "-b"]