# interop-service-a

create app password
https://support.google.com/mail/answer/185833?hl=en

tambahkan jam pada notifikasi pembelian!

=========================================================================================================

Pertama Aktifkan Verifikasi dua langkah pada Email Pengirim, lalu generate APP password dengan prefrensi custom
lalu buat file .env yag berisi credential email dan password diisi dengan app password tadi


![alt text](![alt text](http://url/to/img.png))


lalu load data credential tadi dengan fungsi init

Lalu saat api di hit dengan method post http://localhost:5000/send-email dan diisi body dengan fromat json


![alt text](https://ibb.co/dMNR6Yb)


jika berhasil maka status akan berkode 200

lalu pada email penerima akan terkirim notifikasi pembelian dan akan di ikuti waktu pembelian


![alt text](https://ibb.co/QMfzPGd)
