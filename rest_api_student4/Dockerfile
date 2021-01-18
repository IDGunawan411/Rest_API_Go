FROM golang
# WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
ADD . /go/src/belajar_golang/rest_api_student4
WORKDIR /go/src/belajar_golang/rest_api_student4
RUN go get belajar_golang/rest_api_student4
RUN go install
RUN go get github.com/go-sql-driver/mysql
# Build the Go app
CMD ["go", "run", "/go/src/belajar_golang/rest_api_student4/main.go"]