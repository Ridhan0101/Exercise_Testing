Nama: Muhammad Ridhan Khoirullah
ID Camp: BE9259656

Kode di atas adalah implementasi sederhana dari sebuah kalkulator ilmiah dalam bahasa Go. Kalkulator ini dapat melakukan berbagai operasi matematika dasar dan beberapa operasi trigonometri serta operasi lainnya.

Dalam kode tersebut, terdapat beberapa fungsi yang mengimplementasikan operasi-operasi matematika seperti penjumlahan, pengurangan, perkalian, dan pembagian. Fungsi-fungsi tersebut dapat diakses dari luar package main karena mereka memiliki huruf kapital pada awal nama fungsi, sehingga secara otomatis dianggap diekspor oleh Go.

Selain itu, terdapat pula fungsi-fungsi untuk menghitung akar kuadrat, pangkat dua, dan pangkat tiga dari suatu bilangan. Untuk operasi trigonometri, terdapat fungsi-fungsi untuk menghitung sinus, kosinus, dan tangen baik dalam radian maupun derajat.

Di dalam fungsi main, pengguna diminta untuk memilih operasi yang ingin dilakukan melalui menu. Setelah memilih operasi, pengguna diminta untuk memasukkan angka-angka yang diperlukan untuk operasi tersebut. Selanjutnya, program akan memanggil fungsi yang sesuai berdasarkan pilihan pengguna dan menampilkan hasil operasi tersebut.

Untuk pengujian menggunakan Ginkgo, kita dapat melihat pada  file calculatorscience_test.go. File tersebut akan berisi kode pengujian menggunakan Ginkgo. Dalam file tersebut, kita dapat menuliskan beberapa kasus uji untuk menguji setiap fungsi yang ada dalam kode utama. Ginkgo akan membantu kita dalam menulis dan menjalankan kasus uji tersebut serta menampilkan hasilnya dengan jelas.


File calculatorscience_test.go tersebut merupakan file pengujian untuk memastikan bahwa fungsi-fungsi dalam calculatorscience.go berfungsi dengan benar. Di dalamnya, terdapat beberapa contoh pengujian yang menggunakan Ginkgo.

Pada blok Context("Testing basic arithmetic operations", func() {}), dilakukan pengujian terhadap operasi aritmatika dasar seperti penjumlahan, pengurangan, perkalian, dan pembagian. Setiap operasi diuji dengan menggunakan fungsi Expect dari Gomega untuk membandingkan hasil yang diharapkan dengan hasil yang sebenarnya.

Ginkgo adalah sebuah framework pengujian yang ditulis dalam bahasa Go (Golang) dan dibuat khusus untuk pengujian behavior-driven development (BDD). Dengan menggunakan Ginkgo, Anda dapat menulis spesifikasi pengujian dengan gaya yang lebih deskriptif dan ekspresif, sehingga memudahkan dalam memahami perilaku yang diharapkan dari sebuah sistem atau komponen.