-- auto-generated definition
create table users
(
  id        int auto_increment,
  username  varchar(100) null,
  password  varchar(100) null,
  user_sex  int          null,
  nick_name varchar(100) null,
  constraint user_id_uindex
    unique (id)
);

