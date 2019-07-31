## Membuat Database di CockroachDB Melalui Commant Line

1. Setelah kita jalankan CockroachDB, selanjutnya kita buka terminal baru dan menjalankan perintah ``cockroach sql --insecure --host=localhost:26257``
   
   ![db_01](img/db_01.png)

   perintah tersebut berarti kita akan masuk pada CLI database.

2. Untuk membuat database baru kita gunakan perintah ``CREATE DATABASE <nama db>`` dimana pada projet ini saya membuat database kampus.
   
   ![db_02](img/db_02.png)

3. selanjutnya kita buat 2 buah table yaitu jurusan dan mahasiswa. dengan perintah ``CREATE TABLE <nama table>``.
   
   membuat table jurusan
   ![db_03](img/db_03.png)

   membuat table mahasiswa
   ![db_04](img/db_04.png)

4. jika kita buka dengan browser maka hasilnya sebagai berikut:
   
   ![db_05](img/db_05.png)
   
   ![db_06](img/db_06.png)
   
   ![db_07](img/db_07.png)