FROM golang:latest

##buat folder APP
RUN mkdir /articleapp

##set direktori utama
WORKDIR /articleapp

##copy seluruh file ke completedep
ADD . /articleapp

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]