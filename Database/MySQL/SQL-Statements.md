# SQL Statements - MySQL

## Database

```sql
CREATE DATABASE Tutorial;
DROP DATABASE Tutorial;
USE DATABASE Tutorial;
```

## Data Types

### Numeric Data Types

* INT 11 digits
  * TINYINT 3 digits
  * SMALLINT 5 digits
  * MEDIUMINT 9 digits
  * BIGINT 20 digits
* FLOAT(M, D) the display length (M) and the number of decimals (D) default to 10, 2. Decimal precision can go to 24 places for a FLOAT.
* DOUBLE(M, D) default to 16, 4. Decimal precision can go to 53 places for a DOUBLE.
* DECIMAL(M, D)

### Date and Time Types

* DATE
* DATETIME
* TIMESTAMP
* TIME
* YEAR(M)

### String Types

* CHAR(M): A fixed-length string between 1 and 255 characters in length.
* VARCHAR(M): A variable-length string between 1 and 255 characters in length. You must define a length when creating a VARCHAR field.
* BLOB or TEXT: A field with a maximum length of 65535 .characters
  * TINYBLOB or TINYTEXT
  * MEDIUMBLOB or MEDIUMBLOB
  * LONGBLOB or LONGBLOB
* ENUM: An enumeration, which is a fancy term for list.

## INSERT

```sql
INSERT INTO table_name (field1, field2, ..., fieldN)
VALUES (value1, value2, ..., valueN)
```

## SELECT

```sql
SELECT field1, field2, ..., fieldN
FROM table_name1, table_name2...
[WHERE Clause]
[OFFSET M ][LIMIT N]
```

## WHERE

```sql
SELECT field1, field2,...fieldN table_name1, table_name2...
[WHERE condition1 [AND [OR]] condition2...
```

* The WHERE clause is very useful when you want to fetch the selected rows from a table, especially when you use the MySQL Join. Joins are discussed in another chapter.
* It is a common practice to search for records using the Primary Key to make the search faster.
* If the given condition does not match any record in the table, then the query would not return any row.

## UPDATE

```sql
UPDATE table_name SET field1 = new-value1, field2 = new-value2
[WHERE Clause]
```

## DELEET

```sql
DELETE FROM table_name [WHERE Clause]
```

## LIKE

```sql
SELECT field1, field2,...fieldN table_name1, table_name2...
WHERE field1 LIKE condition1 [AND [OR]] filed2 = 'somevalue'
```

* If the SQL LIKE clause is used along with the % character, then it will work like a meta character (*) as in UNIX, while listing out all the files or directories at the command prompt. Without a % character, the LIKE clause is very same as the equal to sign along with the WHERE clause.

## ORDER BY

```sql
SELECT field1, field2,...fieldN table_name1, table_name2...
ORDER BY field1, [field2...] [ASC [DESC]]
```

## Join

```sql
SELECT a.tutorial_id, a.tutorial_author, b.tutorial_count
FROM tutorials_tbl a, tcount_tbl b
WHERE a.tutorial_author = b.tutorial_author;
```

## NULL Values

* IS NULL
* IS NOT NULL
* <=>: This operator compares values, which (unlike the = operator) is true even for two NULL values.

## Regexps

| Pattern | What the pattern matches |
| - | - |
| Pattern | What the pattern matches |
| ^ | Beginning of string |
| $ | End of string |
| . | Any single character |
| [...] | Any character listed between the square brackets |
| [^...] | Any character not listed between the square brackets |
| p1\|p2\|p3 | Alternation; matches any of the patterns p1, p2, or p3 |
| * | Zero or more instances of preceding element |
| + | One or more instances of preceding element |
| {n} | n instances of preceding element |
| {m,n} | m through n instances of preceding element |

## ALTER

### Drop, Add, Reposition a Column

```sql
ALTER TABLE testalter_tbl  DROP i;
ALTER TABLE testalter_tbl ADD i INT;
ALTER TABLE testalter_tbl ADD i INT FIRST;
ALTER TABLE testalter_tbl ADD i INT AFTER c;
```

### Alter a Column Definition or a Name

```sql
ALTER TABLE testalter_tbl MODIFY c CHAR(10);
ALTER TABLE testalter_tbl CHANGE i j BIGINT;
ALTER TABLE testalter_tbl MODIFY j BIGINT NOT NULL DEFAULT 100;
```

### Alter a Table Type

```sql
ALTER TABLE testalter_tbl TYPE = MYISAM;
```

### Rename a Table

```sql
ALTER TABLE testalter_tbl RENAME TO alter_tbl;
```

## Indexes

```sql
CREATE UNIQUE INDEX index_name ON table_name ( column1, column2,...);
```

### ALTER command to add and drop INDEX

```sql
ALTER TABLE tbl_name ADD PRIMARY KEY (column_list)
ALTER TABLE tbl_name ADD UNIQUE index_name (column_list)
ALTER TABLE tbl_name ADD INDEX index_name (column_list)
ALTER TABLE tbl_name ADD FULLTEXT index_name (column_list)
ALTER TABLE testalter_tbl ADD INDEX (c);
ALTER TABLE testalter_tbl DROP INDEX (c);
```

### ALTER Command to add and drop the PRIMARY KEY

```sql
ALTER TABLE testalter_tbl MODIFY i INT NOT NULL;
ALTER TABLE testalter_tbl ADD PRIMARY KEY (i);
ALTER TABLE testalter_tbl DROP PRIMARY KEY;
```

## Temporary Tables

```sql
CREATE TEMPORARY TABLE table_name;
DROP TABLE table_name;
```

## Clone Tables

```sql
SHOW CREATE TABLE table_name \G;
```

## Using Sequences

* AUTO_INCREMENT = 0

### Renumbering an Existing Sequence

```sql
ALTER TABLE insect DROP id;
ALTER TABLE insect ADD id INT UNSIGNED NOT NULL AUTO_INCREMENT FIRST,
ADD PRIMARY KEY (id);
```

## Handling Duplicates

### Preventing Duplicates from Occurring in a Table

* PRIMARY KEY or UNIQUE
* INSERT IGNORE
* REPLACE INTO

### Counting and Identifying Duplicates

```sql
SELECT COUNT(*) as repetitions
FROM person_tb1
GROUP BY last_name, first_name
HAVING repetitions > 1;
```

### Eliminating Duplicates from a Query Result

```sql
SELECT DISTINCT last_name, first_name FROM person_tbl ORDER BY last_name;
SELECT last_name, first_name FROM person_tbl GROUP BY (last_name, first_name);
```

### Removing Duplicates Using Table Replacement

```sql
CREATE TABLE tmp SELECT last_name, first_name, sex
FROM person_tbl;
GROUP BY (last_name, first_name);

DROP TABLE person_tbl;
ALTER TABLE tmp RENAME TO person_tbl;
```

or

```sql
ALTER IGNORE TABLE person_tbl ADD PRIMARY KEY (last_name, first_name);
```

## Database Export

```sql
SELECT * FROM tutorials_tbl INTO OUTFILE '/tmp/tutorials.txt';
```

### Exporting Tables as Raw Data

```text
mysqldump -u root -p TUTORIALS tutorials_tbl > dump.txt
mysqldump -u root -p TUTORIALS > database_dump.txt
mysqldump -u root -p --all-databases > database_dump.txt
```

### Copying Tables or Databases to Another Host

```text
mysqldump -u root -p database_name \
    mysql -h other-host.com database_name
```

## Database Import

```sql
LOAD DATA LOCAL INFILE 'dump.txt' INTO TABLE mytbl;
```

### Importing Data with mysqlimport

```text
mysqlimport -u root -p --local database_name dump.txt
mysqlimport -u root -p --local --fields-terminated-by = ":" \
   --lines-terminated-by = "\r\n"  database_name dump.txt
mysqlimport -u root -p --local --columns=b,c,a \
   database_name dump.txt
```
