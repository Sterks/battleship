Создать поле


```
POST http://localhost:8000/create-matrix
Accept: application/json

{"range": 10}
```

Расставляем корабли

```
POST http://localhost:8000/ship
Accept: application/json

{"Coordinates": "1A 1B 1C, 9D 9E, 5F 6F 7F, 11F"}
```

Атакуем

```
POST http://localhost:8000/shot
Accept: application/json

{"Coord": "5E"}
```

Очищаем поле

