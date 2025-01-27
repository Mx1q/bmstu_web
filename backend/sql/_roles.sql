create role admin with
    nosuperuser
    nocreatedb
    noreplication
    login
    password 'admin'
    connection limit -1;

grant all privileges
    on all tables in schema saladrecipes, keywords
    to admin;

create role guest with
    nosuperuser
    nocreatedb
    nocreaterole
    noinherit
    nobypassrls
    noreplication
    login
    password 'guest'
    connection limit -1;

grant select
    on all tables in schema saladrecipes
    to guest;

grant insert
    on saladrecipes."user"
    to guest;

create role authorized with
    nosuperuser
    nocreatedb
    nocreaterole
    noinherit
    nobypassrls
    noreplication
    login
    password 'guest'
    connection limit -1;

grant select
    on all tables in schema saladrecipes
    to authorized;

grant insert (saladId, numberofservings, timetocook)
    on saladrecipes.recipe
    to authorized;

grant insert
    on saladrecipes.recipestep,
        saladrecipes.comment,
        saladrecipes.salad
    to authorized;
