Создать задачу
Invoke-RestMethod -Method POST `
  -Uri "http://localhost:8080/tasks" `
  -ContentType "application/json" `
  -Body '{"title":"Первая задача"}'
Получить все задачи
Invoke-RestMethod -Method GET `
  -Uri "http://localhost:8080/tasks"
Получить задачу по id
Invoke-RestMethod -Method GET `
  -Uri "http://localhost:8080/tasks/1"
Обновить задачу
Invoke-RestMethod -Method PUT `
  -Uri "http://localhost:8080/tasks/1" `
  -ContentType "application/json" `
  -Body '{"title":"Обновленная задача","status":true}'
Удалить задачу
Invoke-RestMethod -Method DELETE `
  -Uri "http://localhost:8080/tasks/1"
Проверить, что она удалилась
Invoke-RestMethod -Method GET `
  -Uri "http://localhost:8080/tasks/1"