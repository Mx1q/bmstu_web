<template>
  <div class="d-flex justify-content-center align-items-center center gap-3">
    <div class="flex flex-column gap-5 justify-content-between align-items-between">
          <span class="font-medium text-primary text-xl">
            Комментарии
          </span>
    </div>
  </div>

    <div class="card">
      <DataView :value="comments" paginator :rows="commentRows">
        <template #empty>
          <div class="d-flex justify-content-center align-items-center center">
            <div class="flex">
        <span class="font-medium text-primary text-xl">
          Комментарии не найдены
        </span>
            </div>
          </div>
        </template>

        <template #header v-if="userComment !== null">
          <template v-if="updatingComment === false">
            <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-3" :class="{ 'border-top-1 surface-border': index !== 0 }">
              <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                  <div class="surface-100 p-1" style="border-radius: 30px">
                    <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                        <span class="text-900 font-medium text-md">
                          {{ userComment.rating }}
                        </span>
                      <i class="pi pi-star-fill text-yellow-500"></i>
                    </div>
                  </div>
                  <div>
                    <div class="text-lg font-medium text-900 mt-2">
                      <template v-if="userComment.author==null">
                      </template>
                      <template v-else-if="userComment.author">
                        {{ userComment.author.name }}
                      </template>
                    </div>
                  </div>
                </div>
                <div class="flex flex-column md:align-items-end gap-5">
                    <span class="font-normal text-900">
                      {{ userComment.text }}
                    </span>
                </div>

                <div class="flex flex-column md:align-items-end">
                  <div>
                    <Button icon="pi pi-pencil"
                            style="border-top-left-radius: 8px; border-top-right-radius: 8px"
                            severity="success"
                            @click="startUpdatingComment" />
                  </div>
                  <div>
                    <Button icon="pi pi-trash"
                            style="border-bottom-left-radius: 8px; border-bottom-right-radius: 8px"
                            severity="danger"
                            @click="deleteUserComment" />
                  </div>
                </div>

              </div>
            </div>
          </template>
          <template v-else-if="updatingComment === true">
              <UpdateComment
                  :comment_id="userComment.id"
                  @updatingSuccess="commentUpdateSuccess"
                  @updatingCancelled="commentUpdateCancelled"
                  @updatingError="commentUpdateFailed"
              />
          </template>
        </template>
        <template #header v-if="userComment === null && isAuthValue === true">
          <CreateComment
              :salad_id="this.$route.params.id"
              @creatingSuccess="commentCreateSuccess"
              @creatingError="commentCreateFailed"
          />

        </template>

        <template #list="slotProps">
          <div class="grid grid-nogutter">
            <div v-for="(item, index) in slotProps.items" :key="index" class="col-12">
              <template v-if="item.author_id!==authUserId">
                <div class="flex flex-column sm:flex-row sm:align-items-center p-4 gap-3" :class="{ 'border-top-1 surface-border': index !== 0 }">
                  <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                    <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                      <div class="surface-100 p-1" style="border-radius: 30px">
                        <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                        <span class="text-900 font-medium text-md">
                          {{ item.rating }}
                        </span>
                          <i class="pi pi-star-fill text-yellow-500"></i>
                        </div>
                      </div>
                      <div>
                        <div class="text-lg font-medium text-900 mt-2">
                          <template v-if="item.author==null">
                          </template>
                          <template v-else-if="item.author">
                            {{ item.author.name }}
                          </template>
                        </div>
                      </div>
                    </div>
                    <div class="flex flex-column md:align-items-end gap-5">
                    <span class="font-normal text-900">
                      {{ item.text }}
                    </span>
                    </div>
                  </div>
                </div>
              </template>

            </div>
          </div>
        </template>
      </DataView>
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