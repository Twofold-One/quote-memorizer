-- todo
insert into quotes (author, quote, created) values (
    "Oscar Wilde",
    "Be yourself; everyone else is already taken.",
    cast(now() at time zone 'utc' as date)
)