<div align="center">

# 📋 TODO Scheduler

### Планировщик задач с поддержкой периодических повторений

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=flat&logo=sqlite)](https://www.sqlite.org)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com)
[![JWT](https://img.shields.io/badge/Auth-JWT-000000?style=flat&logo=jsonwebtokens)](https://jwt.io)


*Итоговый проект курса "Go-разработчик" от Яндекс Практикума*

[Возможности](#-возможности) •
[Быстрый старт](#-быстрый-старт) •
[API](#-api-endpoints) •
[Docker](#-docker) •
[Тесты](#-тесты)

</div>


---

## 📖 Описание


**TODO Scheduler** — веб-приложение для управления задачами с системой периодических повторений. Создавайте одноразовые задачи или настраивайте сложные расписания повторений — ежедневные, еженедельные, ежемесячные или ежегодные.


### 🎯 Что умеет приложение?

- ✨ **Создание задач** с датой, заголовком, комментарием и правилами повторения
- 📅 **Умные повторения** — поддержка сложных правил (каждый понедельник, последний день месяца, и т.д.)
- 🔍 **Поиск** по тексту или конкретной дате
- ✏️ **Редактирование** существующих задач
- ✅ **Автоматический расчёт** следующей даты для периодических задач
- 🗑️ **Удаление** ненужных задач
- 🔐 **Защита паролем** через JWT-токены

---

## ⭐ Выполненные задания повышенной сложности

<table>
<tr>
<td width="50%">

### 🔄 Расширенные правила повторения

- ✅ Еженедельные (`w 1,5` — пн, пт)
- ✅ Ежемесячные (`m 1,15` — 1-го и 15-го)
- ✅ Последний день месяца (`m -1`)
- ✅ Фильтрация по месяцам

</td>
<td width="50%">

### 🛡️ Дополнительные возможности

- ✅ Поиск по тексту и дате
- ✅ JWT аутентификация
- ✅ Docker контейнеризация
- ✅ Multi-stage сборка

</td>
</tr>
</table>

---

## 🚀 Быстрый старт

### Требования

- Go 1.21+
- Docker (опционально)

### Установка и запуск

```


# 1. Клонируйте репозиторий

git clone <your-repo-url>
cd go1fl-final

# 2. Установите зависимости

go mod download

# 3. Запустите сервер


go run main.go

# 4. Откройте в браузере

# http://localhost:7540

```

### ⚙️ Переменные окружения

```



# Настройка порта

export TODO_PORT=8080

# Путь к базе данных

export TODO_DBFILE=./data/scheduler.db

# Пароль для защиты (опционально)

export TODO_PASSWORD=mySecretPassword

```

| Переменная | Описание | По умолчанию |
|-----------|----------|--------------|
| `TODO_PORT` | Порт веб-сервера | `7540` |
| `TODO_DBFILE` | Путь к БД | `scheduler.db` |
| `TODO_PASSWORD` | Пароль (если не указан — доступ открыт) | — |

---

## 🐳 Docker

### Быстрый запуск

```


# Сборка образа

docker build -t todo-scheduler .

# Запуск без пароля

docker run -d \
--name todo-app \
-p 7540:7540 \
-v \$(pwd)/data:/data \
todo-scheduler

# Запуск с паролем

docker run -d \
--name todo-app \
-p 7540:7540 \
-v \$(pwd)/data:/data \
-e TODO_PASSWORD=mypassword \
todo-scheduler

```

### Windows (PowerShell)

```



docker run -d `  --name todo-app`
-p 7540:7540 `  -v ${PWD}/data:/data`
-e TODO_PASSWORD=mypassword `
todo-scheduler

```

### Управление

```

docker logs -f todo-app       \# Просмотр логов
docker stop todo-app          \# Остановка
docker restart todo-app       \# Перезапуск
docker rm todo-app            \# Удаление

```

---

## 🧪 Тесты

### Настройка

Отредактируйте `tests/settings.go`:

```

const (

WebServerURL = "http://localhost:7540"
FullNextDate = true  // Полный набор тестов



)

```

### Запуск

```


# Все тесты

go test ./tests/...

# Конкретные тесты

go test -run ^TestNextDate\$ ./tests
go test -run ^TestAddTask\$ ./tests
go test -run ^TestGetTasks\$ ./tests

```


---


## 📡 API Endpoints


### 🔓 Публичные

| Метод | Endpoint | Описание |
|-------|----------|----------|
| `POST` | `/api/signin` | Аутентификация |
| `GET` | `/api/nextdate` | Расчёт следующей даты |

### 🔐 Защищённые

| Метод | Endpoint | Описание |
|-------|----------|----------|
| `GET` | `/api/task?id=<id>` | Получить задачу |
| `POST` | `/api/task` | Создать задачу |
| `PUT` | `/api/task` | Обновить задачу |
| `DELETE` | `/api/task?id=<id>` | Удалить задачу |
| `GET` | `/api/tasks?search=<query>` | Список задач |
| `POST` | `/api/task/done?id=<id>` | Отметить выполненной |

### 📝 Примеры запросов

<details>
<summary><b>Создание задачи</b></summary>

```

curl -X POST http://localhost:7540/api/task \
-H "Content-Type: application/json" \
-d '{
"date": "20240201",
"title": "Поход к врачу",
"comment": "Принести результаты анализов",
"repeat": "m 1"
}'

```
</details>

<details>
<summary><b>Поиск задач</b></summary>

```








# По тексту

curl "http://localhost:7540/api/tasks?search=врач"

# По дате

curl "http://localhost:7540/api/tasks?search=01.02.2024"

```
</details>

<details>
<summary><b>Аутентификация</b></summary>

```

curl -X POST http://localhost:7540/api/signin \
-H "Content-Type: application/json" \
-d '{"password": "mypassword"}'

```
</details>

---


## 📋 Правила повторения


### Синтаксис


| Правило | Формат | Пример | Описание |
|---------|--------|--------|----------|
| **Ежедневно** | `d <число>` | `d 7` | Каждые 7 дней |
| **Ежегодно** | `y` | `y` | Каждый год в ту же дату |
| **Еженедельно** | `w <дни>` | `w 1,5` | Понедельник и пятница |
| **Ежемесячно** | `m <дни> [месяцы]` | `m 1,15` | 1-го и 15-го числа |

### 🎨 Примеры

```



d 1              \# Каждый день
d 30             \# Каждые 30 дней
y                \# Каждый год
w 1,2,3,4,5      \# Рабочие дни (пн-пт)
w 6,7            \# Выходные (сб-вс)
m 1              \# Первое число каждого месяца
m -1             \# Последний день каждого месяца
m -2             \# Предпоследний день месяца
m 15,30          \# 15-го и 30-го каждого месяца
m 10 3,6,9,12    \# 10-го числа ежеквартально
m 1,-1 2,8       \# Первый и последний день февраля и августа

```

---

## 🏗️ Структура проекта

```

go1fl-final/
├── 📄 main.go              \# Точка входа
├── 📦 go.mod               \# Зависимости
├── 🐳 Dockerfile           \# Docker конфигурация
├── 📖 README.md            \# Документация
│
├── 📁 pkg/
│   ├── 🌐 api/             \# HTTP обработчики
│   │   ├── api.go          \# Регистрация маршрутов
│   │   ├── auth.go         \# JWT аутентификация
│   │   ├── signin.go       \# Вход в систему
│   │   ├── handler.go      \# Базовые обработчики
│   │   ├── addtask.go      \# Создание задачи
│   │   ├── gettask.go      \# Получение задачи
│   │   ├── updatetask.go   \# Обновление задачи
│   │   ├── deletetask.go   \# Удаление задачи
│   │   ├── donetask.go     \# Выполнение задачи
│   │   └── tasks.go        \# Список задач
│   │
│   ├── 💾 db/              \# Работа с БД
│   │   ├── db.go           \# Инициализация SQLite
│   │   ├── task.go         \# CRUD операции
│   │   └── nextdate.go     \# Расчёт повторений
│   │
│   └── 🖥️ server/          \# Веб-сервер
│       └── server.go       \# HTTP сервер
│
├── 🎨 web/                 \# Статика (HTML/CSS/JS)
└── 🧪 tests/               \# Тесты
└── settings.go

```

---

## 🛠️ Технологии

<table>
<tr>
<td align="center" width="96">
<img src="https://go.dev/images/go-logo-blue.svg" width="48" height="48" alt="Go" />
<br>Go
</td>
<td align="center" width="96">
<img src="https://www.sqlite.org/images/sqlite370_banner.gif" width="48" alt="SQLite" />
<br>SQLite
</td>
<td align="center" width="96">
<img src="https://jwt.io/img/pic_logo.svg" width="48" height="48" alt="JWT" />
<br>JWT
</td>
<td align="center" width="96">
<img src="https://www.docker.com/wp-content/uploads/2022/03/vertical-logo-monochromatic.png" width="48" height="48" alt="Docker" />
<br>Docker
</td>
</tr>
</table>

- **Backend:** Go 1.21+
- **Database:** SQLite (modernc.org/sqlite)
- **Auth:** JWT (github.com/golang-jwt/jwt/v5)
- **Frontend:** Vanilla JavaScript
- **Container:** Docker (multi-stage build)

---

## 💡 Особенности реализации

### 🔐 Безопасность
- JWT токены с истечением через 8 часов
- SHA-256 хэширование паролей
- Защита от SQL-инъекций через параметризованные запросы
- HTTP-only cookies

### 🎯 Производительность
- Multi-stage Docker сборка
- Индексирование БД по дате
- Эффективные SQL запросы

### 🧹 Качество кода
- Обработка всех ошибок
- Валидация входных данных
- Комплексное тестирование

---

## 🤝 Проект
Проект выполнен в рамках итогового задания курса **"Go-разработчик"** от [Яндекс Практикума](https://practicum.yandex.ru/).

---

<div align="center">

**[⬆ Наверх](#-todo-scheduler)**

</div>