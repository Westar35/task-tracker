Команды Powershell для демонстрации и тестов

---

Залогиниться

$login = Invoke-RestMethod `
  -Method Post `
  -Uri "http://localhost:8080/auth/login" `
  -ContentType "application/json" `
  -Body '{"email":"danis@exnet.ru","password":"Fgrths197+"}'

  ---

  $token = $login.token
  $token

  ---

  Получить все задачи

  Invoke-RestMethod `
  -Method Get `
  -Uri "http://localhost:8080/tasks" `
  -Headers @{ Authorization = "Bearer $token" }

  ---

  Создать задачу

  Invoke-RestMethod `
  -Method Post `
  -Uri "http://localhost:8080/tasks" `
  -ContentType "application/json" `
  -Headers @{ Authorization = "Bearer $token" } `
  -Body '{"title":"Новая задача"}'

  ---

  Получить задачу по id

  Invoke-RestMethod `
  -Method Get `
  -Uri "http://localhost:8080/tasks/1" `
  -Headers @{ Authorization = "Bearer $token" }

  ---

  обновить задачу по id

  Invoke-RestMethod `
  -Method Put `
  -Uri "http://localhost:8080/tasks/1" `
  -ContentType "application/json" `
  -Headers @{ Authorization = "Bearer $token" } `
  -Body '{"title":"Обновленная задача","status":true}'

  ---

  Удалить задачу по id

  Invoke-RestMethod `
  -Method Delete `
  -Uri "http://localhost:8080/tasks/1" `
  -Headers @{ Authorization = "Bearer $token" }