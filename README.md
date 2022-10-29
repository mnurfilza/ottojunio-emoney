# ottojunio-emoney

Panduan cara running service e-money.

1. Menggunakan docker cukup tulis `docker compose up` secara otomatis akan terinstall dan dapat runnig
2. terdapat 2 file yaml untuk perconfigkanya, mostly yang dirubah hanya bagian `db.yaml` saja untuk menyesuiakan database
3. Applikasi berjalan pada port ``8801``
4. berikut dilampirkan link JSON postman ``https://www.getpostman.com/collections/c76dd8fdbbc311cb958a``
5. ketika login pastikan mengambil ``x-token`` pada response header dan letakan diheader dengan key Authorization ketika ingin melakukan request, kecuali untuk API Register dan Login