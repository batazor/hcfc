### ДОРОЖНАЯ КАРТА

#### v1.0 первая версия

- [x] Команда 
    -  [x] `generate` - схема управления генерацией команд 
    -  [x] `-o` выходной каталог 
    -  [x] `-f` путь к файлу значений `deploy.yaml` 
    -  [x] `-t` каталог шаблонов 
- [x] Генерация простой диаграммы из шаблонов 
    -  [x] Chart.yaml 
    -  [x] deploy.yaml 
    -  [x] service.yaml 
    -  [x] добавить поддержку разбора Helm 

#### v1.0.1 Добавить CI / Рефакторинг

- [x] Добавить CI 
    -  [x] Создать двоичный файл 
    -  [x] Создание образа докера 
    -  [x] GitHub Action 

#### v1.1.0 легко создать развернуть конфигурацию

- [x] Использовать собственный регистратор (для форматирования журналов)
- [x] Добавить команду `init` - Создать новую конфигурацию `deploy.yaml` 
    -  [x] Подтвердите создание нового конфига 
    -  [x] Напишите имя, описание, версию 
    -  [x] Развертывание базы генерации, обслуживание, вход (необязательно) 

#### v1.1.1 собачий корм

- [x] Используйте генератор диаграмм :-)
- [x] Генерация `values.yaml`

#### v1.1.4 Исправить синтаксический массив для шаблонов

- [x] От `{{ .Values.Deployment[0].Replicas }}` до `{{ (index .Values.Deployment 0).Replicas }}`

#### v1.2.0 Улучшение

- [x] Улучшить `deployment.yaml` 
    -  [x] Добавить переменную ENV (диапазон) 
- [x] Добавить секрет поддержки (простой)
- [] Улучшить `ingress.yaml` 
    -  [] Добавить секрет TLS 
    -  [] Добавить аннотацию (импорт из ngix-ingress?) 
    -  [] Добавить домен 
        -  [] Добавить патч (маршрут, URL) 
        -  [] Добавить бэкэнд (из сервиса) 
- [] Улучшить `Charts.yaml` 
    -  [] Добавить ключевые слова 
    -  [] Добавить домой 
    -  [] Добавить источники 
    -  [] Добавить mainteiners 
        -  [] Добавить поддержку `MAINTEINERS.md` 
    -  [] Добавить версию по умолчанию для apiVersion / appVersion 
- [] Создать `README.md` 
    -  [ ] описание 
    -  [] таблица с переменной ENV (имя, значение по умолчанию) 
- [] Пропустить комментарии в файле шаблона
- [] Добавить комментарий к `values.yaml`
- [] Улучшить шаблон Ingress
- [] Используйте ENV CI_COMMIT_TAG 
    -  [] Dockerfile 
    -  [] GitHub Action 

#### v1.2.1

- [] Добавить `_helpers.tpl`

#### v1.3.0 Добавить лучшие практики

- [] Добавить больше значков ;-)
- [] Добавить линтеры
- [] Проверьте линтера на действиях GitHub
- [] Используйте структуру k8s / helm

#### v1.3.1 Добавить первый тест

- [] Добавить тестовое создание `deployment.yaml`
- [] Добавить покрытие
- [] Добавить действие GitHub

#### v1.3.2 Добавить еще тест

- [] Добавить тестовое создание, `deployment`
- [] Добавить тест создать `service`
- [] Добавить тест создать `values.yaml`

#### v1.4.0 улучшить документы

- [] добавить пример для команды 
    -  [] генерировать 
    -  [ ] в этом 

#### v1.5.0 мониторинг

- [] Добавить healtcheck 
    -  [] добавить шаблон http (выбрать) 
        -  [] добавить привязку к порту 
- [] Добавить галочку 
    -  [] добавить шаблон http (выбрать) 
        -  [] добавить привязку к порту 
- [] Поддержка Прометей 
    -  [] Добавить шаблон 
        -  [ ] Пинг 
        -  [ ] проверка здоровья 
- [] Панель приборов для графана

#### v1.6.0 улучшить ресурс

- [] Улучшить развертывание 
    -  [] добавить ресурс 
    -  [] добавить nodeSelector 
    -  [] добавить securityContext 
- [] Добавить StatefulSet
- [] Добавить DaemonSet
- [] Добавить RBAC

#### v1.7.0 Улучшение NetworkPolicy

- [] Добавить NetworkPolicy
- [] Добавить пресет: 
    -  [] [Поддержка SCTP](https://kubernetes.io/docs/concepts/services-networking/network-policies/#sctp-support) 
    -  [] По умолчанию запрещен весь входящий и исходящий трафик 
    -  [] По умолчанию разрешить весь выходной трафик 
    -  [] По умолчанию запрещен весь исходящий трафик 
    -  [] По умолчанию разрешить весь входящий трафик 
    -  [] По умолчанию запрещен весь входящий трафик 

#### v1.8.0 Добавить [PodPriority](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/)

- [] Добавить `KubeSchedulerConfiguration`
- [] Добавить `PriorityClass`

#### v1.9.0 Добавить ServiceAccount

- [] Добавить ServiceAccount
- [] Улучшение развертывания, statefulset, daemonset 
    -  [] добавить serviceAccountName 

#### v2.0.0 Поддержка Giltab

- [] Добавить команду `init` 
    -  [] Добавить тип `GitLab` 
    -  [] Создать Dockerfile (как плагины) 
        -  [] NodeJS 
        -  [] Простой HTML 
        -  [] Голанг 
        -  [] Yii 
- [] Поколение gitlab-ci.yaml 
    -  [] Работа: 
        -  [] Построить Dockerfile 
        -  [] Перейдите в реестр 
        -  [] Создать диаграмму Хелма 
        -  [] Развернуть диаграмму 
- [] Поддержка ENV
- [] Добавить тестовый пример

#### v2.0.1 собачий корм

- [] Используйте генератор gitlab :-)

#### v3.0.0 Поддержка GitHub Action

- [] Обновить команду `init` 
    -  [] Добавить тип `GitHub` 

#### v3.0.1 собачий корм

- [] Используйте генератор gitlab :-)

#### V4.0.0

- [] Диаграмма зависимостей поддержки

#### v5.0.0

- [] Добавьте команду `edit` для значения edit `deploy.yaml`