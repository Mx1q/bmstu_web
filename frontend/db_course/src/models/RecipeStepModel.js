class RecipeStepModel {
    constructor(recipeId, name, description, step_num) {
        this.recipe_id = recipeId
        this.name = name
        this.description = description
        this.step_num = step_num
    }
}

export default RecipeStepModel;