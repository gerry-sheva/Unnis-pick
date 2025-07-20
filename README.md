# Project unnis_pick

## Essay
### Jelaskan pemahaman Anda mengenai Depedency Injection dan Pointer ?
Dependency Injection adalah sebuah teknik dalam pemrograman yang memungkinkan objek untuk mendapatkan dependensi yang diperlukan melalui konstruktor, bukan melalui pembuatan objek secara langsung. Dengan kata lain, objek tersebut tidak bertanggung jawab atas dependensi-dependensi yang diperlukannya, menyerahkan kontrol sepenuhnya kepada framework atau yang mengkonstruksi objek tersebut. Dependency Injection memungkinkan kita untuk mengisolasi dependensi dan membuat kode menjadi modular dan lebih mudah untuk diuji dan diubah.

Dalam proyek ini, teknik Dependency Injection digunakan pada package `service` dan `server`. Package `service`, lebih tepatnya `service/product.go`, membutuhkan dependensi dari package `cache` dan `searchengine`. Karena kedua package tersebut menggunakan service dari pihak ketiga, ada kemungkinan kalau kedepannya kita ingin mengganti service tersebut dengan service yang lain. Contohnya, saat ini `cache` diimplementasikan dengan menggunakan Redis, sedangkan `searchengine` diimplementasikan dengan menggunakan Elasticsearch. Bisa saja kalau besok kita memutuskan untuk menggunakan Memcached sebagai cache dan Algolia sebagai search engine. Dalam situasi ini, kita bisa menggunakan teknik Dependency Injection untuk memudahkan penggantian service tersebut. Hal ini dapat dilakukan dengan menggunakan interface sebagai tipe data yang digunakan untuk mendefinisikan kontrak antara package `service` dan `cache` serta `searchengine`, bukan implementasi konkret dari kedua service tersebut. Dengan demikian, kita tidak perlu mengubah kode di package `service` ketika kita ingin mengganti service cache atau search engine. Sama halnya dengan package `server` yang memiliki dependensi terhadap package `service`.

Dalam konteks programming, sebuah program pada dasarnya memiliki 2 jenis memori: Stack dan Heap. Stack adalah jenis memori yang umumnya dimiliki oleh setiap function. Setiap function memiliki akses terhadap memori stacknya masing-masing, dan akan dihapus ketika function tersebut selesai dieksekusi. Karena ukuran dari memori stack terbatas, maka compiler harus mengetahui terlebih dahulu berapa banyak memori yang diperlukan oleh setiap function. Oleh karena itu, jenis-jenis data yang bisa disimpan di memori stack adalah data primitif seperti integer, float, string, dan data struct yang ukurannya tidak terlalu besar dan tidak akan berubah selama runtime.

Untuk jenis-jenis data non-primitif (compound dan complex) seperti struct, array, dan map, data tersebut akan disimpan di memori heap. Dibandingkan dengan stack, heap memiliki ukuran yang lebih besar dan data yang disimpan di heap memiliki masa hidup yang lebih panjang.

Pointer adalah sebuah tipe data yang menyimpan alamat memori di heap. Dengan menggunakan pointer, kita dapat mengakses dan memodifikasi data tersebut. Karena pointer memiliki ukuran yang tetap, maka kita bisa menyimpan pointer di dalam stack, sehingga memungkinkan kita untuk mengelola data yang lebih kompleks dan dinamis. Pointer juga bisa dikirimkan ke stack lain melalui pass by reference. Dengan demikian, function laing juga bisa mengakses data tersebut tanpa melakukan copy terlebih dahulu, yang mana lebih mahal untuk dilakukan.

### Jelaskan pemahaman Anda mengenai Role Base Access Control?
Role-Based Access Control (RBAC) adalah sebuah metode akses kontrol yang berbasis peran (role) yang diberikan kepada pengguna atau entitas lain untuk mengakses resource yang tersedia dalam sistem. Secara garis besar, setiap role yang ada memiliki permission yang berbeda-beda terhadap resource yang tersedia dalam sistem. Permission ini bisa berupa akses read, write, atau execute terhadap resource yang tersedia dalam sistem. Permission ini bisa sangat granular dan bisa ditentukan dengan sangat detail. Oleh karena itu, akan sangat merepotkan jika setiap user harus diberikan permission secara manual. Terlebih lagi, metode ini bisa sangat rentan akan kesalahan. Dengan menggunakan RBAC, kita bisa melekatkan permission-permission yang ada pada peran yang diberikan kepada user. Sehingga, pengelolaan akses kontrol menjadi lebih mudah dan efisien.

### Jelaskan step-step cara menangani issue memory leak di go yang Anda ketahui?
Pada dasarnya, Go memiliki garbage collector (GC) yang akan menghapus data yang tidak lagi memilki pointer atau tidak bisa diakses lagi. Akan tetapi, memory leak masih bisa muncul jika data tersebut sudah tidak lagi digunakan tetapi masih memiliki pointer yang aktif. Dalam konteks web development, memory leak paling rentan muncul ketika berurusan dengan goroutine dan koneksi. Goroutine dan koneksi yang tidak ditutup dengan benar akan menyebabkan memory leak. Sehingga, penanganan memory leak yang paling mudah adalah dengan memastikan bahwa goroutine dan koneksi yang digunakan sudah ditutup dengan benar. Hal ini bisa dilakukan dengan mengikuti best practice seperti menggunakan defer untuk menutup koneksi dan context dan done channel untuk goroutine.

Jika memory leak terjadi, kita bisa menggunakan `pprof` untuk mendiagnosa memori program dengan langkah-langkah berikut:
1. Import `"net/http/pprof"` dan jalankan aplikasi dan pprof di goroutine
2. Sebelum melakukan testing, dapatkan data memori sebagai komparasi dengan
```bash
curl -o baseline.prof http://localhost:6060/debug/pprof/heap
```
3. Lakukan testing di bagian yang memungkinkan terjadinya memory leak
4. Setelah testing selesai, dapatkan data memori lagi dengan
```bash
curl -o aftertest.prof http://localhost:6060/debug/pprof/heap
```
5. Bandingkan data memori sebelum dan sesudah testing dengan
```bash
go tool pprof -http=:8080 baseline.prof aftertest.prof
```
6. Analisis hasil pprof untuk menemukan memory leak

### Apakah Anda berpengalaman dengan AWS ?
Ya, saya pernah menggunakan AWS EC2 dalam environment production.

## Project
Proyek ini menggunakan PostgreSQL docker container sebagai RDBMS dan https://github.com/pressly/goose sebagai alat migrasi database.
Untuk menjalankan aplikasi, jalankan perintah berikut:
```bash
make unnis
```
