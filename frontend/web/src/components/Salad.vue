<template>
  <Toast />
  <ConfirmDialog />

<!--  <Button label="Show" @click="visible = true" />-->

  <Dialog v-model:visible="addStepVisible" modal @after-hide="enterPreviewMode" header="Типы рецепта" @input="enterPreviewMode" :style="{ width: '50rem' }">
    <CreateRecipeStepChild style="margin-bottom: 5px; margin-top: 5px"
                           :recipe_id="recipe.id"
    />
  </Dialog>

  <Dialog v-model:visible="addIngredientVisible" modal @after-hide="enterPreviewMode" header="Добавление ингредиента" @input="enterPreviewMode" :style="{ width: '50rem' }">
    <InputGroup style="margin-bottom: 5px; margin-top: 5px">
      <Dropdown v-model="selectedIngredientId" :options="availableIngredients" optionLabel="name" optionValue="id"
                :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Ингредиент" class="w-full md:w-14rem" />
      <InputNumber v-model="selectedIngredientAmount" inputId="minmax" :min="1" />
      <Dropdown v-model="selectedMeasurement" :options="availableMeasurements" optionLabel="name" placeholder="Единица измерения" class="w-full md:w-14rem" />
      <Button label="Добавить" class="p-button p-component"  @click="addIngredient" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />
    </InputGroup>
  </Dialog>

  <Dialog v-model:visible="editTypesVisible" modal @after-hide="enterPreviewMode" header="Типы рецепта" @input="enterPreviewMode" :style="{ width: '50rem' }">
      <InputGroup style="margin-bottom: 5px; margin-top: 5px">
        <MultiSelect v-model="selectedSaladTypeIds" :options="availableSaladTypes" :maxSelectedLabels="5" optionLabel="name" optionValue="id"
                     @selectall-change="onSelectAllSaladTypesChange($event)" @change="onSelectSaladTypeChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Типы Салатов" class="w-full md:w-20rem" />
        <Button label="Обновить" class="p-button p-component"  @click="updateSaladTypes" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />
      </InputGroup>
  </Dialog>

  <Dialog v-model:visible="visible" modal header="Редактирование рецепта" @input="enterPreviewMode" :style="{ width: '50rem' }">
    <div style="margin-bottom: 5px; margin-top: 5px">
      <UpdateSaladChild
          :salad_id="salad.id"
          @updatingSuccess="saladUpdateSuccess"
      />
    </div>

    <div>
      <template v-if="recipe">
        <UpdateRecipeChild
            :salad_id="salad.id"
            :recipe_id="recipe.id"
            @updatingSuccess="recipeUpdateSuccess"
        />
      </template>
      <template v-else>
        <CreateRecipeChild
            :salad_id="salad.id"
            @creatingSuccess="recipeCreateSuccess"
        />
      </template>

    </div>

<!--    <div>-->
<!--      <InputGroup style="margin-bottom: 5px; margin-top: 5px">-->
<!--        <MultiSelect v-model="selectedSaladTypeIds" :options="availableSaladTypes" :maxSelectedLabels="5" optionLabel="name" optionValue="id"-->
<!--                     @selectall-change="onSelectAllSaladTypesChange($event)" @change="onSelectSaladTypeChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Типы Салатов" class="w-full md:w-20rem" />-->
<!--        <Button label="Обновить" class="p-button p-component"  @click="updateSaladTypes" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />-->
<!--      </InputGroup>-->
<!--    </div>-->
  </Dialog>

<!--  <template v-if="salad !== null && salad.author_id === authUserId && recipe && recipe.status === 4">-->
<!--    <Button label="Снять с публикации" @click="storeSalad" />-->
<!--  </template>-->

  <template v-if="salad !== null && salad.author_id === authUserId">
    <template v-if="editingMode === true">
<!--      <Button label="Предпросмотр" @click="enterPreviewMode" />-->

<!--      <div style="margin-bottom: 5px; margin-top: 5px">-->
<!--        <UpdateSaladChild-->
<!--            :salad_id="salad.id"-->
<!--            @updatingSuccess="saladUpdateSuccess"-->
<!--        />-->
<!--      </div>-->

<!--      <div style="margin-bottom: 5px; margin-top: 5px">-->
<!--        <template v-if="recipe === null">-->
<!--          <CreateRecipeChild-->
<!--              :salad_id="salad.id"-->
<!--              @creatingSuccess="recipeCreateSuccess"-->
<!--          />-->
<!--        </template>-->
<!--        <template v-else>-->
<!--          <UpdateRecipeChild-->
<!--              :salad_id="salad.id"-->
<!--              :recipe_id="recipe.id"-->
<!--              @updatingSuccess="recipeUpdateSuccess"-->
<!--          />-->
<!--        </template>-->
<!--      </div>-->


<!--      <InputGroup style="margin-bottom: 5px; margin-top: 5px">-->
<!--        <MultiSelect v-model="selectedSaladTypeIds" :options="availableSaladTypes" :maxSelectedLabels="5" optionLabel="name" optionValue="id"-->
<!--                     @selectall-change="onSelectAllSaladTypesChange($event)" @change="onSelectSaladTypeChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Типы Салатов" class="w-full md:w-20rem" />-->
<!--        <Button label="Обновить" class="p-button p-component"  @click="updateSaladTypes" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />-->
<!--      </InputGroup>-->

<!--      <template v-if="recipe !== null">-->
<!--        <InputGroup style="margin-bottom: 5px; margin-top: 5px">-->
<!--          <Dropdown v-model="selectedIngredientId" :options="availableIngredients" optionLabel="name" optionValue="id"-->
<!--                    :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Ингредиент" class="w-full md:w-14rem" />-->
<!--          <InputNumber v-model="selectedIngredientAmount" inputId="minmax" :min="1" />-->
<!--          <Dropdown v-model="selectedMeasurement" :options="availableMeasurements" optionLabel="name" placeholder="Единица измерения" class="w-full md:w-14rem" />-->
<!--          <Button label="Добавить" class="p-button p-component"  @click="addIngredient" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />-->
<!--        </InputGroup>-->

<!--        <CreateRecipeStepChild style="margin-bottom: 5px; margin-top: 5px"-->
<!--                               :recipe_id="recipe.id"-->
<!--        />-->
<!--      </template>-->

    </template>
<!--    <template v-else-if="recipe === null || recipe.status !== 2">-->
<!--      <Button label="Редактировать" @click="enterEditMode" />-->
<!--    </template>-->

<!--    <template v-if="recipe && recipe.status === 1 && !editingMode">-->
<!--      <Button label="Отправить на модерацию" @click="sendToModeration" />-->
<!--      <Button v-if="!isAdmin()" label="Удалить" @click="deleteSalad"/>-->
<!--    </template>-->

  </template>

<!--  <template v-if="recipe && recipe.status === 2 && isAdmin()">-->
<!--    <Button label="Опубликовать" @click="approveSalad" />-->
<!--    <Button label="Отклонить" @click="cancelSalad" />-->
<!--  </template>-->

<!--  <template v-if="recipe && isAdmin()">-->
<!--    <Button label="Удалить" @click="deleteSalad"/>-->
<!--  </template>-->

<!--  <template v-if="editingMode === false">-->
    <!--  SALAD AND RECIPE INFO SECTION-->
<!--  <div class="d-flex justify-content-center align-items-center center">-->
<!--    <div class="flex flex-row">-->
<!--    <template v-if="salad !== null">-->
      <div class="d-flex justify-content-center align-items-center center m-1">
        <div class="flex flex-row">

          <template v-if="salad !== null && salad.author_id === authUserId || isAdmin()">
            <Button icon="pi pi-trash"
                    style="border-radius: 8px"
                    severity="danger"
                    class="mr-1"
                    @click="deleteSalad"/>
          </template>

          <template v-if="salad.author_id === authUserId && (recipe === null || recipe.status !== 2)">
            <Button icon="pi pi-pencil"
                    style="border-radius: 8px"
                    severity="success"
                    class="mr-1"
                    @click="enterEditMode"
            />
<!--            @click="visible = true"-->
            <!--                    @click="enterEditMode"-->
          </template>

          <template v-if="recipe && recipe.status === 2 && isAdmin()">
            <Button label="Отклонить"
                    @click="cancelSalad"
                    severity="danger"
                    style="border-radius: 8px"
                    class="m-1"
            />
          </template>

          <span class="font-medium text-primary text-xl align-self-center">
            {{ salad.name }}
          </span>

          <template v-if="recipe && recipe.status === 2 && isAdmin()">
            <Button label="Опубликовать"
                    style="border-radius: 8px"
                    severity="success"
                    @click="approveSalad"
                    class="m-1"
            />
          </template>

          <template v-if="salad !== null && salad.author_id === authUserId">
            <template v-if="recipe && recipe.status === 1">
              <Button
                  severity="success"
                  v-if="recipe && recipe.status === 1"
                  class="ml-1"
                  style="border-radius: 8px"
                  label="Отправить на модерацию"
                  @click="sendToModeration" />
            </template>
          </template>

          <!--        <template v-if="salad !== null && salad.author_id === authUserId && recipe && recipe.status === 4">-->
          <!--          <Button label="Снять с публикации" @click="storeSalad" />-->
          <!--          <Button icon="pi pi-bookmark" severity="secondary" rounded />-->
          <!--        </template>-->
        </div>
      </div>

      <div class="d-flex">
        <div class="flex-row">
          <div class="text-lg font-medium text-900 mt-2">
            Автор: {{ author.name  }}
          </div>
          <div class="text-lg font-medium text-900 mt-2" v-if="recipe !== null">
            Количество порций: {{ recipe.number_of_servings }}, время приготовления {{ recipe.time_to_cook }} минут
          </div>
          <div class="text-lg font-medium text-900 mt-2">
            {{ salad.description  }}
          </div>
        </div>
      </div>
<!--    </template>-->

    <!--  INGREDIENTS LIST SECTION-->
    <DataView :value="ingredients" :layout="'grid'">
      <template #empty>
        <div class="d-flex justify-content-center align-items-center center">
          <div class="flex">
            <Button icon="pi pi-plus"
                    v-if="recipe && recipe.status === 1"
                    style="border-radius: 8px"
                    severity="success"
                    class="mr-1"
                    @click="addIngredientVisible = true"
            />
            <span class="font-medium text-primary text-xl align-self-center">
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
                <Button icon="pi pi-plus"
                        v-if="recipe && recipe.status === 1"
                        style="border-radius: 8px"
                        severity="success"
                        class="mr-1"
                        @click="addIngredientVisible = true"
                />
                <div class="text-lg font-medium text-primary align-self-center">
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
                  <template v-else-if="item.measurement && recipe.status === 1">
                    <div>
                      {{ item.name }} ({{ item.count }} {{ item.measurement.name }})
                    </div>
                    <div class="flex flex-row align-items-center justify-content-center">
                      <div>
                        <Button icon="pi pi-trash"
                                v-if="recipe && recipe.status === 1"
                                style="border-radius: 8px"
                                severity="danger"
                                @click="unlinkIngredient(item.id)" />
                      </div>
                    </div>

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

    <!--  SALAD TYPES SECTION-->
    <DataView :value="saladTypes" :layout="'grid'">
      <template #empty>
        <div class="d-flex justify-content-center align-items-center center">
          <div class="flex">
            <Button icon="pi pi-pencil"
                    v-if="recipe && recipe.status === 1"
                    style="border-radius: 8px"
                    severity="success"
                    class="mr-1"
                    @click="editTypesVisible = true"
            />
            <span class="font-medium text-primary text-xl align-self-center">
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
                <Button icon="pi pi-pencil"
                        v-if="recipe && recipe.status === 1"
                        style="border-radius: 8px"
                        severity="success"
                        class="mr-1"
                        @click="editTypesVisible = true"
                />
                <div class="text-lg font-medium text-primary align-self-center">
                  Типы салата:
                </div>
              </div>
            </div>
          </div>

          <div
              v-for="(item, index) in slotProps.items"
              :key="index"
              class="col-12 sm:col-1 md:col-1 xl:col-1 p-2 align-self-center"
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

    <!--  RECIPE STEPS SECTION-->
    <DataView :value="steps" class="m-1">
      <template #empty>
        <div class="d-flex justify-content-center align-items-center center">
          <div class="flex">

            <Button icon="pi pi-plus"
                    v-if="recipe && recipe.status === 1"
                    style="border-radius: 8px"
                    severity="success"
                    class="mr-1"
                    @click="addStepVisible = true"
            />

          <span class="font-medium text-primary text-xl align-self-center">
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
                class="flex flex-column sm:flex-row sm:align-items-between p-4 gap-3"
                :class="{ 'border-top-1 surface-border': index !== 0 }"
            >
              <div
                  class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4"
              >
                <div class="flex flex-row justify-content-between align-items-center gap-2">
                  <div class="flex flex-row align-items-between" v-if="!item.updating">
                    <template v-if="recipe.status === 1 && !item.updating">
                      <div class="flex flex-column md:align-items-end" style="margin-right: 10px">
                        <div>
                          <Button icon="pi pi-pencil"
                                  style="border-top-left-radius: 8px; border-top-right-radius: 8px"
                                  severity="success"
                                  @click="updateRecipeStep(item.id, index)" />
                        </div>
                        <div>
                          <Button icon="pi pi-trash"
                                  style="border-bottom-left-radius: 8px; border-bottom-right-radius: 8px"
                                  severity="danger"
                                  @click="deleteRecipeStep(item.id, index)" />
                        </div>
                      </div>
                    </template>

                    <div class="flex flex-column">
                    <span class="font-medium text-primary text-xl">
                      Шаг {{ item.step_num }}: {{ item.name }}
                    </span>
                      <div class="text-lg font-medium text-900 mt-2">
                        {{ item.description }}
                      </div>
                    </div>
                  </div>
                  <template v-else-if="recipe.status === 1 && item.updating">
                    <UpdateRecipeStepChild
                        :step_id="item.id" :step_name_prop="item.name" :step_description_prop="item.description"
                        @updatingCancelled="updatingStepCancelled(index)"
                        @updatingSuccess="updatingStepSuccess"
                    />
                  </template>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>
    </DataView>








    <!--  COMMENTS SECTION-->
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
<!--  </template>-->

</template>

<script lang="ts">
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
import {nextTick, ref} from 'vue';
import UpdateSaladChild from "@/components/SaladUpdateChild.vue";
import CreateRecipeChild from "@/components/RecipeCreateChild.vue";
import UpdateRecipeChild from "@/components/RecipeUpdateChild.vue";

import ConfirmDialog from 'primevue/confirmdialog';
import { useConfirm } from "primevue/useconfirm";
import RecipeModel from "@/models/RecipeModel.js";
import CreateRecipeStepChild from "@/components/RecipeStepCreateChild.vue";
import UpdateRecipeStepChild from "@/components/RecipeStepUpdateChild.vue";

export default {
  name: 'SaladPage',
  inheritAttrs: false,
  components: {
    UpdateRecipeStepChild,
    CreateRecipeStepChild,
    UpdateRecipeChild,
    CreateRecipeChild,
    UpdateSaladChild,
    UpdateComment,
    CreateComment,
    Button,
    ButtonGroup,
    DataView,
    Toast,
    ConfirmDialog,
  },
  data() {
    return {
      visible: ref(false),
      addIngredientVisible: ref(false),
      editTypesVisible: ref(false),
      addStepVisible: ref(false),

      editingMode: false,

      updatingComment: false,
      toast: useToast(),
      confirm: useConfirm(),

      index: 0, // just to fix vue warn

      steps: [],
      salad: null,
      recipe: null,
      author: {},

      availableSaladTypes: [], // all salad types
      selectedSaladTypeIds: [], // selected
      saladTypes: [], // printing this
      currentSaladTypePage: 1,
      totalSaladTypePages: 0,
      selectAllSaladTypes: false,



      ingredients: [], // printing this
      availableIngredients: [], // all ingredients
      selectedIngredientId: ref(), // selected

      selectedIngredientAmount: ref(1),
      availableMeasurements: [], // all measurements
      selectedMeasurement: ref(), // selected

      currentIngredientPage: 1,
      totalIngredientPages: 0,
      selectAllIngredients: false,




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
    if (this.isAuthValue) {
      this.authUserId = Utils.getUserIdJWT()
    }

    this.fetchSalad()

    this.fetchAllSaladTypes()
    this.fetchAllIngredients()
    this.fetchAllMeasurements()

    this.fetchFullRecipe() // recipe + steps + ingredients
    this.fetchUserComment()
    this.fetchTypes()
    this.fetchComments()
  },


  methods: {
    isAdmin() {
      if (this.isAuthValue) {
        return Utils.getUserRoleJWT() === 'admin';
      }
      return false
    },

    fetchAllMeasurements() {
      MeasurementService.getAll()
          .then(response => {
            this.availableMeasurements = [...response.data.data.measurements]
            this.selectedMeasurement = this.availableMeasurements[0]
          })
    },

    fetchAllIngredients() {
      IngredientService.getAllByPage(this.currentIngredientPage)
          .then(response => {
            this.availableIngredients = [...response.data.data.ingredients]
            this.totalIngredientPages = response.data.data.num_pages
            this.selectedIngredientId = this.availableIngredients[0].id
          })
    },
    addIngredient() {
      console.log('adding ingredient', this.selectedIngredientId, this.selectedIngredientAmount, this.selectedMeasurement.id)

      IngredientService.link(
          this.recipe.id,
          this.$route.params.id,
          this.selectedIngredientId,
          this.selectedMeasurement.id,
          this.selectedIngredientAmount,
      ).then(() =>
          this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Ингредиент успешно добавлен', life: 3000 })
      )
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при добавлении ингредиента: ${error.response.data.error}`, life: 3000 });
          })
    },

    fetchAllSaladTypes() {
      SaladTypeService.getAllByPage(this.currentSaladTypePage)
          .then(response => {
            this.availableSaladTypes = response.data.data.salad_types
            this.totalSaladTypePages = response.data.data.num_pages

            console.log(this.availableSaladTypes)
          })
    },
    onSelectAllSaladTypesChange(event) {
      this.selectedSaladTypeIds.value = event.checked ? this.availableSaladTypes.map((item) => item.id) : [];
      this.selectAllSaladTypes = event.checked;
    },
    onSelectSaladTypeChange(event) {
      this.selectAllSaladTypes = event.value.length === this.availableSaladTypes.length;
    },
    updateSaladTypes() {
      const isLEqual = (id, saladType) => id === saladType.id;
      const isREqual = (saladType, id) => id === saladType.id;
      const onlyInLeft = (left, right, compareFunction) =>
          left.filter(leftValue =>
              !right.some(rightValue =>
                  compareFunction(leftValue, rightValue)));

      const toUnlink = onlyInLeft(this.saladTypes, this.selectedSaladTypeIds, isREqual)
      const toLink = onlyInLeft(this.selectedSaladTypeIds, this.saladTypes, isLEqual)

      Promise.all(
          toLink.map(id => {
            SaladTypeService.link(this.$route.params.id, id)}),
          toUnlink.map(saladType => {
            SaladTypeService.unlink(this.$route.params.id, saladType.id)})
      )
          .then(() => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Типы салата успешно обновлены', life: 3000 })

            // console.log(this.saladTypes)
            // this.fetchTypes()
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при обновлении типов салата: ${error.response.data.error}`, life: 3000 });
          })

      // this.fetchTypes();
      // this.$nextTick(() => {
      //   this.fetchTypes()
      // })
    },

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
          })
    },

    fetchTypes() {
      const saladId = this.$route.params.id
      SaladTypeService.getSaladTypes(saladId)
          .then(response => {
            this.saladTypes = [...response.data.data.types]
            this.selectedSaladTypeIds = this.saladTypes.map(value => value.id)
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
      if (errorText.includes("rate out of range")) {
        errorText = "пожалуйста, оцените рецепт."
      }

      this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при создании комментария: ${errorText}`, life: 3000 });
    },

    recipeCreateSuccess(recipe) {
      this.recipe = recipe
      console.log(this.recipe) // FIXME
      // this.fetchRecipe()
    },

    recipeUpdateSuccess(recipe) {
      this.recipe = recipe
    },

    saladUpdateSuccess(salad) {
      this.salad = salad
    },

    enterEditMode() {
      if (this.recipe !== null && this.recipe.status !== 1 && this.recipe.status !== 3 && this.recipe.status !== 5) {
        this.confirm.require({
          message: 'В случае редактирования рецепта, Вам снова придется отправить его на модерацию',
          header: 'Подтверждение',
          icon: 'pi pi-exclamation-triangle',
          rejectClass: 'p-button-secondary p-button-outlined',
          rejectLabel: 'Отмена',
          acceptLabel: 'Принять',
          accept: () => {
            const recipe = new RecipeModel(0, 0, this.salad.id, 1)
            RecipeService.updateRecipe(this.recipe.id, recipe)
                .then(response => {
                  this.editingMode = true
                  this.toast.add({ severity: 'info', summary: 'Подтверждено', detail: 'Вы перешли в режим редактирования. После завершения необходимых действий не забудьте отправить рецепт на модерацию', life: 3000 });
                  this.recipe = response.data.data.recipe

                  this.visible = true
                })
                .catch(error => {
                  this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при обновлении рецепта: ${error.response.data.error}`, life: 3000 });
                })

          },
          reject: () => {
            this.toast.add({ severity: 'info', summary: 'Отклонено', detail: 'Вы не перешли в режим редактирования', life: 3000 });
          }
        });
      } else {
        // this.editingMode = true

        if (this.recipe && this.recipe.status !== 1) {
          const recipe = new RecipeModel(0, 0, this.salad.id, 1)
          RecipeService.updateRecipe(this.recipe.id, recipe)
              .then(response => {
                this.editingMode = true
                // this.toast.add({ severity: 'info', summary: 'Подтверждено', detail: 'Вы перешли в режим редактирования. После завершения необходимых действий не забудьте отправить рецепт на модерацию', life: 3000 });
                this.recipe = response.data.data.recipe

                this.visible = true
              })
              .catch(error => {
                console.log(error)
              })
        } else {
          this.editingMode = true
          this.visible = true
        }
      }
    },

    enterPreviewMode() {
      this.editingMode = false
      this.fetchTypes()

      if (this.recipe !== null) {
        this.fetchIngredients(this.recipe.id)
        this.fetchSteps(this.recipe.id)
      }
    },

    unlinkIngredient(ingredientId) {
      IngredientService.unlink(this.recipe.id, this.$route.params.id, ingredientId)
          .then(() => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Ингредиент успешно удален', life: 3000 });
            let i = 0
            this.ingredients.map(ingredient => {
              if (ingredient.id === ingredientId) {
                this.ingredients.splice(i, 1)
              }
              i++
            })
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при удалении ингредиента: ${error.response.data.error}`, life: 3000 });
          })
    },

    updateRecipeStep(id, index) {
      this.steps[index].updating = true
    },
    updatingStepCancelled(index) {
      this.steps[index].updating = false
    },
    updatingStepSuccess(step) {
      let index = step.step_num - 1
      this.steps[index].updating = false
      this.steps[index].name = step.name
      this.steps[index].description = step.description
    },


    deleteRecipeStep(id, index) {
      console.log('delete', id)

      RecipeStepService.deleteById(id)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Шаг успешно удален', life: 3000 });
            // this.steps.splice(index, 1)
            this.fetchSteps(this.recipe.id)
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при удалении шага: ${error.response.data.error}`, life: 3000 });
          })
    },

    sendToModeration() {
      const recipe = new RecipeModel(0, 0, this.$route.params.id, 2)
      RecipeService.updateRecipe(this.recipe.id, recipe)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт успешно отправлен на модерацию', life: 3000 });
            this.$router.push('/salads');
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при попытке отправить рецепт на модерацию: ${error.response.data.error}`, life: 3000 });
          })
    },

    approveSalad() {
      console.log("approving salad")
      const recipe = new RecipeModel(0, 0, this.$route.params.id, 4)
      RecipeService.updateRecipe(this.recipe.id, recipe)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт успешно опубликован', life: 3000 });
            this.$router.push('/salads');
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при попытке опубликовать: ${error.response.data.error}`, life: 3000 });
          })
    },

    cancelSalad() {
      console.log("cancelling salad")
      const recipe = new RecipeModel(0, 0, this.$route.params.id, 3)
      RecipeService.updateRecipe(this.recipe.id, recipe)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт отклонен', life: 3000 });
            this.$router.push('/salads');
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при попытке отклонить рецепт: ${error.response.data.error}`, life: 3000 });
          })
    },

    deleteSalad() {
      SaladService.deleteSalad(this.$route.params.id)
          .then(response => {
            this.$router.push(`/salads`)
                .then(() => { this.$router.go(0) })
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт удален', life: 3000 });
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при попытке удалить рецепт: ${error.response.data.error}`, life: 3000 });
          })
    },

    storeSalad() {
      // console.log("should store salad")
      const recipe = new RecipeModel(0, 0, this.$route.params.id, 5)
      RecipeService.updateRecipe(this.recipe.id, recipe)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт успешно снят с публикации', life: 3000 });
            this.$router.push('/mySalads');
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при попытке снять рецепт с публикации: ${error.response.data.error}`, life: 3000 });
          })
    }

  }
}
</script>
<script setup lang="ts">
</script>