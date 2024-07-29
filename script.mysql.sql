
create table users
(
    id_user integer not null auto_increment,
    date_add datetime default now(),
    date_update datetime default now(),
    name varchar(80) not null,
    email varchar(80) not null,
    primary key (id_user)
);


create table users_comments
(
    id_comment  int  not null primary key auto_increment,
    id_user     int  not null,
    date_add    datetime default CURRENT_TIMESTAMP null,
    date_update datetime default CURRENT_TIMESTAMP null,
    text        text                               null,
    status      tinyint  default 0                 null
);

create table config_empresas
(
    id                   int auto_increment
        primary key,
    razao_social         varchar(200) not null,
    nome_fantasia        varchar(200) not null,
    cnpj                 char(14)     not null,
    inscricao_estadual   char(20)     null,
    inscricao_municipal  varchar(20)  null,
    endereco_cep         char(10)     null,
    endereco_logradouro  varchar(80)  null,
    endereco_numero      char(10)     null,
    endereco_bairro      varchar(80)  null,
    endereco_complemento varchar(80)  null,
    endereco_municipio   varchar(80)  null,
    endereco_uf          char(2)      null,
    email                varchar(50)  null,
    site_oficial         varchar(150) null
);




