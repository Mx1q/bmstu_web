<template>
  <DataView :value="salads">

<!--    <template #header>-->
<!--      <InputGroup>-->
<!--        <MultiSelect v-model="selectedIngredients" :options="ingredients" :maxSelectedLabels="5" optionLabel="name" optionValue="id"-->
<!--                     @selectall-change="onSelectAllIngredientsChange($event)" @change="onSelectIngredientsChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Ингредиенты" class="w-full md:w-20rem" />-->
<!--        <MultiSelect v-model="selectedSaladTypes" :options="saladTypes" :maxSelectedLabels="5" optionLabel="name" optionValue="id"-->
<!--                     @selectall-change="onSelectAllSaladTypesChange($event)" @change="onSelectSaladTypeChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Типы Салатов" class="w-full md:w-20rem" />-->

<!--        <Button label="Поиск" class="p-button p-component"  @click="fetchSalads" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />-->
<!--      </InputGroup>-->
<!--    </template>-->

    <template #empty>
      <div class="d-flex justify-content-center align-items-center center">
        <div class="flex">
        <span class="font-medium text-primary text-xl">
          По вашему запросу ничего не найдено :(
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
                  <div class="text-lg font-medium text-900 mt-2">
                    <template v-if="item==null">
                      It's an error, please contact admins
                    </template>
                    <template v-else-if="item.name">
                      <div class="flex flex-row gap-3">
                        {{ item.name }} ({{ statuses[item.status] }})
                        <Rating v-model="item.rating" readonly :cancel="false" />
                        ({{ item.rating }} / 5)
                      </div>
                    </template>
                  </div>
                </div>
                <div v-if="item !== null"
                    class="surface-100 p-1"
                    style="border-radius: 30px"
                >
                  {{ item.description }}
                </div>
              </div>
              <div class="flex flex-column md:align-items-end gap-5">
                <div v-if="item===null"></div>
                <div v-else-if="item.id" class="flex flex-row-reverse md:flex-row gap-2">
                  <RouterLink :to="`/legacy/salads/${item.id}`">
                    <Button
                        icon="pi pi-info-circle"
                        label="Подробнее"
                        class="flex-auto md:flex-initial white-space-nowrap"
                    >
                    </Button>
                  </RouterLink>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </DataView>
</template>

<script>
import SaladService from '@/services/salad.service';
import RecipeService from "@/services/recipe.service.js";

import IngredientService from "@/services/ingredient.service.js";
import SaladTypeService from "@/services/saladType.service.js";

import Button from 'primevue/button';
import DataView from 'primevue/dataview';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';

import MultiSelect from 'primevue/multiselect';
import VirtualScroller from 'primevue/virtualscroller';
import { ref } from 'vue';

export default {
  name: 'SaladsPage',
  components: {
    Button,
    DataView,
    InputGroup,
    InputGroupAddon,
    VirtualScroller,
    MultiSelect
  },
  data() {
    return {
      rows: 30,
      totalPages: 0,
      salads: [],
      currentPage: 1,

      ingredients: [],
      currentIngredientPage: 1,
      totalIngredientPages: 0,
      selectedIngredients: [],
      selectAllIngredients: false,

      saladTypes: [],
      currentSaladTypePage: 1,
      totalSaladTypePages: 0,
      selectedSaladTypes: [],
      selectAllSaladTypes: false,

      statuses: {
        1: 'редактирование',
        2: 'отправлен на модерацию',
        3: 'отклонен',
        4: 'опубликован',
        5: 'снят с публикации',
      }

    }
  },
  created() {
    this.fetchSalads()
  },
  methods: {
    fetchSalads() {
      SaladService.getUserSalads(this.currentPage, this.selectedIngredients, this.selectedSaladTypes)
          .then(response => {
            this.salads = [...response.data.data.salads]
            this.totalPages = response.data.data.num_pages
            Promise.all(
                this.salads
                    // .filter(salad => salad !== null)
                    .map(salad => {
                      RecipeService.getRecipeDetails(salad.id)
                          .then(response => {
                            salad.rating = response.data.data.recipe.rating
                            if (!salad.rating) {
                              salad.rating = 0
                            }
                            salad.status = response.data.data.recipe.status
                          })
                          .catch(error => {
                            salad.status = 1
                            salad.rating = 0.0
                          })
                    })
            ).then(() => {})
          })
          .catch(error => {
            console.error(error)
          })
    },

    onPageChange(event) {
      this.currentPage = event.page + 1;
      this.fetchSalads();
    },
  }
}
</script>
