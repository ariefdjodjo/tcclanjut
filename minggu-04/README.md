### MINGGU KE-04

Dalam berbagaikasus kita dapat menggunakan dua atau lebih container pada satu Server Docket. Container-container tersebut nantinya akan dapat terhubung dengan menggunakan Docker Media komunikasi yaitu Network atau menggunakan Link. Pada kesematan ini kita akan sedikit mempraktekan penggunaan Link dan Network pada docker dengan playgroud Katakoda.


## A. KONEKSI ANTARA 2 NODE MENGGUNAKAN NETWORK
"Communicating between containers".

1. Jalankan container
    Kita Coba cek network yang tersedia dalam docker dengan perintah ``docker network ls``

    ![01](img/01.png)

    Pada gambar diatas dapat kita lihat bahwa terdapat 3 (tiga)
    