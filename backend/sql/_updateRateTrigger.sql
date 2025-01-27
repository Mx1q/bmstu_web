create or replace function calcRate(saladId uuid)
    returns float
as $$
begin
    return (select avg(rating)
            from saladrecipes.comment
            where salad = saladId);
end;
$$ language plpgsql;

create or replace function updateRate()
    returns trigger
as $$
begin
    update saladRecipes.recipe
    set rating = calcRate(new.salad)
    where saladId = new.salad;
    return new;
end;
$$ language plpgsql;

create trigger updateRateTrigger after insert on saladrecipes.comment
    for each row execute procedure updateRate();
create trigger updateRateUpdate after update on saladrecipes.comment
    for each row execute procedure updateRate();
create trigger updateRateDelete after delete on saladrecipes.comment
    for each row execute procedure updateRate();
