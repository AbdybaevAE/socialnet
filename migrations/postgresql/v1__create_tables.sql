create table if not exists lang
(
    lang_id bigserial not null primary key,
    lang_name varchar(255) not null, lang_code varchar(255) not null unique
);


insert into lang (lang_name, lang_value) values (
                                                 "english(UK)", "english_uk",
                                                );
insert into lang (lang_name, lang_value) values (
                                                    "english(USA)", "english_usa",
                                                );
