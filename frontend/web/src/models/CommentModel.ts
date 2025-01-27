class CommentModel {
    rating: number;
    text: string;
    salad_id: string;

    constructor(rating: number, text: string, salad_id: string) {
        this.text = text;
        this.rating = rating;
        this.salad_id = salad_id;
    }
}

export default CommentModel;