# quiz_sanber_batch68

## cara penggunaan
- terdapat 2 fitur dalam project ini yaitu CRUD untuk books/buku dan categories/kategori

untuk mengakses CRUD tersebut, perlu dilakukan authentifikasi terlebih dahulu melalui tab authorization pada aplikasi postman.
untuk username dan password pada project ini adalah basic auth dan bentuknya statis

silahkan dimasukkan 
username : admin
password : admin

setelah itu, silahkan pindah ke tab body pada postman dan pilih menu raw. 
cara penggunaan API nya ini yaitu sebagai berikut
"https://quizsanberbatch68-production.up.railway.app/api/"

## books
API books ini terdiri dari Create, Get, dan Delete. cara mengaksesnya yaitu sebagai berikut :
### GET
- Get semua book : https://quizsanberbatch68-production.up.railway.app/api/books (Pilih method GET)
- Get book berdasarkan id nya : https://quizsanberbatch68-production.up.railway.app/api/books/:id (Pilih method GET)

kegunaan API ini adalah untuk melihat data buku buku yang terdapat pada database.


### POST
- Create book : https://quizsanberbatch68-production.up.railway.app/api/books (Pilih method POST)
pilih menu "body" di postman dan silahkan input data tersebut
{
  "title": ,
  "description": ,
  "image_url": ,
  "release_year": ,
  "price": ,
  "total_page": ,
  "category_id": 
}

kegunaan API ini adalah untuk membuat data buku buku baru dan menyimpannya ke dalam database.


### DELETE
- DELETE book by id : https://quizsanberbatch68-production.up.railway.app/api/books/:id (Pilih method DELETE)

kegunaan API ini adalah untuk menghapus data buku yang ada pada database.


## categories
API category ini terdiri dari Create, Get, dan Delete juga, mirip seperti API book, namun perbedaannya adalah cara mengaksesnya. cara mengaksesnya yaitu sebagai berikut :
### GET
- Get semua category : https://quizsanberbatch68-production.up.railway.app/api/categories (Pilih method GET)
- Get book category id nya : https://quizsanberbatch68-production.up.railway.app/api/categories/:id (Pilih method GET)

kegunaan API ini adalah untuk melihat data kategori yang ada pada database.

### POST
- Create category : https://quizsanberbatch68-production.up.railway.app/api/categories (Pilih method POST)
pilih menu "body" di postman dan silahkan input data tersebut
{
  "name": 
}

kegunaan API ini adalah untuk membuat data kategori baru dan menyimpannya dalam database.


### DELETE
- DELETE category by id : https://quizsanberbatch68-production.up.railway.app/api/categories/:id (Pilih method DELETE)

kegunaan API ini adalah untuk menghapus data kategori dari database berdasarkan id yang dipilih.
