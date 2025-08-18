# Book Management API

Bu proje, **Go (Gin framework)** ve **PostgreSQL** kullanarak oluşturulmuş basit bir kitap yönetim API’sidir. Docker ve Docker Compose ile çalışacak şekilde yapılandırılmıştır.

---

## 🛠 Teknolojiler

- Go 1.24
- Gin Web Framework
- GORM (ORM)
- PostgreSQL 15
- Docker & Docker Compose

---

## 📂 Proje Yapısı

book_management_API/
├── controllers/ # API controller’ları
├── initializers/ # DB ve env yükleme
├── models/ # GORM modelleri
├── main.go # Uygulama giriş noktası
├── Dockerfile # Docker image yapılandırması
├── docker-compose.yml # Docker Compose yapılandırması
├── wait-for-db.sh # DB hazır olana kadar bekleme scripti
├── .env # Çevresel değişkenler

---

## ⚙ Kurulum ve Çalıştırma

1. Projeyi klonlayın:

```bash
git clone https://github.com/myrathEfe/book_management_API/
cd book_management_API
.env dosyasını oluşturun:
DB_USER=efe
DB_PASSWORD=123
DB_NAME=bookdb
DB_HOST=db
DB_PORT=5432
Docker Compose ile uygulamayı başlatın:
docker-compose up --build
PostgreSQL container’ı bookdb olarak çalışacak.
API container’ı book-api olarak çalışacak.
API varsayılan olarak http://localhost:8080 üzerinde çalışır.
🚀 API Kullanımı
Kitap Ekleme
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "Gödel, Escher, Bach", "author": "Douglas Hofstadter", "published_year": 1979}'
Tüm Kitapları Listeleme
curl http://localhost:8080/books
ID’ye Göre Kitap Getirme
curl http://localhost:8080/books/1
Kitap Güncelleme
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Gödel, Escher, Bach", "author": "Douglas Hofstadter", "published_year": 1980}'
Kitap Silme
curl -X DELETE http://localhost:8080/books/1


📌 Notlar
* Database container’ı hazır olana kadar wait-for-db.sh scripti çalışır.
* .env dosyası mutlaka olmalı, aksi halde uygulama başlamaz.
* API Gin debug modunda çalışır. Production için GIN_MODE=release ayarlayın.
