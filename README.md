# go-netflow2squid

## Transpiler from netflow v5 format to default squid log format

### How To Build

If you do not have a golang, please use the [installation instruction](https://golang.org/doc/install).

and next

    git clone https://github.com/Rid-lin/go-netflow2squid.git
    cd go-netflow2squid
    make

### How to use

    flow-cat ft-v05.YYYY-MM-DD.HHMMSS+Z | flow-print | go-netflow2squid > access_netflow.log

After that, you need to use any **analyzer of squid** logs.

### Thanks

- [Translation of Squid documentation](http://break-people.ru)
- [Very simple log format description](https://wiki.enchtex.info/doc/squidlogformat)

-------------------------------------------------

### Это транспилятор из формата netflow v5 в формат логов Squid по-умолчанию

### Как установить

Если у Вас не установлен Golang, пожалуйста воспользуйтесь [инструкцией по установке](https://golang.org/doc/install)

и далее

    git clone https://github.com/Rid-lin/go-netflow2squid.git
    cd go-netflow2squid
    make

После этого нужно использовать любой **аналозатор логов Squid-а**.

### Как использовать

    flow-cat ft-v05.YYYY-MM-DD.HHMMSS+Z | flow-print | go-netflow2squid > access_netflow.log

### Благодарности

- [Перевод документации по Squid](http://break-people.ru/)
- [Очень простое описание формата логов](https://wiki.enchtex.info/doc/squidlogformat)
