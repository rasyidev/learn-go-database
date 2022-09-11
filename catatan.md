## Database Driver
- [drivers](golang.org/s/sqldrivers)
- `go get -u github.com/go-sql-driver/mysql`

## Membuat Koneksi ke Database
- `sql.Open(driver, dataSourceName)`
  - driver: "mysql", etc
  - dataSourceName, url database
- Jika object `sql.DB` tidak digunakan lagi, sebaiknya ditutup menggunakan function `Close()`

## Database Pooling
- Manajemen koneksi database
- Dalam Go Lang, manajemen koneksi database sudah diatur, tidak perlu manual lagi.
- Dapat menentukan jumlah minimal dan maksimal koneksi

**Pengaturan Database Pooling**
- `SetMaxIdleConns(number)`, Mengatur jumlah koneksi minimal
- `SetMaxOpenConns(number)`, Mengatur jumlah koneksi maksimal
- `SetConnMaxIdle(duration)`, Mengatur seberapa lama koneksi idle (tidak digunakan)
- `SetConnMaxLifetime(duration)`, Mengatur seeberapa lama koneksi dapat digunakan