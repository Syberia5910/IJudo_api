FROM golang:alpine3.19 as sources  

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY models/*.go ./models/
COPY config/*.go ./config/
COPY controllers/*.go ./controllers/

RUN CGO_ENABLED=0 GOOS=linux go build -o /ijudo_api

############### Final images without sources

FROM golang:alpine3.19

COPY --from=sources /ijudo_api /ijudo_api

CMD ["/ijudo_api"]