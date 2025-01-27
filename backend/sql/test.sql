-- drop schema saladrecipes cascade ;

insert into saladrecipes.ingredienttype(name)
values
    ('фрукт'),
    ('овощ'),
    ('мясо'),
    ('рыба'),
    ('молоко');

insert into saladrecipes.ingredient(id, name, calories, type)
values ('f1fc4bfc-799c-4471-a971-1bb00f7dd30a', 'яблоко', 1, (select id from saladrecipes.ingredienttype where name = 'фрукт'));

insert into saladrecipes.ingredient(name, calories, type)
values
    ('морковь', 2, (select id from saladrecipes.ingredienttype where name = 'овощ')),
    ('говядина', 3, (select id from saladrecipes.ingredienttype where name = 'мясо')),
    ('лосось', 4,  (select id from saladrecipes.ingredienttype where name = 'рыба')),
    ('молоко', 5, (select id from saladrecipes.ingredienttype where name = 'молоко'));

insert into saladrecipes.salad(id, name)
values ('fbabc2aa-cd4a-42b0-b68d-d3cf67fba06f', 'цезарь');

insert into saladrecipes.salad(name)
values
    ('овощной'),
    ('сезонный'),
    ('сельдь под шубой'),
    ('греческий');

insert into saladrecipes.saladtype(id, name)
values ('7e17866b-2b97-4d2b-b399-42ceeebd5480', 'зима');

insert into saladrecipes.saladtype(name)
values
    ('лето'),
    ('осень'),
    ('весна'),
    ('мясной');

insert into saladrecipes.typesofsalads(saladid, typeid)
values
    ((select id from saladrecipes.salad where name = 'цезарь'),
     (select id from saladrecipes.saladtype where name = 'зима')),

    ((select id from saladrecipes.salad where name = 'овощной'),
     (select id from saladrecipes.saladtype where name = 'лето')),
    ((select id from saladrecipes.salad where name = 'овощной'),
     (select id from saladrecipes.saladtype where name = 'зима')),

    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'лето')),
    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'зима')),
    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'весна')),
    ((select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.saladtype where name = 'осень')),

    ((select id from saladrecipes.salad where name = 'сельдь под шубой'),
     (select id from saladrecipes.saladtype where name = 'зима')),
    ((select id from saladrecipes.salad where name = 'сельдь под шубой'),
     (select id from saladrecipes.saladtype where name = 'мясной')),

    ((select id from saladrecipes.salad where name = 'греческий'),
     (select id from saladrecipes.saladtype where name = 'зима'));

insert into saladrecipes.modstatus(name)
values
    ('редактирование'),
    ('на модерации'),
    ('отклонено'),
    ('опубликовано'),
    ('снято с публикации');

insert into saladrecipes.recipe(id, saladid, status, numberofservings, timetocook)
values
    ('01000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'цезарь'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     1, 1),
    ('02000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'овощной'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     2, 2),
    ('03000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'сельдь под шубой'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     3, 3),
    ('04000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'сезонный'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     4, 4),
    ('05000000-0000-0000-0000-000000000000', (select id from saladrecipes.salad where name = 'греческий'),
     (select id from saladrecipes.modstatus where name = 'опубликовано'),
     5, 5);

insert into saladrecipes.measurement(id, name, grams)
values
    ('01000000-0000-0000-0000-000000000000', 'грамм', 1);

insert into saladrecipes.measurement(name, grams)
values
    ('чайная ложка', 1),
    ('штук', 1),
    ('килограмм', 1000);

insert into saladrecipes.recipeingredient(recipeid, ingredientid, measurement, amount)
values
    ((select id from saladrecipes.recipe where numberofservings = 1),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     1),

    ((select id from saladrecipes.recipe where numberofservings = 2),
     (select id from saladrecipes.ingredient where name = 'морковь'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     2),
    ((select id from saladrecipes.recipe where numberofservings = 2),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     3),

    ((select id from saladrecipes.recipe where numberofservings = 3),
     (select id from saladrecipes.ingredient where name = 'говядина'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     4),
    ((select id from saladrecipes.recipe where numberofservings = 3),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     5),

    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'говядина'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     6),
    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     7),
    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'морковь'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     8),
    ((select id from saladrecipes.recipe where numberofservings = 4),
     (select id from saladrecipes.ingredient where name = 'лосось'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     9),

    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'говядина'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     10),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'яблоко'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     11),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'морковь'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     12),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'лосось'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     13),
    ((select id from saladrecipes.recipe where numberofservings = 5),
     (select id from saladrecipes.ingredient where name = 'молоко'),
     (select id from saladrecipes.measurement where name = 'грамм'),
     14);

insert into saladrecipes.recipestep(name, description, recipeid, stepnum)
values
    ('first', 'first step', '01000000-0000-0000-0000-000000000000', 1),
    ('second', 'second step', '01000000-0000-0000-0000-000000000000', 2),
    ('third', 'third step', '01000000-0000-0000-0000-000000000000', 3),
    ('fourth', 'fourth step', '01000000-0000-0000-0000-000000000000', 4),
    ('fifth', 'fifth step', '01000000-0000-0000-0000-000000000000', 5);

insert into saladrecipes.user(name, email, login, password, role)
values
    ('user', 'user@mail.ru', 'user', '$2a$10$phPF85LGpSXa0RG5wRf0JORJDz/jeqmKXHBq672yZQwTxIUKf/6Na', 'user'),
    ('admin', 'admin@mail.ru', 'admin', '$2a$10$SIpA96ZIPUHT1m6O96B.fu.WILd1jc0CjOg/pTsqy0cIBo2yUnZA.', 'admin');