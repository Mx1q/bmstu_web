class RecipeModel {
    number_of_servings: number;
    time_to_cook: number;
    salad_id: string;
    status: number;

    constructor(nos: number, ttc: number, saladId: string, status: number) {
        this.salad_id = saladId;
        this.number_of_servings = nos;
        this.time_to_cook = ttc;
        this.status = status;
    }
}

export default RecipeModel;