# MySQL - Administration

## MySQL Server

```text
[root@host]# ps -ef | grep mysqld

[root@host]# cd /usr/bin
./safe_mysqld &

[root@host]# mysql -u root -p
Enter password: ********

mysql> exit
Byte
```

## Privileges

FLUSH PRIVILEGES tells the server to reload the grant tables.

* Select_priv
* Insert_priv
* Update_priv
* Delete_priv
* Create_priv
* Drop_priv
* Reload_priv
* Shutdown_priv
* Process_priv
* File_priv
* Grant_priv
* References_priv
* Index_priv
* Alter_priv

## Administrative Command

* USE Databasename
* SHOW DATABASES
* SHOW TABLES
* SHOW COLUMNS FROM tablename
* SHOW INDEX FROM tablename
* SHOW TABLE STATUS LIKE tablename\G
