create extension if not exists pgcrypto;

create schema if not exists saladRecipes;
create schema if not exists keywords;

create table if not exists keywords.word (
    id uuid default gen_random_uuid() primary key,
    word varchar(32)
);

create table if not exists saladRecipes.user (
    id uuid default gen_random_uuid() primary key,
    name varchar(64) not null,
    email text not null check ( email like '%@%.%' ) unique, -- todo: unique?
    login varchar(64) not null unique, -- todo: unique?
    password varchar(256) not null,
    role varchar(25) not null default 'user'-- todo: add to ER diagram (course work) verify?
);

create table if not exists saladRecipes.saladType (
    id uuid default gen_random_uuid() primary key,
    name varchar(25) not null,
    description varchar(256) default ''
);

create table if not exists saladRecipes.salad (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    authorId uuid references saladRecipes.user(id),
    description varchar(32) not null default ''
);

create table if not exists saladRecipes.modStatus (
    id serial primary key,
    name varchar(32),
    description varchar(256)
);

create table if not exists saladRecipes.recipe (
    id uuid default gen_random_uuid() primary key,
    saladId uuid not null references saladRecipes.salad(id) on delete cascade,
    status int not null references saladRecipes.modStatus(id),
    numberOfServings int not null,
    timeToCook int,
    rating decimal(3, 1) default 0
);

create table if not exists saladRecipes.ingredientType (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    description varchar(256)
);

create table if not exists saladRecipes.ingredient (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    calories int not null,
    type uuid not null references saladRecipes.ingredientType(id)
);

create table if not exists saladRecipes.comment (
    id uuid default gen_random_uuid() primary key,
    author uuid not null references saladRecipes.user(id),
    salad uuid not null references saladRecipes.salad(id) on delete cascade,
    text text,
    rating int, check ( rating >= 1 ), check ( rating <= 5 ),
    unique (author, salad)
);

create table if not exists saladRecipes.measurement (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    grams int
);

create table if not exists saladRecipes.recipeStep (
    id uuid default gen_random_uuid() primary key,
    name varchar(32) not null,
    description text not null,
    recipeId uuid not null references saladRecipes.recipe(id) on delete cascade,
    stepNum int not null
);

-- Links between tables
create table if not exists saladRecipes.recipeIngredient (
    id uuid default gen_random_uuid() primary key,
    recipeId uuid not null references saladRecipes.recipe(id),
    ingredientId uuid not null references saladRecipes.ingredient(id),

    measurement uuid not null references saladRecipes.measurement(id) default '01000000-0000-0000-0000-000000000000',
    amount int not null default 1,
    unique(recipeId, ingredientId)
);

create table if not exists saladRecipes.typesOfSalads (
    id uuid default gen_random_uuid() primary key,
    saladId uuid not null references saladRecipes.salad(id) on delete cascade,
    typeId uuid not null references saladRecipes.saladType(id),
    unique (saladId, typeId)
);

-- TRIGGERS
create or replace function updateRating()
    returns trigger
as $$
begin
    update saladrecipes.recipe
    set
        rating = (select avg(rating)
                  from saladrecipes.comment
                  where salad = new.salad)
    where saladId = new.salad;
    return new;
end;
$$ language plpgsql;

create or replace trigger updateRatingTrigger after insert on saladrecipes.comment
    for each row execute procedure updateRating();
