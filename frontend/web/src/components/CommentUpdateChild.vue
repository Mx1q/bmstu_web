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
      <Button @click="updatingCancelled">
        отмена
      </Button>
      <Button @click="updateComment" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px">
        Обновить
      </Button>
    </InputGroup>
  </div>
</template>

<script lang="ts">
import CommentModel from "../models/CommentModel.ts";
import CommentService from "../services/comment.service.ts";
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { ref } from 'vue';

export default {
  name: 'UpdateCommentChild',

  props: {
    comment_id: String,
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
    this.fillInfo();
  },
  methods: {
    fillInfo() {
      CommentService.getById(this.comment_id)
          .then(response => {
            const comment = response.data.data.comment

            this.commentText = comment.text
            this.rating = comment.rating
          })
          .catch(error => {
            console.error(error)
          })
    },

    updateComment() {
      const comment = new CommentModel(this.rating, this.commentText)

      CommentService.updateComment(this.comment_id, comment)
          .then(response => {
            this.updatingSuccess()
          })
          .catch(error => {
            this.updatingError(error.response.data.error)
          })
    },

    updatingSuccess() {
      this.$emit('updatingSuccess')
    },

    updatingCancelled() {
      this.$emit('updatingCancelled')
      // this.updatingError("testing updating error")
    },

    updatingError(errorText: string) {
      this.$emit('updatingError', errorText)
    }
  }
}

</script>