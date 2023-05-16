1. Check all albums
http://localhost:8080/albums  GET

2. Add one album
http://localhost:8080/albums  POST
```
{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}
```

3. Query an album
http://localhost:8080/albums/2
```
{
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
}
```
