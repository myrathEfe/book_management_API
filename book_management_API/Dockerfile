#GO İMAJI
FROM golang:1.21-alpine
#CALISMA DİZİNİ
WORKDIR /app

#MODÜLLERİ KOPYALA VE YÜKLE
COPY go.mod .
COPY go.sum .

#TÜM PROJEYİ KOPYALA

COPY . .

#UYGULAMAYI BUILD ET
RUN go build -o main .

#KONTEYNER CALISTIRILDIGINDA
CMD ["./main"]



#Özetle:
#Temel Go ortamını alıyoruz (FROM)
#Çalışma dizinini belirliyoruz (WORKDIR)
#Modülleri indirip cache avantajı sağlıyoruz (COPY + go mod download)
#Kodumuzu konteynere kopyalıyoruz (COPY . .)
#Go kodunu derleyip binary oluşturuyoruz (RUN go build)
#Konteyner başlatıldığında uygulamayı çalıştırıyoruz (CMD ["./main"])