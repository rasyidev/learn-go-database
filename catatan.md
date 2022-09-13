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

## Eksekusi Perintah SQL
- Di Go Lang, terdapat function `(DB) ExecContext(context, query, params)` yang dapat digunakan untuk mengirim perintah SQL yang tidak mengembalikan data. Cth: insert dll
- Terdapat `context` sebagai parameter, sehingga dapat mengirim sinyal cancel jika ingin membatalkan perintah SQL
- `(DB) QueryContext(context, query, params)(rows, err)`, untuk mengirim perintah sql yang mengembalikan data. Cth: select
  - rows disarankan diClose menggunakan `rows.Close()`

## Tipe Data Column dan Representasi pada Go Lang
**Tipe data database -> tipe data Go Lang**
- VARCHAR, CHAR -> string
- INT, BIGINT -> int32, int64
- FLOAT, DOUBLE -> float32, float64
- BOOLEAN -> bool
- DATE, DATETIME, TIME, TIMESTAMP -> time.Time

**Driver Go Lang secara default tidak support time.Time**
Strategi:
DATE, DATETIME, TIMESTAMP -> []byte -> parsing ke string -> time.Time (tapi ribet)
setting  di driver golang parseDate=True pada url `dataSourceName`


## Nullable Type
- Driver Go lang tidak mengerti tipe database
- Khusus kolom yang nullable, `rows.Scan` akan menjadi error, perlu dihandle
- Error:

```bash
--- FAIL: TestPerintahSelectComplex (0.01s)
panic: sql: Scan error on column index 2, name "email": converting NULL to string is unsupported [recovered]
        panic: sql: Scan error on column index 2, name "email": converting NULL to string is unsupported
```
- Konversi otomatis untuk tipe database NULL tidak didukung oleh Driver MySQL Go Lang
- Khusus tipe kolom nullable, perlu menggunakan tipe data yang ada dalam package `sql`

**Tipe Data Nullable package `sql`**
- string -> `database.sql.NullString`