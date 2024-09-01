Запросы:
1) создание пользователя
http://localhost:8080/users
{
  "id": "1",
  "first_name": "Иван",
  "last_name": "Иванов",
  "age": 30,
  "recording_date": "2022-01-01T00:00:00Z"
}

2) все пользователи
http://localhost:8080/users/all

3) по дате и возрасту
http://localhost:8080/users/by_date_and_age
{
  "date_from": "2022-01-01T00:00:00Z",
  "date_to": "2022-01-31T00:00:00Z",
  "age_from": 20,
  "age_to": 40
}