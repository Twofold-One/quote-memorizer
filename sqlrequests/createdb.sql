create table quotes (
    id serial primary key,
    author varchar(128) not null,
    quote text not null,
    created timestamp not null
);

create index idx_quotes_created ON quotes(created);