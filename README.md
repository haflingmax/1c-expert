# 1c-expert
Все, что должен знать эксперт.

## 1. Настройка технологического журнала (ТЖ):
  
Файл настроек ТЖ __<logcfg.xml>__ обычно располагается по адресу:  
  ### Windows:
  ```
  C:\Program Files\1cv8\conf\logcfg.xml                       - для всех версий платформы 
  C:\Program File\1cv8\x.x.xx.xxxx\bin\conf\logcfg.xml        - для отдельной версии
  C:\Users\<UserName>\AppData\Local\1C\1cv8\conf\logcfg.xml   - для конкретного пользователя (клиентский ТЖ)
  ```

  ### Linux:
  ```
  /opt/1cv8/conf/logcfg.xml                                   - для всех версий
  ``` 
Переопределить расположение фалов настроек можно в файле __<conf.cfg>__:
  ### Windos:
  ```
  C:\Program Files\1cv8\conf\conf.cfg                         - для всех версий платформы; 
  C:\Program File\1cv8\x.x.xx.xxxx\bin\conf\conf.cfg          - для отдельной версии;
  ```
### Linux:
  ```
  /opt/1cv8/conf/conf.cfg                                     - для всех версий
  ``` 