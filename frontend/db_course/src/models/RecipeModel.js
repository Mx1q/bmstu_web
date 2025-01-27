class RecipeModel {
    constructor(nos, ttc, saladId, status) {
        this.salad_id = saladId;
        this.number_of_servings = nos;
        this.time_to_cook = ttc;
        this.status = status;
    }
}

export default RecipeModel;