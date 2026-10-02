# 🐹 GoLang: изучение основ

Репозиторий с учебными материалами для изучения языка **Go** по основным темам.
Каждый файл или пакет посвящён одной теме и содержит примеры кода с пояснениями.

Материалы пополняются по мере изучения языка.

---

## 📚 Содержание

| №  | Тема                        | Файл                                        | Что разобрано                                                                 |
|----|------------------------------|----------------------------------------------|--------------------------------------------------------------------------------|
| 1  | Hello, World                | [`1-hello.go`](./1-hello.go)                 | Структура программы, пакет `main`, функция `main`, вывод в консоль            |
| 2  | Циклы                       | [`2-loops.go`](./2-loops.go)                 | Цикл `for` во всех формах, `break`, `continue`, `range`                       |
| 3  | Срезы и map                 | [`3-slices.go`](./3-slices.go)               | Массивы и срезы, `len`/`cap`, `append`, срезы, `copy`, `map` и comma-ok idiom  |
| 4  | Структуры и указатели       | [`4-struct.go`](./4-struct.go)               | Структуры, поля, методы (value/pointer receiver), автоматическое взятие адреса |
| 5  | Обработка ошибок            | [`task/task.go`](./task/task.go)             | Тип `error`, `errors.New`, sentinel-ошибки, оборачивание `fmt.Errorf(%w)`, `errors.Is` |
| 6  | Модульное тестирование      | [`task/task_test.go`](./task/task_test.go)   | Пакет `testing`, обычные тесты и table-driven тесты, `t.Run`                   |
| 7  | Пакеты и структура проекта  | [`task/task.go`](./task/task.go)             | Разделение на `package task`/`package main`, экспортируемые имена, импорт своего пакета |
| 8  | Интерфейсы                  | [`task/task.go`](./task/task.go)             | Неявная реализация интерфейсов, `fmt.Stringer`, метод `String()`               |

---

## 🗂 Структура проекта

```
GoLang/
├── 1-hello.go        # Первая программа
├── 2-loops.go        # Циклы
├── 3-slices.go       # Срезы и map
├── 4-struct.go       # Структуры и указатели
├── scan-check.go     # Доп. эксперименты
├── task/
│   ├── task.go         # Доменная логика: Task, CreateTask, AddTask, String
│   └── task_test.go    # Тесты пакета task
├── go.mod             # Описание модуля
├── go.md              # Журнал прогресса для передачи контекста между сессиями
├── http.go            # методы обработки json http запросов
├── main.go            # Точка входа, использует пакет task
└── README.md
```

---

## 🚀 Как запустить

### Требования

- Установленный [Go](https://go.dev/dl/) (рекомендуется последняя стабильная версия)

Проверить версию:

```bash
go version
```

### Клонирование репозитория

```bash
git clone https://github.com/<your-username>/<repo-name>.git
cd <repo-name>
```

### Запуск программы

```bash
go run .
```

### Запуск тестов

```bash
# все тесты
go test ./...

# с подробным выводом
go test -v ./...

# конкретный тест по имени
go test -run TestИмяТеста -v
```

---

## 🛣 Планы

Темы, которые планируется добавить:

- [ ] HTTP и JSON: `net/http`, handler-функции, кодирование/декодирование JSON, статусы ответа
- [ ] REST API для задач: CRUD сначала в памяти
- [ ] PostgreSQL: `database/sql`, драйвер, миграции, repository
- [ ] `context`, middleware, graceful shutdown
- [ ] Горутины, каналы, `select`
- [ ] Пакет `sync` (Mutex, WaitGroup), race conditions (`go test -race ./...`)
- [ ] Работа с файлами
- [ ] Дженерики
- [ ] Pet-project: тесты handler/service/repository, линтеры, Docker, подготовка к собеседованиям

---

## 📖 Полезные ресурсы

- [Официальная документация Go](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/) — интерактивный тур по языку
- [Go by Example](https://gobyexample.com/) — примеры кода по темам
- [Effective Go](https://go.dev/doc/effective_go) — рекомендации по стилю

---

## 📝 Лицензия

Материалы свободны для использования в учебных целях.
