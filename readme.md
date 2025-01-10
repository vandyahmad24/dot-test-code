# DOT Test By Vandy Ahmad
## How to running API

### Step 1: Clone this repository
Please clone this repository with command `git clone https://github.com/vandyahmad24/dot-test-vandy.git`

### Step 2: Install dependencies
Please install dependencies with command `go mod tidy`

### Step 3: Running API
Please running API with command `make run-api`

### Step 4: Access API
Please access API with URL `http://localhost:8000/`

## How to running Migration

### Step 1: Running Migration
Please running migration with command `make migrate-up`



### Design Pattern
Aplikasi ini menggunakan arsitektur **hexagonal architecture**. Dalam aplikasi ini, kita memiliki lapisan-lapisan yang memisahkan logika bisnis dari antarmuka pengguna dan infrastruktur. Dengan menggunakan arsitektur hexagonal, kita dapat dengan mudah mengganti komponen eksternal, seperti database, redis, atau thrid party lain tanpa mempengaruhi logika bisnis utama, karena setiap komponen terhubung melalui antarmuka yang jelas.

### Kelebihan Arsitektur Hexagonal

1. **Pemeliharaan yang Lebih Mudah**: Dengan memisahkan logika bisnis dari infrastruktur, perubahan pada salah satu lapisan tidak akan mempengaruhi lapisan lainnya. Ini memudahkan tim pengembang untuk memelihara dan memperbarui aplikasi tanpa khawatir akan dampak negatif.

2. **Dapat Diuji dengan Lebih Baik**: Logika bisnis yang terisolasi memungkinkan pengujian unit yang lebih efektif. Kita dapat menguji logika bisnis secara terpisah dari interaksi dengan sistem luar, sehingga mengurangi kompleksitas dan meningkatkan keandalan pengujian.

3. **Peningkatan Kualitas Kode**: Dengan mengikuti prinsip pemisahan kekhawatiran, kode menjadi lebih bersih dan terstruktur dengan baik. Ini membantu dalam menjaga kualitas kode secara keseluruhan dan memudahkan pengembang baru untuk memahami dan berkontribusi pada proyek.

4. **Reusabilitas**: Komponen logika bisnis dapat digunakan kembali di berbagai konteks atau aplikasi, menghemat waktu dan usaha pengembangan saat memperluas atau membuat aplikasi baru.

Arsitektur hexagonal menyediakan kerangka kerja yang kuat untuk membangun aplikasi yang dapat beradaptasi dengan perubahan dan bertahan dalam jangka panjang.

