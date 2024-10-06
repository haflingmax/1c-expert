# 1c-expert
Все, что должен знать эксперт.

## 1. Настройка технологического журнала (ТЖ):
### 1.1 Файлы ТЖ:  
#### Файл настроек ТЖ __<logcfg.xml>__ обычно располагается по адресу:  
  #### Windows:
  ```
  C:\Program Files\1cv8\conf\logcfg.xml                       - для всех версий платформы 
  C:\Program File\1cv8\x.x.xx.xxxx\bin\conf\logcfg.xml        - для отдельной версии
  C:\Users\<UserName>\AppData\Local\1C\1cv8\conf\logcfg.xml   - для конкретного пользователя (клиентский ТЖ)
  ```

  #### Linux:
  ```
  /opt/1cv8/conf/logcfg.xml                                   - для всех версий
  ``` 
### Переопределить расположение фалов настроек можно в файле __<conf.cfg>__:
  #### Windos:
  ```
  C:\Program Files\1cv8\conf\conf.cfg                         - для всех версий платформы; 
  C:\Program File\1cv8\x.x.xx.xxxx\bin\conf\conf.cfg          - для отдельной версии;
  ```
  #### Linux:
  ```
  /opt/1cv8/conf/conf.cfg                                     - для всех версий
  ``` 

### 1.2 Настройки ТЖ:
[Подробная документация по logcfg.xml на ИТС](https://its.1c.ru/db/v8325doc#bookmark:adm:TI000000393)
#### Полный ТЖ:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<config xmlns="http://v8.1c.ru/v8/tech-log">
  <log location="C:/v8/logs/" history="4">
    <event>
      <ne property="Name" value="" />
    <event>
    <property name="all">
  </log>
</config>
```
#### Блокировки:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<config xmlns="http://v8.1c.ru/v8.tech-log">
  <log location="C:/v8/logs/" history="4">
    <event>
      <eq property="Name" value="EXCP"/>
    </event>
    <event>
      <eq property="Name" value="TLOCK"/>
    </event>
    <event>
      <eq property="Name" value="TTIMEOUT"/>
    </event>
    <event>
      <eq property="Name" value="TDEADLOCK"/>
    </event>
    <event>
      <eq property="Name" value="SDBL"/>
      <eq property="Func" value="BeginTransaction"/>
    </event>
    <event>
      <eq property="Name" value="SDBL"/>
      <eq property="Func" value="RollbackTransaction"/>
    </event>
    <event>
      <eq property="Name" value="SDBL"/>
      <eq property="Func" value="CommitTransaction"/>
    </event>
    <property name="all"/>
  </log>
</config>
```

#### Аварийное завершение процессов __<rphost__>:
##### Основные причины:
 1. Перерасход по памяти (утечки памяти, большие запросы)
 ```xml
<?xml version="1.0" encoding="UTF-8"?>
<config xmlns="http://v8.1c.ru/v8/tech-log">
  <log location="C:/v8/logs/" history="4">
    <event>
      <eq property="Name" value="EXCP"/>
    </event>
    <event>
      <eq property="Name" value="ADMIN"/>
    </event>
    <event>
      <eq property="Name" value="ATTN"/>
    </event>
    <event>
      <eq property="Name" value="PROC"/>
    </event>
    <property name="all"/>
  </log>
</config>
 ```
> [!NOTE]
> Примеры строк ТЖ:
> 
>1. Принудительное завершение rphost из диспетчера задач пользователем или по иной причине: 
> 
>**PROC**
>
>`22:13.201000-0,PROC,0,level=INFO,process=ragent,OSThread=32644,Txt='Supervision time expired. ProcessID=d6ce89a2-85a6-4a2e-8940-ae9a4b6554e1 pid=19348 started=1955846464.'`
>
>`22:14.014007-2,PROC,0,level=INFO,process=ragent,OSThread=12880,Txt='Run process. Prog=C:\Program Files\1cv8\8.3.25.1394\bin\rphost.exe, Command=("C:\Program Files\1cv8\8.3.25.1394\bin\rphost.exe" -range 1560:1591 -reghost haflingmax -regport 1541 -pid d95eae3b-2bb2-487e-b3b1-f0dabd111392 -fromsrvc), success, pid=33216'`
> 
>В событии \<PROC\> нет явного указания на аварийное завершение процесса. Однако, мы можем увидеть сообщение 'Supervision time expired' говорящее, что ragent больше не видит rphost, а так же сообщение 'Run process' сигнализирующие о запуске нового рабочего процесса. Можно предположить, что перед этим случилось аварийное завершение предыдущего процесса. 
>
>**ATTN**
>
>`22:16.982007-0,ATTN,0,level=WARNING,process=ragent,OSThread=39564,Descr=Process finished,Url=tcp://haflingmax:1560,AgentUrl=,ProcessId=d6ce89a2-85a6-4a2e-8940-ae9a4b6554e1,Pid=19348`