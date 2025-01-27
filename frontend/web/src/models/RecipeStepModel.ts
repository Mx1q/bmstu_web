class RecipeStepModel {
    recipe_id: string;
    name: string;
    description: string;
    step_num: number;

    constructor(recipeId: string, name: string, description: string, step_num: number) {
        this.recipe_id = recipeId
        this.name = name
        this.description = description
        this.step_num = step_num
    }
}

export default RecipeStepModel;