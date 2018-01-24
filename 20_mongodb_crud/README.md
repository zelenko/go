# Notes
#### Make a request via Curl
```
curl -i 192.168.0.2:8080/items
```


#### Insert items using Curl
```
curl -i -X POST -d "isbn=001-8484314701&title=How to run Windows&author=Bill Gates&price=1.90" 192.168.0.2:8080/items/create/process
```

#### Insert items in MongoDB shell
```
db.books.insert([{"isbn":"978-1505255607","title":"The Time Machine","author":"H. G. Wells","price":5.99},{"isbn":"978-1503261960","title":"Wind Sand \u0026 Stars","author":"Antoine","price":14.99},{"isbn":"978-1503261961","title":"West With The Night","author":"Beryl Markham","price":14.99}])
```