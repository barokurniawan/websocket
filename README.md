# websocket
Saya menggunakan service ini untuk melakukan update pada sebuah halaman berdasarkan interaksi dihalaman lain, ini lebih efisien dari pada menggunakan long polling ajax.

# INSTALASI
1. clone repository ini `git clone git@github.com:barokurniawan/websocket.git`
2. install [go-dep](https://github.com/golang/dep)
3. jalankan `dep ensure` untuk menginstall dependency

# Menjalankan Contoh
- Rename file appconfig.toml.example menjadi appconfig.toml `cp appconfig.toml.example appconfig.toml`
  Konfigurasi Address, Port, AllowedOrigins ada di file ini.

- Ketikan `go run main.go`
- Buka http://localhost:3001/ atau dengan yang sesuai di file appconfig.toml 
  Buka address tersebut di 2 tab, regular dan incognito 
- ketik di console dan cek di masing-masing console : 
```
window.customSocket.send(JSON.stringify({
    Channel: "MINE", //channel bisa diganti dengan PITSTOP_FIRST atau MINE
    Message: "CHANGE"
}));
```

# Thanks
[Noval Agung](https://github.com/novalagung/dasarpemrogramangolang) - tutorial web socket
