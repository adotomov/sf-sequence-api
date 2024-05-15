FROM golang:1.22 as builder

COPY . .

RUN go mod download && go mod verify
RUN go build -o sf-sequence-api

FROM golang:1.22

COPY --chown=root:root --from=builder . /bin

CMD [ "/bin/sf-sequence-api" ]
