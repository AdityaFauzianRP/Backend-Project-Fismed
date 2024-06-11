# Backend Project Fismed App

![Go](https://img.shields.io/badge/Go-1.22.4-blue)
![Gin](https://img.shields.io/badge/Gin-1.7.7-green)

## Deskripsi
Project ini adalah bagian dari Backend untuk aplikasi SMART ERP. Backend ini dibuat menggunakan Go dan Gin.

## Prasyarat
Pastikan Anda telah menginstal Go di sistem Anda. Anda dapat mengunduhnya dari [golang.org](https://golang.org/dl/).

## Cara Menjalankan Proyek

### Langkah - Langkah: Mengatur Modul
Jalankan perintah berikut untuk memastikan semua dependensi proyek terunduh dan file `go.mod` serta `go.sum` diperbarui.

```bash or any Terminal
1. go mod tidy
2. go mod vendor
3. go run main.go
4. go build -o {name_build}.exe
5. ./{name_build}.exe
