# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1
4. [Настройте](https://docs.github.com/ru/authentication/connecting-to-github-with-ssh) доступ к своему Github-аккаунту
   по SSH-ключу.
5. Склонируйте себе репозиторий: `git clone <путь к вашему репозиторию>`
6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
   go mod tidy
```

7. Подготовьте два разных аккаунта TestBoard и укажите их учётные данные в файле `.env.override` в корне проекта.

   ```dotenv
   TEST_LOGIN=owner@example.com
   TEST_PASSWORD=owner_password
   TEST_BUYER_LOGIN=buyer@example.com
   TEST_BUYER_PASSWORD=buyer_password
   ```

   `TEST_LOGIN` и `TEST_PASSWORD` относятся к владельцу объявлений. `TEST_BUYER_LOGIN` и `TEST_BUYER_PASSWORD` относятся к покупателю, который добавляет объявления в избранное. При запуске `SetupSuite` загружает `.env`, затем `.env.override`. Оба аккаунта должны успешно входить в TestBoard. Для теста лимита требуется покупатель с пустым избранным. Используйте его для одного прогона за раз.

   Для полного запуска у владельца должно быть хотя бы одно ранее опубликованное объявление. Тест `myAdvertisement` проверяет описание и фото у всех объявлений из профиля, которые вернул API. Исходный тест поиска использует объявление `QAAvitoTestForPracticeLab` из каталога.

8. Для запуска всех тестов достаточно выполнить команду в корне проекта:

   ```bash
   go test -count=1 -v ./...
   ```

   Для проверки повторяемости запустите каждый тест три раза.

   ```bash
   go test -count=3 -v ./...
   ```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -v ./tests/scenarios/myAdvertisement/
```

## Логирование

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```
