## INSTALL GO (GOLANG) DI WINDOWS

Pada kesempatan ini saya menggunakan golang versi 1.12.5 Windows 32bit. Untuk link Download dapat di download di [Go-Download](https://golang.org/dl/). Selanjutnya kita jalankan file hasil download kita sehingga tahapannya sebagai berikut.

1. Jalankan File yang kita download tadi. Kemudian klik ``Next > ``
![01](img/01.png)

2. Selanjutnya kita akan disuguhkan License Agreement untuk go. Kita beri check pada Agreement selanjutnya kita klik ``Next``.

![02](img/02.png)

3. Selanjutnya kita pilih Destination tempat menginstall go. Disini kita letakkan di folder ``C:\go\``. Kemudian tekan ``Next``.

![03](img/03.png)

4. Golang siap di install. kita klik ``Install``.

![04](img/04.png)

5. Tunggu hingga proses install selesai. 

![05](img/05.png)

Setelah selesai kita klik ``Finish``

![06](img/06.png)

6. Kita cek apakah go yang kita install berhasil atau tidak. Kita buka Command Line (CMD), selanjutnya kita ketikka perintah ``go version`` maka akan menampilkan versi go yang kita install

![07](img/07.png)

Untuk melihat setting ``env`` golang, kita ketikkan perintah ``go env``. disini kita dapat melihat setting ``GOPATH`` dimana terletak pada direktori ``C:\Users\arief\go``. ``GOPATH`` digunakan untuk meletakkan project-project yang akan kita buat.

![08](img/08.png)