<template>
<!--  FIXME не готовый компонент, хотя надо было бы в salad уже слишком сложно...-->
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

  <script>
  import CommentModel from "@/models/CommentModel.js";
  import CommentService from "@/services/comment.service.js";
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

      updatingError(errorText) {
        this.$emit('updatingError', errorText)
      }
    }
  }

  </script>





























<!--TODO-->
<!--TODO-->
<!--TODO-->
<!--TODO-->
<!-- TODO: снизу предыдущая версия salad-->
<template>
  <Toast />

  <!--  <UpdateComment-->
  <!--      :comment_id="userComment.id"-->
  <!--      @updatingSuccess="commentUpdateSuccess"-->
  <!--      @updatingCancelled="commentUpdateCancelled"-->
  <!--      @updatingError="commentUpdateFailed"-->
  <!--  />-->

  <!--  <template v-if="editingMode === true && this.salad.author_id === authUserId">-->

  <!--    <UpdateSaladChild-->
  <!--        :salad_id="salad.id" />-->
  <!--    -->
  <!--  </template>-->

  <!--  <template v-if="editingMode === false">-->








  <div class="d-flex justify-content-center align-items-center center">
    <div class="flex">
          <span class="font-medium text-primary text-xl">
            {{ salad.name }}
          </span>
    </div>
  </div>

  <div class="d-flex">
    <div class="flex-row">
      <div class="text-lg font-medium text-900 mt-2">
        Автор: {{ author.name  }}
      </div>
      <div class="text-lg font-medium text-900 mt-2">
        Количество порций: {{ recipe.number_of_servings }}, время приготовления {{ recipe.time_to_cook }} минут
      </div>
      <div class="text-lg font-medium text-900 mt-2">
        {{ salad.description  }}
      </div>
    </div>
  </div>

  <DataView :value="ingredients" :layout="'grid'">
    <template #empty>
      <div class="d-flex justify-content-center align-items-center center">
        <div class="flex">
          <span class="font-medium text-primary text-xl">
            Необходимые ингредиенты не найдены :(
          </span>
        </div>
      </div>
    </template>

    <template #grid="slotProps">
      <div class="grid grid-nogutter">
        <div class="col-12 sm:col-2 md:col-2 xl:col-2 p-2">
          <div
              class="p-2 border-0 surface-border surface-card border-round flex flex-column"
          >
            <div
                class="flex flex-row justify-content-center align-items-start gap-2"
            >
              <div class="text-lg font-medium text-primary">
                Ингредиенты:
              </div>
            </div>
          </div>
        </div>

        <div
            v-for="(item, index) in slotProps.items"
            :key="index"
            class="col-12 sm:col-2 md:col-2 xl:col-2 p-2"
        >
          <div
              class="p-2 border-1 surface-border surface-card border-round flex flex-column"
          >
            <div
                class="flex flex-row justify-content-center align-items-start gap-2"
            >
              <div class="text-md font-medium text-900">
                <template v-if="item==null">
                </template>
                <template v-else-if="item.measurement">
                  {{ item.name }} ({{ item.count }} {{ item.measurement.name }})
                </template>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </DataView>

  <DataView :value="saladTypes" :layout="'grid'">
    <template #empty>
      <div class="d-flex justify-content-center align-items-center center">
        <div class="flex">
          <span class="font-medium text-primary text-xl">
            Типы рецепта не указаны :(
          </span>
        </div>
      </div>
    </template>

    <template #grid="slotProps">
      <div class="grid grid-nogutter">
        <div class="col-12 sm:col-2 md:col-2 xl:col-2 p-2">
          <div
              class="p-2 border-0 surface-border surface-card border-round flex flex-column"
          >
            <div
                class="flex flex-row justify-content-center align-items-start gap-2"
            >
              <div class="text-lg font-medium text-primary">
                Типы салата:
              </div>
            </div>
          </div>
        </div>

        <div
            v-for="(item, index) in slotProps.items"
            :key="index"
            class="col-12 sm:col-1 md:col-1 xl:col-1 p-2"
        >
          <div
              class="p-2 border-1 surface-border surface-card border-round flex flex-column"
          >
            <div
                class="flex flex-row justify-content-center align-items-start gap-2"
            >
              <div class="text-md font-medium text-900">
                {{ item.name }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </DataView>
  <!--  </template>-->








  <DataView :value="steps">
    <!--    <template #header>-->

    <!--      <div class="d-flex justify-content-center align-items-center center">-->
    <!--        <div class="flex">-->
    <!--          <span class="font-medium text-primary text-xl">-->
    <!--            {{ salad.name }}-->
    <!--          </span>-->
    <!--        </div>-->
    <!--      </div>-->

    <!--      <div class="d-flex">-->
    <!--        <div class="flex-row">-->
    <!--          <div class="text-lg font-medium text-900 mt-2">-->
    <!--            Автор: {{ author.name  }}-->
    <!--          </div>-->
    <!--          <div class="text-lg font-medium text-900 mt-2">-->
    <!--            Количество порций: {{ recipe.number_of_servings }}, время приготовления {{ recipe.time_to_cook }} минут-->
    <!--          </div>-->
    <!--          <div class="text-lg font-medium text-900 mt-2">-->
    <!--            {{ salad.description  }}-->
    <!--          </div>-->
    <!--        </div>-->
    <!--      </div>-->

    <!--      <DataView :value="ingredients" :layout="'grid'">-->
    <!--        <template #empty>-->
    <!--          <div class="d-flex justify-content-center align-items-center center">-->
    <!--            <div class="flex">-->
    <!--          <span class="font-medium text-primary text-xl">-->
    <!--            Необходимые ингредиенты не найдены :(-->
    <!--          </span>-->
    <!--            </div>-->
    <!--          </div>-->
    <!--        </template>-->

    <!--        <template #grid="slotProps">-->
    <!--          <div class="grid grid-nogutter">-->
    <!--            <div class="col-12 sm:col-2 md:col-2 xl:col-2 p-2">-->
    <!--              <div-->
    <!--                  class="p-2 border-0 surface-border surface-card border-round flex flex-column"-->
    <!--              >-->
    <!--                <div-->
    <!--                    class="flex flex-row justify-content-center align-items-start gap-2"-->
    <!--                >-->
    <!--                  <div class="text-lg font-medium text-primary">-->
    <!--                    Ингредиенты:-->
    <!--                  </div>-->
    <!--                </div>-->
    <!--              </div>-->
    <!--            </div>-->

    <!--            <div-->
    <!--                v-for="(item, index) in slotProps.items"-->
    <!--                :key="index"-->
    <!--                class="col-12 sm:col-2 md:col-2 xl:col-2 p-2"-->
    <!--            >-->
    <!--              <div-->
    <!--                  class="p-2 border-1 surface-border surface-card border-round flex flex-column"-->
    <!--              >-->
    <!--                <div-->
    <!--                    class="flex flex-row justify-content-center align-items-start gap-2"-->
    <!--                >-->
    <!--                  <div class="text-md font-medium text-900">-->
    <!--                    <template v-if="item==null">-->
    <!--                    </template>-->
    <!--                    <template v-else-if="item.measurement">-->
    <!--                      {{ item.name }} ({{ item.count }} {{ item.measurement.name }})-->
    <!--                    </template>-->
    <!--                  </div>-->
    <!--                </div>-->
    <!--              </div>-->
    <!--            </div>-->
    <!--          </div>-->
    <!--        </template>-->
    <!--      </DataView>-->

    <!--      <DataView :value="saladTypes" :layout="'grid'">-->
    <!--        <template #empty>-->
    <!--          <div class="d-flex justify-content-center align-items-center center">-->
    <!--            <div class="flex">-->
    <!--          <span class="font-medium text-primary text-xl">-->
    <!--            Типы рецепта не указаны :(-->
    <!--          </span>-->
    <!--            </div>-->
    <!--          </div>-->
    <!--        </template>-->

    <!--        <template #grid="slotProps">-->
    <!--          <div class="grid grid-nogutter">-->
    <!--            <div class="col-12 sm:col-2 md:col-2 xl:col-2 p-2">-->
    <!--              <div-->
    <!--                  class="p-2 border-0 surface-border surface-card border-round flex flex-column"-->
    <!--              >-->
    <!--                <div-->
    <!--                    class="flex flex-row justify-content-center align-items-start gap-2"-->
    <!--                >-->
    <!--                  <div class="text-lg font-medium text-primary">-->
    <!--                    Типы салата:-->
    <!--                  </div>-->
    <!--                </div>-->
    <!--              </div>-->
    <!--            </div>-->

    <!--            <div-->
    <!--                v-for="(item, index) in slotProps.items"-->
    <!--                :key="index"-->
    <!--                class="col-12 sm:col-1 md:col-1 xl:col-1 p-2"-->
    <!--            >-->
    <!--              <div-->
    <!--                  class="p-2 border-1 surface-border surface-card border-round flex flex-column"-->
    <!--              >-->
    <!--                <div-->
    <!--                    class="flex flex-row justify-content-center align-items-start gap-2"-->
    <!--                >-->
    <!--                  <div class="text-md font-medium text-900">-->
    <!--                    {{ item.name }}-->
    <!--                  </div>-->
    <!--                </div>-->
    <!--              </div>-->
    <!--            </div>-->
    <!--          </div>-->
    <!--        </template>-->
    <!--      </DataView>-->
    <!--    </template>-->


    <template #empty>
      <div class="d-flex justify-content-center align-items-center center">
        <div class="flex">
          <span class="font-medium text-primary text-xl">
            Шаги рецепта не найдены :(
          </span>
        </div>
      </div>
    </template>

    <template #list="slotProps">
      <div class="grid grid-nogutter">
        <div
            v-for="(item, index) in slotProps.items"
            :key="index"
            class="col-12 relative"
        >
          <div
              class="flex flex-column sm:flex-row sm:align-items-center p-4 gap-3"
              :class="{ 'border-top-1 surface-border': index !== 0 }"
          >
            <div
                class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4"
            >
              <div
                  class="flex flex-row md:flex-column justify-content-between align-items-start gap-2"
              >
                <div>
                  <span class="font-medium text-primary text-xl">
                    Шаг {{ item.step_num }}: {{ item.name }}
                  </span>
                  <div class="text-lg font-medium text-900 mt-2">
                    {{ item.description }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!--    <template #footer>-->
    <!--      <div class="d-flex justify-content-center align-items-center center gap-3">-->
    <!--        <div class="flex flex-column gap-5 justify-content-between align-items-between">-->
    <!--          <span class="font-medium text-primary text-xl">-->
    <!--            Комментарии-->
    <!--          </span>-->
    <!--        </div>-->
    <!--      </div>-->
    <!---->
    <!--      <div class="card">-->
    <!--        <DataView :value="comments" paginator :rows="commentRows">-->
    <!--          <template #empty>-->
    <!--            <div class="d-flex justify-content-center align-items-center center">-->
    <!--              <div class="flex">-->
    <!--          <span class="font-medium text-primary text-xl">-->
    <!--            Комментарии не найдены-->
    <!--          </span>-->
    <!--              </div>-->
    <!--            </div>-->
    <!--          </template>-->

    <!--          <template #header v-if="userComment !== null">-->
    <!--            <template v-if="updatingComment === false">-->
    <!--              <div class="flex flex-column sm:flex-row sm:align-items-center p-2 gap-3" :class="{ 'border-top-1 surface-border': index !== 0 }">-->
    <!--                <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">-->
    <!--                  <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">-->
    <!--                    <div class="surface-100 p-1" style="border-radius: 30px">-->
    <!--                      <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">-->
    <!--                          <span class="text-900 font-medium text-md">-->
    <!--                            {{ userComment.rating }}-->
    <!--                          </span>-->
    <!--                        <i class="pi pi-star-fill text-yellow-500"></i>-->
    <!--                      </div>-->
    <!--                    </div>-->
    <!--                    <div>-->
    <!--                      <div class="text-lg font-medium text-900 mt-2">-->
    <!--                        <template v-if="userComment.author==null">-->
    <!--                        </template>-->
    <!--                        <template v-else-if="userComment.author">-->
    <!--                          {{ userComment.author.name }}-->
    <!--                        </template>-->
    <!--                      </div>-->
    <!--                    </div>-->
    <!--                  </div>-->
    <!--                  <div class="flex flex-column md:align-items-end gap-5">-->
    <!--                      <span class="font-normal text-900">-->
    <!--                        {{ userComment.text }}-->
    <!--                      </span>-->
    <!--                  </div>-->

    <!--                  <div class="flex flex-column md:align-items-end">-->
    <!--                    <div>-->
    <!--                      <Button icon="pi pi-pencil"-->
    <!--                              style="border-top-left-radius: 8px; border-top-right-radius: 8px"-->
    <!--                              severity="success"-->
    <!--                              @click="startUpdatingComment" />-->
    <!--                    </div>-->
    <!--                    <div>-->
    <!--                      <Button icon="pi pi-trash"-->
    <!--                              style="border-bottom-left-radius: 8px; border-bottom-right-radius: 8px"-->
    <!--                              severity="danger"-->
    <!--                              @click="deleteUserComment" />-->
    <!--                    </div>-->
    <!--                  </div>-->

    <!--                </div>-->
    <!--              </div>-->
    <!--            </template>-->
    <!--            <template v-else-if="updatingComment === true">-->
    <!--                <UpdateComment-->
    <!--                    :comment_id="userComment.id"-->
    <!--                    @updatingSuccess="commentUpdateSuccess"-->
    <!--                    @updatingCancelled="commentUpdateCancelled"-->
    <!--                    @updatingError="commentUpdateFailed"-->
    <!--                />-->
    <!--            </template>-->
    <!--          </template>-->
    <!--          <template #header v-if="userComment === null && isAuthValue === true">-->
    <!--            <CreateComment-->
    <!--                :salad_id="this.$route.params.id"-->
    <!--                @creatingSuccess="commentCreateSuccess"-->
    <!--                @creatingError="commentCreateFailed"-->
    <!--            />-->

    <!--          </template>-->

    <!--          <template #list="slotProps">-->
    <!--            <div class="grid grid-nogutter">-->
    <!--              <div v-for="(item, index) in slotProps.items" :key="index" class="col-12">-->
    <!--                <template v-if="item.author_id!==authUserId">-->
    <!--                  <div class="flex flex-column sm:flex-row sm:align-items-center p-4 gap-3" :class="{ 'border-top-1 surface-border': index !== 0 }">-->
    <!--                    <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">-->
    <!--                      <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">-->
    <!--                        <div class="surface-100 p-1" style="border-radius: 30px">-->
    <!--                          <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">-->
    <!--                          <span class="text-900 font-medium text-md">-->
    <!--                            {{ item.rating }}-->
    <!--                          </span>-->
    <!--                            <i class="pi pi-star-fill text-yellow-500"></i>-->
    <!--                          </div>-->
    <!--                        </div>-->
    <!--                        <div>-->
    <!--                          <div class="text-lg font-medium text-900 mt-2">-->
    <!--                            <template v-if="item.author==null">-->
    <!--                            </template>-->
    <!--                            <template v-else-if="item.author">-->
    <!--                              {{ item.author.name }}-->
    <!--                            </template>-->
    <!--                          </div>-->
    <!--                        </div>-->
    <!--                      </div>-->
    <!--                      <div class="flex flex-column md:align-items-end gap-5">-->
    <!--                      <span class="font-normal text-900">-->
    <!--                        {{ item.text }}-->
    <!--                      </span>-->
    <!--                      </div>-->
    <!--                    </div>-->
    <!--                  </div>-->
    <!--                </template>-->

    <!--              </div>-->
    <!--            </div>-->
    <!--          </template>-->
    <!--        </DataView>-->
    <!--      </div>-->
    <!--    </template>-->


  </DataView>
  <!--  </template>-->

</template>

<script>
import Button from 'primevue/button';
import ButtonGroup from "primevue/buttongroup";
import Toast from 'primevue/toast';
import DataView from 'primevue/dataview';
import SaladService from '@/services/salad.service';
import RecipeService from "@/services/recipe.service.js";
import RecipeStepService from "@/services/recipeStep.service.js";
import SaladTypeService from "@/services/saladType.service.js";
import IngredientService from "@/services/ingredient.service.js";
import IngredientTypeService from "@/services/ingredientType.service.js";
import MeasurementService from "@/services/measurement.service.js";
import CommentService from "@/services/comment.service.js";
import UserService from "@/services/user.service.js";
import Utils from "@/services/auth-header.js"
import UpdateComment from "@/components/CommentUpdateChild.vue";
import CreateComment from "@/components/CommentCreateChild.vue";
import { useToast } from "primevue/usetoast";
import {auth} from "@/store/auth.module.js";
import { ref } from 'vue';
// import UpdateSaladChild from "@/components/SaladUpdateChild.vue";

export default {
  name: 'SaladPage',
  inheritAttrs: false,
  components: {
    // UpdateSaladChild,
    UpdateComment,
    CreateComment,
    Button,
    ButtonGroup,
    DataView,
    Toast
  },
  data() {
    return {
      editingMode: false,

      updatingComment: false,
      toast: useToast(),

      index: 0, // just to fix vue warn

      steps: [],
      salad: {},
      recipe: null,
      author: {},
      saladTypes: [],
      ingredients: [],

      userComment: null,
      comments: [],
      commentRows: 30,
      totalCommentPages: 0,
      currentCommentPage: 1,

      isAuthValue: false,
      authUserId: null,
    }
  },
  created() {
    this.isAuth()
    // if (this.isAuthValue) {
    //   this.authUserId = Utils.getUserIdJWT()
    // }

    this.fetchSalad()

    this.fetchFullRecipe() // recipe + steps + ingredients
    this.fetchUserComment()
    this.fetchTypes()
    this.fetchComments()
  },


  methods: {
    fetchFullRecipe() {
      const saladId = this.$route.params.id
      RecipeService.getRecipeDetails(saladId)
          .then(response => {
            this.recipe = response.data.data.recipe
            this.fetchSteps(this.recipe.id)
            this.fetchIngredients(this.recipe.id)
          })
          .catch(error => {
            this.editingMode = true
          })
    },

    fetchSalad() {
      const saladId = this.$route.params.id
      SaladService.getSaladDetails(saladId)
          .then(response => {
            this.salad = response.data.data.salad
            this.fetchAuthor(this.salad.author_id)

            console.log(this.salad) // FIXME
          })
    },

    fetchAuthor(authorId) {
      UserService.getUserDetails(authorId)
          .then(response => {
            this.author = response.data.data.user
          })
    },

    fetchSteps(recipeId) {
      RecipeStepService.getRecipeSteps(recipeId)
          .then(response => {
            this.steps = [...response.data.data.steps]

            console.log("STEPS", this.steps)
          })
    },

    fetchTypes() {
      const saladId = this.$route.params.id
      SaladTypeService.getSaladTypes(saladId)
          .then(response => {
            this.saladTypes = [...response.data.data.types]
          }).catch(error => {
        console.log(error)
      })
    },

    fetchIngredients(recipeId) {
      IngredientService.getRecipeIngredients(recipeId)
          .then(response => {
            this.ingredients = [...response.data.data.ingredients]
            Promise.all(
                this.ingredients
                    .filter(ingredient => ingredient !== null)
                    .map(ingredient => {
                      IngredientTypeService.getType(ingredient.type_id).then(response => {
                        ingredient.type = response.data.data.ingredientType
                      })
                      MeasurementService.getMeasurementByRecipe(ingredient.id, recipeId).then(response => {
                        ingredient.measurement = response.data.data.measurement
                        ingredient.count = response.data.data.count
                      })
                    })
            ).catch(error => {
              console.log(error)
            })
          })
    },

    fetchComments() {
      const saladId = this.$route.params.id
      CommentService.getAllBySaladId(saladId, this.currentCommentPage)
          .then(response => {
            this.comments = [...response.data.data.comments]
            this.totalCommentPages = response.data.data.num_pages

            Promise.all(
                this.comments
                    .filter(comment => comment !== null)
                    .map(comment => {
                      UserService.getUserDetails(comment.author_id)
                          .then(response => {
                            comment.author = response.data.data.user
                          })
                    })
            ).catch(error => {
              console.log(error)
            })

          })
    },

    isAuth() {
      this.isAuthValue = Utils.isAuth()

    },

    fetchUserComment() {
      if (this.isAuthValue) {
        this.authUserId = Utils.getUserIdJWT()

        const saladId = this.$route.params.id
        CommentService.getUserComment(saladId, this.authUserId)
            .then(response => {
              if (response.statusText === "OK") {
                this.userComment = response.data.data.comment

                UserService.getUserDetails(this.userComment.author_id)
                    .then(response => {
                      this.userComment.author = response.data.data.user
                    })
              }
            }).catch(error => {
          console.log("user comment not found")
        })
      } else {
        console.log("it's an unauthorized guest")
      }
    },

    deleteUserComment() {
      CommentService.deleteComment(this.userComment.id)
          .then(response => {
            if (response.status === 200) {
              this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Комментарий успешно удален', life: 3000 });
              this.userComment = null
            }
          })
    },

    startUpdatingComment() {
      this.updatingComment = true
    },

    commentUpdateSuccess() {
      this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Комментарий успешно обновлен', life: 3000 });
      this.updatingComment = false
      this.fetchUserComment()
    },

    commentUpdateCancelled() {
      this.updatingComment = false
    },

    commentUpdateFailed(errorText) {
      this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при обновлении данных: ${errorText}`, life: 3000 });
      // this.updatingComment = false
    },

    commentCreateSuccess() {
      this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Комментарий успешно создан', life: 3000 });
      this.fetchUserComment()
      this.fetchComments()
    },

    commentCreateFailed(errorText) {
      this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при создании комментария: ${errorText}`, life: 3000 });
    },

  }
}
</script>
