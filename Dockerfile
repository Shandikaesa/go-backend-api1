# stage 1 : build (compile kode go)
FROM golang:1.26-alpine AS builder

WORKDIR /app

#Copi file go mod dan go sum dulu ( biar cache efesien)
COPY go.mod go.sum ./
RUN go mod download

#copi semua file go
COPY . .

#build binary
RUN go build -o server . 

#stage 2 : run (jalanakan binary)
FROM alpine:latest

WORKDIR /root/

#copy binary dari stage  1
COPY --from=builder /app/server .

# #copy file .env (biar konfigurasi kebaca)
# COPY --from=builder /app/.env . 

#expose port 8080
EXPOSE 8080

#jalankan binary 
CMD ["./server"]