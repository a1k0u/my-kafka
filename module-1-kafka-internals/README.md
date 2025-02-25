## Модуль 1. Введение в Apache Kafka: продюсеры, консьюмеры, топики и партиции

### Как запустить?

Последовательный вызов make команд:

- `make step-1`
  - Запуск Kafka через Docker Compose
- `make step-2`
  - Создание топика `yandex-practicum` с тремя партициями и двумя репликами
- `make step-3`
  - Удаление всех созданных consumer группы
  - Сброс оффсета в топике во всех партициях
  - Запуск приложения с продюсером и двумя консюмерами

Удалить/остановить все созданные ресурсы можно через `make cleanup`

### Что делает приложение?

Запускается продюсер, записывает в топик `yandex-practicum`
структуру `Order`, состояющую из одного uint64 поля, с секундным интервалом.

Параллельно с ним работает два консюмера. Один иммитирует модель `push` (*получение новых данных сразу после отправки в топик*), а второй работает в стандартном для Kafka режиме `pull`.

По логам можем убедиться, что `push`-консюмер получает сообщение практически сразу после того, как продюсер отправляет данные в топик.
```
Producer: delivered message to yandex-practicum [1] [123 34 112 114 105 99 101 34 58 52 56 48 49 48 52 53 49 50 50 57 52 57 56 49 50 56 49 55 125] with 122 offset
Push Consumer: get from yandex-practicum [1] {4801045122949812817} with 122 offset
Producer: delivered message to yandex-practicum [0] [123 34 112 114 105 99 101 34 58 56 50 49 51 55 55 56 51 51 56 48 50 49 48 56 53 49 48 51 125] with 141 offset
Push Consumer: get from yandex-practicum [0] {8213778338021085103} with 141 offset
Producer: delivered message to yandex-practicum [2] [123 34 112 114 105 99 101 34 58 56 49 55 50 48 48 51 52 51 53 52 55 54 55 55 50 56 55 54 125] with 138 offset
Push Consumer: get from yandex-practicum [2] {8172003435476772876} with 138 offset
Producer: delivered message to yandex-practicum [0] [123 34 112 114 105 99 101 34 58 50 55 55 49 48 54 50 55 50 56 48 57 55 56 51 53 56 53 54 125] with 142 offset
Push Consumer: get from yandex-practicum [0] {2771062728097835856} with 142 offset
Producer: delivered message to yandex-practicum [0] [123 34 112 114 105 99 101 34 58 49 55 54 54 55 54 56 52 49 52 53 56 50 55 52 53 56 53 52 50 125] with 143 offset
Push Consumer: get from yandex-practicum [0] {17667684145827458542} with 143 offset
Producer: delivered message to yandex-practicum [2] [123 34 112 114 105 99 101 34 58 51 52 54 50 52 52 57 57 57 51 53 52 51 54 52 56 49 57 48 125] with 139 offset
Push Consumer: get from yandex-practicum [2] {3462449993543648190} with 139 offset
Producer: delivered message to yandex-practicum [0] [123 34 112 114 105 99 101 34 58 49 52 48 55 50 50 48 53 53 48 48 57 56 54 49 51 49 55 48 52 125] with 144 offset
Push Consumer: get from yandex-practicum [0] {14072205500986131704} with 144 offset
Producer: delivered message to yandex-practicum [1] [123 34 112 114 105 99 101 34 58 49 48 52 54 54 55 52 56 51 48 50 54 48 49 50 48 51 53 49 56 125] with 123 offset
Push Consumer: get from yandex-practicum [1] {10466748302601203518} with 123 offset
Pull Consumer: get from yandex-practicum [0] {8213778338021085103} with 141 offset
Pull Consumer: get from yandex-practicum [0] {2771062728097835856} with 142 offset
Pull Consumer: get from yandex-practicum [0] {17667684145827458542} with 143 offset
Pull Consumer: get from yandex-practicum [0] {14072205500986131704} with 144 offset
Pull Consumer: get from yandex-practicum [1] {4801045122949812817} with 122 offset
Pull Consumer: get from yandex-practicum [1] {10466748302601203518} with 123 offset
Pull Consumer: get from yandex-practicum [2] {8172003435476772876} with 138 offset
Pull Consumer: get from yandex-practicum [2] {3462449993543648190} with 139 offset
Producer: delivered message to yandex-practicum [2] [123 34 112 114 105 99 101 34 58 49 55 56 51 53 48 53 49 49 52 48 53 55 55 53 57 53 49 50 54 125] with 140 offset
Push Consumer: get from yandex-practicum [2] {17835051140577595126} with 140 offset
Producer: delivered message to yandex-practicum [0] [123 34 112 114 105 99 101 34 58 52 50 53 53 50 57 51 53 54 50 55 51 53 50 51 52 49 51 57 125] with 145 offset
Push Consumer: get from yandex-practicum [0] {4255293562735234139} with 145 offset
Producer: delivered message to yandex-practicum [1] [123 34 112 114 105 99 101 34 58 49 52 49 50 55 53 49 51 56 48 56 50 50 54 48 56 57 55 54 49 125] with 124 offset
Push Consumer: get from yandex-practicum [1] {14127513808226089761} with 124 offset
```
