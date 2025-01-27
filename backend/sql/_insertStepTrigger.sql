-- Поиск максимального номера шага в рецепта с идентификатором rID
create or replace function maxStepNum(rId uuid)
    returns int
as $$
begin
    return (
        select case
                   when max(stepNum) is null then 0
                   else max(stepNum)
                   end as freeStep
        from saladrecipes.recipeStep
        where recipeid = rId);
end;
$$language plpgsql;

create or replace function insertStep()
    returns trigger
as $$
begin
    update saladrecipes.recipeStep
    set stepNum = maxStepNum(new.recipeid)+1
    where id = new.id;
    return new;
end;
$$ language plpgsql;

create trigger insertStepTrigger after insert on saladrecipes.recipeStep
    for each row execute procedure insertStep();


insert into saladrecipes.recipestep(name, description, recipeid, stepnum)
values ('testTrigger', 'testTrigger', '05000000-0000-0000-0000-000000000000', 0);

-- drop trigger if exists insertStepTrigger on saladrecipes.recipestep;

-- insert into saladrecipes.recipestep(name, description, recipeid, stepnum)
-- values('testingTrigger', 'testingTrigger', uuid('05000000-0000-0000-0000-000000000000'), 0);
