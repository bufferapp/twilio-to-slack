FROM golang:1.7-alpine
MAINTAINER Dan Farrelly <dan@buffer.com>
ADD twilio-to-slack twilio-to-slack
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/twilio-to-slack"]
