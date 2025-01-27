create or replace function strToUUIDs(strUUIDs text)
    returns table(
                     id uuid
                 )
as $$
begin
    if trim(strUUIDs) = '' then return;
    else
        return query
            select uuid(regexp_split_to_table)
            from regexp_split_to_table(strUUIDs, '\s+')
            group by regexp_split_to_table;
    end if;
end
$$ language plpgsql;
-- drop function if exists strToUUIDs(strUUIDs text);
-- select * from strToUUIDs('2516a0e8-38c9-4b38-bc4a-03f1f66409d7 17ff8c02-3d4d-465c-9173-940afa87b998');

create or replace function filterSaladsByIngredientIDs(ingredientIDs text)
returns table(
    id uuid,
    name varchar(32),
    authorId uuid,
    description varchar(32)
) as $$
    begin
        if trim(ingredientIDs) = '' then
        return query
        select *
        from saladrecipes.salad;
        else
        return query
        with
            matchCounts as (select recipeId, count(*) as matches
                             from saladRecipes.recipeIngredient
                             where ingredientId in (select * from strToUUIDs(ingredientIDs))
                             group by recipeId),
            totalIngredients as (select recipeId, count(*) as ingredientsCount
                                  from saladRecipes.recipeIngredient
                                  group by recipeId)
            select *
            from saladrecipes.salad
            where salad.id in (select saladId
                         from saladrecipes.recipe
                         where recipe.id in (select matchCounts.recipeid
                                      from matchCounts join totalIngredients on matchCounts.recipeId = totalIngredients.recipeId
                                      where matches = ingredientsCount));
        end if;
    end;
$$ language plpgsql;
-- drop function if exists filterSaladsByIngredientIDs(ingredientIDs text);
-- select * from filterSaladsByIngredientIDs('2516a0e8-38c9-4b38-bc4a-03f1f66409d7');

create or replace function filterSaladsByTypeIDs(typeIDs text)
    returns table(
                     id uuid,
                     name varchar(32),
                     authorId uuid,
                     description varchar(32)
                 ) as $$
begin
    if trim(typeIDs) = '' then
        return query
            select *
            from saladrecipes.salad;
    else
        return query
            select *
            from saladrecipes.salad
            where salad.id in (select saladid
                               from saladrecipes.typesofsalads
                               where typesofsalads.typeid in (select * from strToUUIDs(typeIDs)));
    end if;
end;
$$ language plpgsql;
-- drop function if exists filterSaladsByTypeIDs(typeIDs text);
-- select * from filterSaladsByTypeIDs('b1dea1e1-a5c5-4087-9961-887be1ec50fa');

create or replace function filterSalads(ingredIDs text, saladTypeIDs text, minRate float)
    returns table(
                     id uuid,
                     name varchar(32),
                     authorId uuid,
                     description varchar(32)
                 ) as $$
begin
    return query
        with sorted as (
            select *
            from filterSaladsByIngredientIDs(ingredIDs)
            intersect
            select *
            from filterSaladsByTypeIDs(saladTypeIDs))
        select *
        from sorted
        where sorted.id in (select saladId
                           from saladrecipes.recipe
                           where rating is null or rating > minRate);
end;
$$ language plpgsql;
-- drop function filterSalads(ingredIDs text, saladTypeIDs text, minRate float);
-- TODO: write an example of calling this filtering function
