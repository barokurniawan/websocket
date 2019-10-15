# websocket
library ini adalah hasil pengembangan dari chat app webscoket noval agung, untuk dokumentasi asli  [klik disini](https://dasarpemrogramangolang.novalagung.com/C-28-golang-web-socket.html). Saya menggunakan websocket ini untuk melakukan update pada 
sebuah halaman berdasarkan interaksi dihalaman lain, ini lebih efisien dari pada menggunakan sistem polling ajax.

# INSTALASI
1. clone repository ini
2. install [go-dep](https://github.com/golang/dep)
3. jalankan `dep ensure` untuk menginstall dependency

# Menjalankan Contoh
masuk ke folder project lalu ketikan "go run main.go" .disini host yang digunakan localhost dengan port 8080.
Untuk pengaksesan simulasi :
 - akses channel PITSTOP_FIRST
    localhost:8080/
 - akses channel MINE
    localhost:8080/other

Prinsip kerja nya sama seperti chat group, setiap pesan yang di kirim hanya bisa dibaca oleh connection yang ada 
pada channel yang sama. 
Buka masing masing halaman simulasi diatas pada browser biasa dan incognito, 
kemudia ketik di console : 
```
window.customSocket.send(JSON.stringify({
    Channel: "MINE", //channel bisa diganti dengan PITSTOP_FIRST atau MINE
    Message: "CHANGE"
}));
```


# Thanks
[Noval Agung](https://github.com/novalagung/dasarpemrogramangolang) - tutorial web socket
