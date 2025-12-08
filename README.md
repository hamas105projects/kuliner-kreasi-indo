# Kuliner Kreasi Indo 🍲

## 📌 Deskripsi
Kuliner Kreasi Indo adalah proyek API yang dikembangkan sebagai bagian dari kebutuhan interview.  
API ini dirancang untuk mengelola data kuliner secara terstruktur, sehingga memudahkan integrasi dengan aplikasi lain maupun sistem internal.

## 🚀 Fitur Utama
- **Modular API**: Endpoint terpisah untuk user dan sale order.  
- **CRUD Operations**: Mendukung create, read, update, delete data.  
- **Response Utility**: Format respons JSON yang konsisten.  
- **Scalable Structure**: Folder modular (handler, model, repository, service).  

## 🛠️ Teknologi
- [Go (Golang)](https://golang.org/) sebagai bahasa utama backend  
- REST API untuk komunikasi data  
- Git untuk version control  

# 📦 Structure Folder Overview
/pos-backend/
│
├── server/
│   └── server.go                 # entrypoint
│
├── config/
│   └── config.go               # load env, db connection
│
├── modules/                    # semua fitur disatukan dalam modules
│   ├── auth/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── jwt.go
│   │
│   ├── user/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   │
│   └── saleorder/
│       ├── handler.go
│       ├── service.go
│       ├── repository.go
│       └── model.go
│
├── middleware/
│   ├── auth.go                 # parse JWT
│   └── rbac.go                 # role based access
│
├── utils/
│   ├── response.go             # success/failed response
│   └── utils.go                # pagination, hash, etc (disatukan)
│
├── migrations/
│   └── kki.sql
│
├── go.mod
├── go.sum
├── main.go
└── .env

# 📦 Dependencies Overview

Project ini menggunakan beberapa package penting untuk membangun backend service dengan **Golang**. Berikut penjelasan mengapa setiap package dipakai:

---

## 🚀 Web Framework
- **[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)**
  - Framework HTTP yang ringan, cepat, dan powerful.
  - Mendukung middleware, routing yang fleksibel, serta JSON binding.
  - Cocok untuk membangun REST API dengan performa tinggi.

---

## 🗄️ ORM & Database Driver
- **[gorm.io/gorm](https://gorm.io/gorm)**
  - ORM (Object Relational Mapping) untuk Golang.
  - Memudahkan query database dengan model berbasis struct.
  - Mendukung fitur seperti auto migration, hooks, dan associations.

- **[gorm.io/driver/postgres](https://gorm.io/driver/postgres)**
  - Driver PostgreSQL untuk GORM.
  - Digunakan agar GORM bisa berkomunikasi dengan database PostgreSQL.

- **[github.com/jackc/pgx/v5](https://github.com/jackc/pgx)**
  - PostgreSQL driver murni untuk Golang.
  - Memberikan performa lebih tinggi dibanding `database/sql` default.
  - Bisa digunakan langsung atau sebagai backend untuk GORM.

---

## 🔑 Authentication & Security
- **[github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)**
  - Library untuk membuat dan memverifikasi JSON Web Token (JWT).
  - Digunakan untuk sistem autentikasi berbasis token.
  - Mendukung berbagai algoritma signing (HS256, RS256, dll).

- **[golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**
  - Implementasi algoritma hashing **bcrypt**.
  - Digunakan untuk menyimpan password dengan aman.
  - Memberikan proteksi terhadap brute-force attack.

---

## 🆔 Utilities
- **[github.com/google/uuid](https://github.com/google/uuid)**
  - Library untuk membuat UUID (Universally Unique Identifier).
  - Berguna untuk ID unik pada user, transaksi, atau resource lain.

- **[github.com/joho/godotenv](https://github.com/joho/godotenv)**
  - Membaca file `.env` dan memuat environment variables ke dalam aplikasi.
  - Memudahkan pengaturan konfigurasi (database, secret key, dll) tanpa hardcode.

---

## **Testing Checklist **

### 1️⃣ Authentication (JWT)
| Endpoint | Method | Test Case | Expected Result |
|----------|--------|-----------|----------------|
| `/api/auth/login` | POST | Login dengan email/password valid | 200, JWT token dikembalikan, status: success |
| `/api/auth/login` | POST | Login dengan email/password salah | 400, status: failed, message error |
| `/api/auth/logout` | POST | Logout user (opsional) | 200, status: success, JWT invalid (jika diimplementasikan) |

---

### 2️⃣ RBAC Middleware
| Role | Endpoint | Test Case | Expected Result |
|------|----------|-----------|----------------|
| Cashier | `/users/cashier/*` | Akses endpoint owner-only | 403 Forbidden, status: failed |
| Owner | Semua | Akses semua endpoint | 200 OK, status: success |

---

### 3️⃣ CREATE ADMIN
Turn off with comment this func temporary
//salesGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
//salesGroup.Use(middleware.RBACMiddleware([]string{"admin", "cashier"}))

| Endpoint | Method | Test Case | Expected Result |
|----------|--------|-----------|----------------|

| `/api/users/cashier` | POST | create user via cshier | 200, role = admin, status: success |



salesGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
salesGroup.Use(middleware.RBACMiddleware([]string{"admin", "cashier"}))

### 3️⃣ CRUD User Cashier (Owner Only)
| Endpoint | Method | Test Case | Expected Result |
|----------|--------|-----------|----------------|
| `/api/users/cashier` | POST | Create new cashier | 200, role = cashier, status: success |
| `/api/users/casier` | GET | List cashiers dengan pagination (`?page=1&limit=5`) | 200, status: success, data array sesuai limit |
| `/api/users/cashier/:id` | GET | Ambil detail cashier valid | 200, status: success, data cashier |
| `/api/users/cashier/:id` | GET | Ambil detail cashier invalid | 404, status: failed, message "not found" |
| `/api/users/cashier/:id` | PUT | Update cashier (owner) | 200, data updated, status: success |
| `/api/users/cashier/:id` | PUT | Update cashier (cashier) | 403 Forbidden, status: failed |
| `/api/users/cashier/:id` | DELETE | Hapus cashier (owner) | 200, status: success |
| `/api/users/cashier/:id` | DELETE | Hapus cashier (cashier) | 403 Forbidden, status: failed |

---

### 4️⃣ CRUD Sale Order (Cashier & Owner)
| Endpoint | Method | Test Case | Expected Result |
|----------|--------|-----------|----------------|
| `/api/sale-orders` | POST | Buat order dengan beberapa item | 200, total_amount dihitung benar, status: success |
| `/api/sale-orders` | GET | List orders dengan pagination (`?page=1&limit=5`) | 200, status: success, data sesuai limit |
| `/api/sale-orders/:id` | GET | Ambil order valid | 200, status: success, data termasuk items |
| `/api/sale-orders/:id` | GET | Ambil order invalid | 404, status: failed, message "not found" |
| `/api/sale-orders/:id` | PUT | Update order items | 200, total_amount berubah sesuai update, status: success |
| `/api/sale-orders/:id` | DELETE | Hapus order | 200, items ikut terhapus (CASCADE), status: success |

---

### 5️⃣ Pagination & Limit
- Test `?page=2&limit=5` → pastikan data sesuai halaman & jumlah limit  
- Test `?page=1&limit=100` → menampilkan semua jika total < 100  

---

### 6️⃣ Response Format & HTTP Code
- **Success:** 200, status: "success", message, data  
- **Failed / Validation:** 400, status: "failed", message error  
- **Forbidden / RBAC:** 403, status: "failed", message error  
- **Not Found:** 404, status: "failed", message "not found"  
- **Server error:** 500, status: "failed", message error  

---

### 7️⃣ Security & Edge Cases
- JWT expired → endpoint protected menolak akses  
- Role invalid → endpoint forbidden  
- Cashier tidak bisa update/delete cashier lain  
- Sale order tidak bisa dibuat tanpa item  
- Qty item harus > 0, price_snapshot ≥ 0 → validasi gagal → 400  

---

### 8️⃣ Optional: Product CRUD (Owner Only)
| Endpoint | Method | Test Case | Expected Result |
|----------|--------|-----------|----------------|
| `/api/products` | POST | Create product valid | 200, status: success |
| `/api/products/:id` | GET | Ambil product valid | 200, status: success |
| `/api/products/:id` | PUT | Update product | 200, status: success |
| `/api/products/:id` | DELETE | Delete product | 200, status: success |
| `/api/products` | GET | List products dengan pagination | 200, status: success |

---

> **Catatan:**  
> Semua endpoint **harus menggunakan JWT** pada header `Authorization: Bearer <token>` untuk akses.  
> Semua response **mengikuti standar** success / failed seperti dijelaskan.


## 🙏 Penutup
Project **Kuliner Kreasi Indo** ini dikembangkan khusus sebagai bagian dari proses **interview** di Kuliner Kreasi Indo.  
Tujuan utama repo ini adalah untuk menunjukkan kemampuan dalam:
- Membangun API dengan **Golang** secara modular
- Mengelola data dengan struktur yang rapi dan maintainable
- Menyediakan dokumentasi yang jelas untuk memudahkan review

Semoga repo ini dapat menjadi gambaran nyata atas skill dan pendekatan saya dalam membangun solusi backend yang efisien.