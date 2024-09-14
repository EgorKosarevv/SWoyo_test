URL Shortener 

Этот проект представляет собой простой сервис для сокращения URL-адресов. Он позволяет пользователям отправлять длинные URL для их сокращения и получать исходный URL по короткому URL.

Требования:
Go 1.18 или выше
PostgreSQL (если используется база данных)


Структура проекта
Проект имеет следующую структуру:

SWOYO_test/

├── config/

   &emsp└── config.yml       # Конфигурационный файл

   └── config.go        # Файл для работы с конфигурацией

├── controllers/

   └── url_controller.go # Контроллеры для обработки запросов

├── database/

   └── db.go            # Подключение к базе данных

├── models/

   └── url.go           # Логика работы с URL

├── store/

   ├── db_store.go      # Хранилище данных в базе данных

   └── memory_store.go  # Хранилище данных в памяти

   └── store.go         # Интерфейс для работы с хранилищем

├── main.go              # Главный файл приложения

├── go.mod               # Модульные зависимости

├── go.sum  

├── .gitignore           # Файл игнорирования для Git

└── README.md            


Установка:

Клонируйте репозиторий:

      git clone https://github.com/ваш_репозиторий/SWOYO_test.git

      cd SWOYO_test

Установите зависимости:

      go mod tidy


Запуск приложения:

Запустите сервер с использованием базы данных:

      go run main.go -d

Или, если вы хотите использовать память вместо базы данных:

      go run main.go

Сервер будет слушать на порту 8080 по умолчанию. Вы можете изменить порт в коде, если необходимо.


Использование:

Создание короткого URL

Отправьте POST-запрос на / с параметром originalURL. Приложение вернет короткий URL.

Перенаправление по короткому URL

Отправьте GET-запрос на /<shortURL>. Приложение перенаправит вас на оригинальный URL.


Работа с базой данных:

Подключение к базе данных:


Используйте утилиту командной строки для подключения к вашей базе данных PostgreSQL.

Проверка содержимого базы данных:


Вы можете выполнять SQL-запросы для просмотра содержимого таблиц.

Управление таблицами:


Используйте SQL-запросы для изменения структуры таблиц или добавления новых таблиц.

Резервное копирование и восстановление базы данных

Используйте команды для резервного копирования и восстановления базы данных.

