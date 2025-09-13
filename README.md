# 🔐 SSO_BE_API

Sistem **Single Sign-On (SSO)** berbasis **Go (Gin Framework)**.  
Mendukung autentikasi terpusat untuk berbagai consumer (aplikasi client), dengan alur mirip **OAuth2 Authorization Code Flow**.

---

## ✨ Features
- Registrasi & Login user melalui SSO
- Redirect ke consumer dengan **one-time token**
- Tukar one-time token → JWT (access_token + refresh_token)
- Verify access token untuk cek validitas & permission user
- Refresh token untuk memperpanjang sesi
- Manajemen aplikasi consumer (client key & callback URL)
- Ambil profile user berdasarkan access token

---

## 📂 API Routes

### 🔑 Auth
| Method | Endpoint              | Deskripsi |
|--------|-----------------------|-----------|
| POST   | `/api/auth/login`         | Login SSO user |
| POST   | `/api/auth/logout`        | Logout user (invalidate token) |
| POST   | `/api/auth/register`      | Register user baru |
| POST   | `/api/auth/refresh`       | Refresh JWT token |
| POST   | `/api/auth/verify_access` | Verifikasi client key & callback URL sebelum memberikan izin pakai SSO |
| POST   | `/api/auth/sso`           | Tukar one-time token menjadi JWT (sign in ke consumer app) |

---

### 🛠️ Application
| Method | Endpoint                            | Deskripsi |
|--------|-------------------------------------|-----------|
| GET    | `/api/application/`                     | List semua aplikasi client |
| POST   | `/api/application/create`               | Buat aplikasi client baru |
| GET    | `/api/application/:id`                  | Lihat detail aplikasi client |
| GET    | `/api/application/:id/refresh`          | Generate client key baru |
| DELETE | `/api/application/:id`                  | Hapus aplikasi client |
| POST   | `/api/application/:id/callback`         | Tambah callback URL |
| PATCH  | `/api/application/:id/:callback_id`     | Update callback URL |
| DELETE | `/api/application/:id/:callback_id`     | Hapus callback URL |

---

### 👤 User
| Method | Endpoint      | Deskripsi |
|--------|---------------|-----------|
| GET    | `/api/user/`      | Ambil informasi user berdasarkan access token |

---

## 🛠️Tect Stack
Go 1.24.0

Gin (github.com/gin-gonic/gin v1.10.1)

JWT (github.com/golang-jwt/jwt/v5 v5.3.0)

GORM + PostgreSQL
(gorm.io/gorm v1.30.5, gorm.io/driver/postgres v1.6.0)

Validator (github.com/go-playground/validator/v10 v10.27.0)

dotenv loader (github.com/joho/godotenv v1.5.1)

JSON parser (github.com/goccy/go-json v0.10.5)

Dependensi lain: sonic, pgx, protobuf, yaml.v3, dll (lihat go.mod)

---

## ⚙️ Setup & Instalation
1. Clone repo
   git clone https://github.com/username/SSO_BE_API.git
   cd SSO_BE_API
2. Buat file .env
   Salin .env.example ke .env lalu sesuaikan nilainya:
   cp .env.example .env
   Isi default .env.example:
## Database Config
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=sso

## App Config
ENV=development
PORT=8080

## JWT Config
JWT_KEY=your_jwt_secret
JWT_REFRESH_TOKEN_KEY=your_jwt_refresh_secret

3. Install dependencies
   go mod tidy
4. Run server
   go run main.go
   Server akan jalan di http://localhost:8080.

---

## 🔑Security Notes

access_token → masa aktif pendek (misalnya 15 menit).

refresh_token → masa aktif lebih panjang (misalnya 7 hari), simpan hanya di BE consumer, jangan di FE.

code (one-time token) hanya berlaku sekali & expired cepat (30–60 detik).

---