# go-netflowsquidlog-filter

## filter of logs received from the program go-netflow2squid

### How To Build

If you do not have a golang, please use the [installation instruction](https://golang.org/doc/install).

and next

    git clone https://github.com/Rid-lin/go-netflowsquidlog-filter.git
    cd go-netflowsquidlog-filter
    make

### How to use

    flow-cat ft-v05.YYYY-MM-DD.HHMMSS+Z | flow-print | go-netflow2squid > access_netflow_temp.log

    go-netflowsquidlog-filter -с go-netflowsquidlog-filter.json -in access_netflow_temp.log > access_netflow.log

After that, you need to use any **analyzer of squid** logs.

-------------------------------------------------

### Это фильтратор логов полученных отпрограммы go-netflow2squid

### Установка

Если у Вас не установлен Golang, пожалуйста воспользуйтесь [инструкцией по установке](https://golang.org/doc/install)

и далее

    git clone https://github.com/Rid-lin/go-netflowsquidlog-filter.git
    cd go-netflowsquidlog-filter
    make

После этого нужно использовать любой **анализатор логов Squid-а**.

### Как использовать

    flow-cat ft-v05.YYYY-MM-DD.HHMMSS+Z | flow-print | go-netflow2squid > access_netflow_temp.log

    go-netflowsquidlog-filter -с go-netflowsquidlog-filter.json -in access_netflow_temp.log > access_netflow.log

### Благодарности

- [Перевод документации по Squid](http://break-people.ru/)
- [Очень простое описание формата логов](https://wiki.enchtex.info/doc/squidlogformat)
