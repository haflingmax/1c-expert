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
#### 1. Перерасход по памяти (утечки памяти, большие запросы).

 Настройки ТЖ.
 ```xml
<?xml version="1.0" encoding="UTF-8"?>
<config xmlns="http://v8.1c.ru/v8/tech-log">
  <log location="C:/v8/logs/" history="4">
    <event>
      <eq property="Name" value="ATTN"/>
    </event>
    <event>
      <eq property="Name" value="PROC"/>
    </event>
    <event>
      <eq property="Name" value="CLSTR"/>
    </event>
    <property name="all"/>
  </log>
</config>
 ```
> [!NOTE]
> Примеры строк ТЖ:
> 
>#### 1. Принудительное завершение rphost пользователем из диспетчера задач, через kill или иным способом: 
> 
>**PROC**
>
>`22:13.201000-0,PROC,0,level=INFO,process=ragent,OSThread=32644,Txt='Supervision time expired. ProcessID=d6ce89a2-85a6-4a2e-8940-ae9a4b6554e1 pid=19348 started=1955846464.'`
>
>`22:14.014007-2,PROC,0,level=INFO,process=ragent,OSThread=12880,Txt='Run process. Prog=C:\Program Files\1cv8\8.3.25.1394\bin\rphost.exe, Command=("C:\Program Files\1cv8\8.3.25.1394\bin\rphost.exe" -range 1560:1591 -reghost haflingmax -regport 1541 -pid d95eae3b-2bb2-487e-b3b1-f0dabd111392 -fromsrvc), success, pid=33216'`
> 
>В событии \<PROC\> нет явного указания на аварийное завершение процесса. Однако, мы можем увидеть сообщение 'Supervision time expired' говорящее, что ragent больше не видит rphost, а так же сообщение 'Run process' сигнализирующие о запуске нового рабочего процесса. Можно предположить, что перед этим случилось принудительное завершение предыдущего процесса.  
>
>**ATTN**
>
>`22:16.982007-0,ATTN,0,level=WARNING,process=ragent,OSThread=39564,Descr=Process finished,Url=tcp://haflingmax:1560,AgentUrl=,ProcessId=d6ce89a2-85a6-4a2e-8940-ae9a4b6554e1,Pid=19348`
>
>Данное событие явно указывает, что процесс завершен: 'Process finished'. Бзе пояснения причины. 
>
>**CLSTR**
>
>`59:03.506007-0,CLSTR,2,level=INFO,process=rmngr,p:processName=RegMngrCntxt,p:processName=ServerJobExecutorContext,OSThread=29384,t:clientID=7546,t:applicationName=AgentProcess,t:computerName=haflingmax,Event=Process requirements changed,Obsolete=(1):94b83924-893b-481c-902d-2d0ac77a147c;`
>
>`59:03.506008-0,CLSTR,2,level=INFO,process=rmngr,p:processName=RegMngrCntxt,p:processName=ServerJobExecutorContext,OSThread=29384,t:clientID=7546,t:applicationName=AgentProcess,t:computerName=haflingmax,Event=Unregister rphost,Txt=processID:94b83924-893b-481c-902d-2d0ac77a147c hostName:haflingmax`
>
>`59:04.584007-0,CLSTR,2,level=INFO,process=rmngr,p:processName=RegMngrCntxt,p:processName=ServerJobExecutorContext,OSThread=29384,t:clientID=7546,t:applicationName=AgentProcess,t:computerName=haflingmax,Event=Process deficit detected,Host=haflingmax,Connections=0,Infobases=0,Deficit=1`
>
>`59:04.584008-0,CLSTR,2,level=INFO,process=rmngr,p:processName=RegMngrCntxt,p:processName=ServerJobExecutorContext,OSThread=29384,t:clientID=7546,t:applicationName=AgentProcess,t:computerName=haflingmax,Event=Process requirements changed,Registered=(1):402768f0-90a4-4087-91ad-40cc6e027a51(haflingmax);`
>
>`59:04.584009-0,CLSTR,2,level=INFO,process=rmngr,p:processName=RegMngrCntxt,p:processName=ServerJobExecutorContext,OSThread=29384,t:clientID=7546,t:applicationName=AgentProcess,t:computerName=haflingmax,Event=Register rphost,Txt=processID:402768f0-90a4-4087-91ad-40cc6e027a51 hostName:haflingmax`
>
>В событии CLSTR так же можно наблюдать завершение rphost, через событие 'Unregister rphost' с последующей регистрацией нового рабочего процесса.
>В рамках перезапуска rphost, событие CLSTR ничего ценного не несет. Вся нужная информация есть в событиях PROC и ATTN. 
>
>#### 2. Завершение rphost по причине перерасхода/нехватки ресурсов:
>
>**PROC**
>
>Важные события: 
>- Abandoned process was alive too long time ‑ Рабочий процесс «завис» в памяти.
>- Process excess memory limit ‑ Рабочий процесс превысил ограничения по памяти.
>- Process has generated too big amount of exceptions ‑ Рабочий процесс формирует очень большое количество ошибок или исключений.
>- Process not respond ‑ Рабочий процесс не отвечает.
>- Process will be killed ‑ Процесс будет принудительно завершен.
>- Memory shortage detected — Свободной оперативной памяти осталось меньше безопасного расхода за один вызов.
>- Memory exceeded critical limit - Память процессов на сервере превысила критический объём
>
>`37:03.085001-0,ATTN,2,level=WARNING,process=rmngr,OSThread=29384,t:clientID=17951,t:applicationName=AgentProcess,t:computerName=haflingmax,Descr=Memory exceeded critical limit,ServerId=4bbfcd0e-2348-45cd-ba16-d7c3de52f3d4,Host=haflingmax,MemoryLimits='[tempAllowed: 13695641190 tempAllowedTimeLimit: 300 sec, criticalMem: 16263573913]',TotalMemory=17119551488,ExcessStartTime=20241007223307,ExcessDurationSec=236`
>
>`35:52.983001-0,ATTN,2,level=WARNING,process=rmngr,OSThread=38532,t:clientID=17931,t:applicationName=AgentProcess,t:computerName=haflingmax,Descr=Memory shortage detected,ServerId=4bbfcd0e-2348-45cd-ba16-d7c3de52f3d4,Host=haflingmax,FreeMemory=354754560,SafeLimit=1369564119`
>
>
>`37:08.107008-0,ATTN,0,level=WARNING,process=ragent,OSThread=33532,Descr=Process finished,Url=tcp://haflingmax:1560,AgentUrl=,ProcessId=2b12eb0e-9c56-408e-bdab-cb2e93a2eaa2,Pid=31172`
>
>В данном случае, мы видим так же завершение процесса, но по вполне понятным причинам. Перерасход оперативной памяти. 
>
>#### 2.1 Поймать утечки памяти
>
>Настройки ТЖ:
>```xml
><?xml version="1.0" encoding="UTF-8"?>
><config xmlns="http://v8.1c.ru/v8/tech-log">
> <log>
>  <event>
>   <eq property="Name" value="CALL">
>  </event>
>  <event>
>   <eq property="Name" value="LEAKS">
>  </event>
>  <property name="all"/>
> </log>
> <leaks collect="1">
>  <point call="server"/>
> </leaks>
></config>
>```
>Скрипт поиска утечек памяти с выгрузкой результат в файл:
>
>```bash
> cat rphost_*/*.log | awk -vORS= '{if($0 ~ /^[0-9][0-9]:/){ print "\n"$0 }else{ print "@"$0 }}' | grep -E ',LEAKS' | sed 's/@/\n/g' | tee ../logs.log
>```
>
>В результате мы можем увидеть сообщения типа (сообщение приведено не полностью, так как правило очень объемное):
>
>```bsl
>09:10.883004-0,LEAKS,1,level=DEBUG,process=rphost,OSThread=38724,t:clientID=21,t:applicationName=BackgroundJob,t:computerName=haflingmax,t:connectID=1156,Descr='
>KeyAndValue:
>ОбщийМодуль.УтечкиПамяти.Модуль : 4 : СоздатьУтечкуМин();
>	ОбщийМодуль.УтечкиПамяти.Модуль : 927 : Структура2 = Новый Структура("СтруктураЦикла, БольшойМассив", Неопределено, БольшойМассив);
>
>KeyAndValue:
>ОбщийМодуль.УтечкиПамяти.Модуль : 4 : СоздатьУтечкуМин();
>	ОбщийМодуль.УтечкиПамяти.Модуль : 927 : Структура2 = Новый Структура("СтруктураЦикла, БольшойМассив", Неопределено, БольшойМассив);
>
>KeyAndValue:
>ОбщийМодуль.УтечкиПамяти.Модуль : 4 : СоздатьУтечкуМин();
>	ОбщийМодуль.УтечкиПамяти.Модуль : 926 : Структура1 = Новый Структура("СтруктураЦикла, БольшойМассив", Неопределено, БольшойМассив);
>```
>
>Так же не будет лишним проверить, какие события в системе потребляют наибольшее количество памяти. Скрипт парсинга ТЖ:
>
>```bash
>cat rphost_10016/*.log | awk -vORS= '{if($0 ~ /^[0-9][0-9]:/){ print "\n"$0 }else{ print "@"$0 }}' | grep -e ',CALL.*processName=v83.*Module=' | sed 's/^.*,Module=//' | sed 's/,Method=/:/' | sed 's/,.*Memory=/,/' | awk -F ',' '{ sumMemory[$1] += $2; } END { for(i in sumMemory) { print i": "sumMemory[i]/1000000 } }' | sort -nrb -k 2 | sed 's/@/\n/g' | tee ../logs.log
>```
 
>