# websocket
Saya menggunakan service ini untuk melakukan update pada sebuah halaman berdasarkan interaksi dihalaman lain, ini lebih efisien dari pada menggunakan long polling ajax.

## INSTALASI
1. clone repository ini `git clone git@github.com:barokurniawan/websocket.git`
2. install [go-dep](https://github.com/golang/dep)
3. jalankan `dep ensure` untuk menginstall dependency

## RUN
```
<script type="text/javascript">
    const SOCKET_CHANNEL = "PITSTOP_FIRST";
    window.customSocket = {};

    window.customSocket.init = function () {
        if (!(window.WebSocket)) {
            console.log("Browser doesnot support websocket");
            return
        }

        window.customSocket = new WebSocket("ws://localhost:3001/socket?channel=" + SOCKET_CHANNEL)

        window.customSocket.onopen = function () {
            console.log("Connect to websocket, send and receive on chanel " + SOCKET_CHANNEL);
        }

        window.customSocket.onmessage = function (event) {
            console.log("incoming...");
            console.log(event);
        }

        window.customSocket.onclose = function () {
            console.log("Disconnected from websocket");
        }
    }

    window.customSocket.doSendMessage = function () {
        window.customSocket.send(JSON.stringify({
            Channel: SOCKET_CHANNEL,
            Message: "CHANGE"
        }));
    }

    window.onload = window.customSocket.init
</script>
```

```
window.customSocket.send(JSON.stringify({
    Channel: "PITSTOP_FIRST", // or MINE
    Message: "CHANGE"
}));
```
### Thanks
[Noval Agung](https://github.com/novalagung/dasarpemrogramangolang) - tutorial web socket
