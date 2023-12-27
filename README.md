# PROMETHEUS/GRAFANA

Этот проект использует Prometheus/Grafana для измерения и отображения метрик, а также Alertmanager для отправки уведомлений, если количество вызовов конечной точки превышает заданное число.

## Начало работы

1. **Клонирование проекта**
   
   В терминале выполните команду:

    ```bash 
    git clone git@github.com:shuklarituparn/Prometheus_Metrics.git



2. **Переход в директорию проекта**

    Выполните команду:
        
        cd Prometheus_Metrics


3. **Сборка и запуск с помощью Docker**

Откройте терминал и выполните команды:

         docker compose build
         docker compose up
    


4. **Вход в Grafana**

Откройте `localhost:3000/login`, чтобы войти в панель Grafana. Используйте имя пользователя и пароль `admin`.

5. **Добавление источника данных**

Нажмите на "Добавить новый источник данных" и выберите Prometheus. В URL Prometheus добавьте `http://prometheus:9090` и нажмите "Сохранить".

6. **Создание новой панели управления**

Создайте новую панель управления и в метриках выберите нужную вам метрику. Нажмите "Выполнить запрос".

7. **Измерение метрик**

Теперь вы можете измерять метрики.

![Dashboard_Grafana](https://github.com/shuklarituparn/Prometheus_Metrics/assets/66947051/af756eec-6777-49cd-8de4-05a84cd6b517)

![Dashboard_Grafana_2](https://github.com/shuklarituparn/Prometheus_Metrics/assets/66947051/65342bba-2570-446b-b18b-6918ed016128)

# ENDPOINTS

Проект также включает две конечные точки под названием "cat" и "dog", которые отображают случайные изображения кошек или собак дня и увеличивают счетчик при каждом вызове.


![Cat_Endpoint](https://github.com/shuklarituparn/Prometheus_Metrics/assets/66947051/bf0861ad-1e14-4d21-9cc9-11d5f9974293)

![Dog_Endpoint](https://github.com/shuklarituparn/Prometheus_Metrics/assets/66947051/13fba459-bc6a-41ff-b2a3-49858ac95cdf)

Для доступа к ним откройте `localhost:8090/cat` или `localhost:8090/dog`.

# ALERTMANAGER

Проект также содержит Alertmanager, который активирует уведомление, если метрика превышает заданный лимит.

## Настройка вебхука

1. **Переход в директорию Alertmanager**

Выполните команду:
    
          cd alertmanager



2. **Конфигурация**

Отредактируйте файл конфигурации следующим образом:

```yaml
global:
  resolve_timeout: 5m
route:
  receiver: webhook_receiver
receivers:
  - name: webhook_receiver
    webhook_configs:
      - url: '<YOUR_WEBHOOK_URL>'
        send_resolved: false
```

Вместо ```<YOUR_WEBHOOK_URL>``` вставьте свой URL вебхука.

![Prometheus_Alert](https://github.com/shuklarituparn/Prometheus_Metrics/assets/66947051/a9edc0c5-950c-4196-8500-e54359accd8e)


![AlertManager_Alerts](https://github.com/shuklarituparn/Prometheus_Metrics/assets/66947051/2322649f-0618-4b13-a04c-90f47a7ed6e9)




