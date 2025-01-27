<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputGroupAddon>
        <Rating v-model="rating" :cancel="false" />
      </InputGroupAddon>
      <InputText v-model="commentText" placeholder="Текст" />
      <Button @click="createComment" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px">
        Создать
      </Button>
    </InputGroup>
  </div>
</template>

<script>
import CommentModel from "@/models/CommentModel.js";
import CommentService from "@/services/comment.service.js";
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { ref } from 'vue';

export default {
  name: 'CreateCommentChild',

  props: {
    salad_id: String,
  },

  components: {
    Button,
    InputGroup,
    InputGroupAddon,
    InputText,
  },
  setup() {
    const commentText = ref('');
    const rating = ref(0)

    return {
      commentText,
      rating,
    };
  },
  created() {
  },
  methods: {
    createComment() {
      const comment = new CommentModel(this.rating, this.commentText, this.salad_id)

      CommentService.createComment(comment)
          .then(response => {
            this.creatingSuccess()
          })
          .catch(error => {
            this.creatingError(error.response.data.error)
          })
    },

    creatingSuccess() {
      this.$emit('creatingSuccess')
    },

    creatingError(errorText) {
      this.$emit('creatingError', errorText)
    }
  }
}

</script>