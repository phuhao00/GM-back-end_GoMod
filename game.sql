-- auto-generated definition
create table games
(
    id               int auto_increment,
    name             varchar(200) null,
    price            int          null,
    commentId        int          null,
    downloadQuantity bigint       null,
    score            int          null,
    url              varchar(100) null,
    supplierId       int          null,
    constraint game_id_uindex
        unique (id)
);

