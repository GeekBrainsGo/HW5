# ДЗ

Задания со звёздочкой не обязательны.
### 1. * Напишите функции для создания, редактирования и удаления данных в БД.
### 2. * Перепишите свой сервис, используя BeeGo (создайте контроллеры и напишите алгоритмы для работы с данными посредством веб-сайта).
### 3. * Сделайте отдельную ветку в Git и перепишите работу с БД с использованием BeeGo orm.
### 4. Перевести ваш блог на одну из ORM: `gorm` || `beego-orm` || `sqlboiler` (рекомендую sqlboiler)

# Пакеты

1) https://github.com/volatiletech/sqlboiler - SQL Boiler
2) https://github.com/jinzhu/gorm - Gorm
3) https://github.com/astaxie/beego - BeeGo

# Статьи

1) https://4gophers.ru/articles/veb-prilozhenie-s-beego/ - Как писать веб-приложения на BeeGo
2) https://habr.com/ru/post/444022/ - Почему BeeGo уже всё

# Подсказки

1) Не забудьте добавить $GOPATH/bin в переменную окружения PATH
2) Перегенерация моделей для boiler: `sqlboiler mysql -c sqlboiler.yaml --wipe`
3) Пример конфига для бойлера в репозитории