create table if not exists lang
(
    lang_id bigserial not null primary key,
    lang_name varchar(255) not null,
    lang_code varchar(255) not null
);


insert into lang (lang_name, lang_code) values (
                                                   'english(UK)', 'english_uk'
                                               );
insert into lang (lang_name, lang_code) values (
                                                   'english(USA)', 'english_usa'
                                               );


create table if not exists locale(
    locale_id bigserial not null primary key,
    locale_translations jsonb not null default '{}'
);