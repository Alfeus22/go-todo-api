# Go-Todo-API 🚀

Project API To-Do List sederhana menggunakan **Golang** dengan sistem autentikasi **JWT** dan database **MongoDB**. Project ini dibuat sebagai bagian dari roadmap pembelajaran Backend Developer Internship.

## ✨ Fitur
- **User Authentication**: Register & Login menggunakan JWT (JSON Web Token).
- **Task Management (CRUD)**:
  - Membuat tugas baru.
  - Melihat daftar tugas (hanya milik user yang login).
  - Update status tugas.
  - Menghapus tugas.
- **Data Isolation**: Menjamin user tidak dapat mengakses, mengubah, atau menghapus data milik user lain.

## 🛠️ Tech Stack
- **Language**: [Golang](https://golang.org/)
- **Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: [MongoDB](https://www.mongodb.com/)
- **Drivers**: `mongo-driver` (v1)
- **Auth**: `golang-jwt`

## 📋 Prasyarat
Sebelum menjalankan project ini, pastikan kamu sudah menginstall:
- Go (versi 1.20+)
- MongoDB (Local atau Atlas)
- Postman (untuk testing API)

## 🚀 Cara Menjalankan
1. Clone repository:
   ```bash
   git clone [https://github.com/username-kamu/nama-repo.git](https://github.com/username-kamu/nama-repo.git)