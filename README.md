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
> cat rphost_*/*.log |
>awk -vORS= '{if($0 ~ /^[0-9][0-9]:/){ print "\n"$0 }else{ print "@"$0 }}' |
>grep -E ',LEAKS' |
>sed 's/@/\n/g' |
>tee ../logs.log
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
>cat rphost_*/*.log |
>sed 's/,Context,/,@,/g' |
>awk -vORS= '{if($0 ~ /^[0-9]{2}:[0-9]{2}\.[0-9]{6}-\d*,[^@]/){ print "\n"$0 }else{ print "$"$0 }}' |
>grep -E ',CALL.*,Context=' |
>sed 's/^.*,Context=//g' |
>sed 's/,[A-z]*=.*MemoryPeak=/MemoryPeak=/g' |
>awk -F 'MemoryPeak=' '{ count[$1] += 1; sumMemory[$1] += $2;} END { for(i in count){ print "Count: "count[i]", MemorySum: "sumMemory[i]" = "i"\n" }}' |
>sort -nrb -k 4 -t " " |
>sed -r 's/$/\n/g' |
>tee ../result.logs
>``` 
>
## 2. Администрирование
### 1. MS SQL
#### 1.1. Статистика
1. Посмотреть текущее состояние статистики:
```sql
SELECT objects.name AS objectName, stat.name AS statName, last_updated, rows, modification_counter   
FROM sys.objects AS objects
LEFT JOIN sys.stats AS stat  
ON objects.object_id = stat.object_id
CROSS APPLY sys.dm_db_stats_properties(stat.object_id, stat.stats_id) AS sp  
WHERE objects.name = '_InfoRg50';
```
**Главные показатели в этом запросе:**

**last_updated** - Дата и время последнего обновления объекта статистики.  
**modification_counter** - Общее количество изменений в начальном столбце статистики (на основе которого строится гистограмма) с момента последнего обновления статистики.

При настройках базы данных по умолчанию, обновление статики происходит автоматически (параметр AUTO_UPDATE_STATISTICS). Если он включен, оптимизатор запросов проверяет, устарела ли статистика перед выполнением запроса, и при необходимости обновляет ее. 

AUTO_UPDATE_STATISTICS_ASYNC - если установить данный параметр, запрос выполниться со старой статистикой. Параллельно будет запущен процесс пересчета статистики. 

Автообновление статистики применяется только для статистики помеченной сервером, как "устаревшая". Сервер ставит такую отметку по определенным критериям:

| Тип таблицы	            | Кардинальность таблицы (n)	| Пороговое значение повторной компиляции (количество модификаций)  |
| :---                    | :---:                       |                                                             ---:  |
|Временные процедуры      |n< 6                         |6                                                                  |
|Временные процедуры      |6 <= n<= 500                 |500                                                                |
|Постоянный               |n<= 500                      |500                                                                |
|Временная или постоянная |n> 500                       |`До MS SQL 2014 (12.x):` 500 + (0,20 * n) <br/> `Начиная с MS SQL 2016 (13.x) :` IN ( 500 + (0.20 * n), SQRT(1,000 * n) )

2. Обновить статистику:
- UPDATE STATISTICS - обновляет статистику для таблицы или индексированного представления.
- EXEC sp_updatestats - полное обновление статистики с выводом информации (Эта хранимая процедура проходит по вашей базе данных с помощью цикла WHILE и выполняет команду UPDATE STATISTICS).

Варианты обновления статистики:
- FULL SCAN - обновление статистики с полным сканированием всех строк. 
- SAMPLE 		- обновление статистики для указанного процента или количества строк (Update STATISTICS <объект> WITH SAMPLE 10 PERCENT).
		
> [!NOTE]
>	`MAXDOP = max_degree_of_parallelism` - ограничение числа процессоров используемых при параллельном выполнении планов.
>
>Параметр max_degree_of_parallelism может иметь одно из следующих значений:
>
>1 — ограничивает максимальное количество процессоров, используемых в параллельных операциях со статистиками, заданным или меньшим числом в зависимости от текущей рабочей нагрузки системы;
>
>0 - (по умолчанию)
В зависимости от текущей рабочей нагрузки системы использует реальное или меньшее число процессоров.
>
>**Полное сканирование (FULL SCAN)** - может выполняться как параллельная операция, что может значительно ускорить процесс. <br/>
>**Использованием выборки (SAMPLE)** - для возможно SQL Server 2016 (или выше) параллельное выполнение. Для версий ниже, только однопоточный режим.
>
>**Что бы установить режима параллелизма для конкретной операции пересчета статистики выполните команду:**
>```sql
>EXEC sp_configure 'max degree of parallelism', 16;  
>GO  
>RECONFIGURE WITH OVERRIDE;  
>GO 
> ```

#### 1.2. Индексы
Основная проблема индексов - **фрагметация**. Ядро СУБД автоматически изменяет индексы при каждом выполнении операций вставки, обновления или удаления в базовые данные. Спустя время возникает ситуация, когда для некоторых страниц индекса логический порядок, основанный на значении ключа, не совпадает с физическим порядком страниц индексов, что приведет к понижению производительности.

Дефрагментация индексов происходит при помощи операций - реорганизации или перестроения.

> [!CAUTION] Фрагментация индексов является актуальной проблемой только для систем у которых скорость <произвольного> чтения значительно медленнее скорости <последовательного> чтения. Обычно это сервера с HDD дисками. Для серверов на SSD дефрагментация не является проблемой.
>
>Часто люди неправильно полагают, что улучшение в скорости работы запросов связано с перестроением индекса, которое снизило фрагментацию и увеличило плотность страниц. Но на практике прирост скорости связан с обновлением статистики, которая автоматически происходит после перестроения индекса.
>
>Затраты на обновление статистики ресурсов являются незначительными по сравнению с перестроения индексами, и операция часто завершается в минутах. Перестроение индекса может занять несколько часов.

##### 1. Посмотреть текущий уровень фрагментации индексов:
```sql
SELECT 
  objects.name AS objectName, 
  indexes.name AS indexName,
  index_stats.index_type_desc, 
  index_stats.index_depth, 
  index_stats.avg_fragmentation_in_percent 
FROM sys.objects AS objects 
LEFT JOIN sys.indexes AS indexes
  ON objects.object_id = indexes.object_id
LEFT JOIN sys.dm_db_index_physical_stats(DB_ID(), NULL, NULL, NULL, NUll) AS index_stats
  ON indexes.index_id = index_stats.index_id
WHERE NOT indexes.name is NULL
ORDER BY 
  index_stats.avg_fragmentation_in_percent DESC, 
  objectName, 
  indexes.name;
```

##### 2. Реорганизация индекса (дефрагментация)
###### 2.1. Реорганизация всех индексов в таблице:
```sql
ALTER INDEX ALL ON <Table> REORGANIZE; 
```
>[!NOTE]Важно! Начиная с версии платформы 8.3.22 необходимо выполнять дефрагментацию индексов по следующему алгоритму:
>
>- До дефрагментации индекса необходимо включить страничные блокировки. Пример команды: 
>```sql
>ALTER INDEX index_name ON table_name SET (ALLOW_PAGE_LOCKS = ON, ALLOW_ROW_LOCKS = ON);
>```
>- Выполнить дефрагментацию.
>- Обратно выключить страничные блокировки. Пример команды: 
>```sql
>ALTER INDEX index_name ON table_name SET (ALLOW_PAGE_LOCKS = OFF, ALLOW_ROW_LOCKS = ON);
>```

##### 3. Перестройка индекса
```sql

```