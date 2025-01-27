<template>
    <DataView :value="salads" paginator :rows="rows" :totalRecords="totalPages*rows" @page="onPageChange">

      <template #header>
        <InputGroup>
            <MultiSelect v-model="selectedIngredients" :options="ingredients" :maxSelectedLabels="3" optionLabel="name" optionValue="id"
                         @selectall-change="onSelectAllIngredientsChange($event)" @change="onSelectIngredientsChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Ингредиенты" class="w-full md:w-10rem" />
            <MultiSelect v-model="selectedSaladTypes" :options="saladTypes" :maxSelectedLabels="3" optionLabel="name" optionValue="id"
                         @selectall-change="onSelectAllSaladTypesChange($event)" @change="onSelectSaladTypeChange($event)" :virtualScrollerOptions="{ itemSize: 30 }" placeholder="Типы Салатов" class="w-full md:w-10rem" />

<!--            FIXME-->
<!--            <FloatLabel>-->
              <InputNumber id="minRateInput" v-model="selectedMinRate" :minFractionDigits="0" :maxFractionDigits="2" :min="0" :max="5" prefix="Мин. рейтинг: " suffix="&#9733" />
<!--              <label for="minRateInput" class="text outline-no >Минимальный рейтинг</label>-->
<!--            </FloatLabel>-->

<!--            FIXME: check user in localstorage-->
            <Dropdown v-if="isAdmin() === true" v-model="selectedStatus" :options="statuses" optionLabel="name" placeholder="Статус" class="w-full md:w-14rem" />

<!--            <Button label="testing" @click="testing" />-->

            <Button label="Поиск" class="p-button p-component"  @click="fetchSalads" style="border-bottom-right-radius: 8px; border-top-right-radius: 8px" />
        </InputGroup>
      </template>





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
                      <template v-if="item===null">
                        It's an error, please contact admins
                      </template>
                      <template v-else-if="item.name">
                        <div class="flex flex-row gap-3">
                          {{ item.name }}
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
                  <div v-if="item==null"></div>
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
import Utils from '@/services/auth-header.js'

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

      selectedMinRate: ref(1.00),

      isAuthValue: false,

      selectedStatus: {},
      statuses: [
        {id: 1, name: 'редактирование'},
        {id: 2, name: 'отправлен на модерацию'},
        {id: 3, name: 'отклонен'},
        {id: 4, name: 'опубликован'},
        {id: 5, name: 'снят с публикации'},
      ]
    }
  },
  created() {
    this.isAuth()
    this.fetchSalads()
    this.fetchIngredients()
    this.fetchSaladTypes()

    this.selectedStatus = this.statuses[3]
  },
  methods: {
    isAuth() {
      this.isAuthValue = Utils.isAuth()
    },

    isAdmin() {
      if (this.isAuthValue) {
        return Utils.getUserRoleJWT() === 'admin';
      }
      return false
    },

    fetchSalads() {
      console.log(this.currentPage)

      let status = 4
      if (this.selectedStatus.id) {
        status = this.selectedStatus.id
      }

      if (this.isAdmin()) {
        SaladService.getSaladsByStatus(this.currentPage, this.selectedIngredients, this.selectedSaladTypes, this.selectedMinRate, status)
            .then(response => {
              this.salads = [...Array((this.currentPage - 1) * 30).fill(null), ...response.data.data.salads]
              this.totalPages = response.data.data.num_pages
              Promise.all(
                  this.salads
                      .filter(salad => salad !== null)
                      .map(salad =>
                          RecipeService.getRecipeRating(salad.id).then(response => {
                            salad.rating = response.data.data.rating
                          })
                      )
              ).then(() => {})
            })
            .catch(error => {
              console.error(error)
            })
      } else {
        SaladService.getSalads(this.currentPage, this.selectedIngredients, this.selectedSaladTypes, this.selectedMinRate)
            .then(response => {
              this.salads = [...Array((this.currentPage - 1) * 30).fill(null), ...response.data.data.salads]
              this.totalPages = response.data.data.num_pages
              Promise.all(
                  this.salads
                      .filter(salad => salad !== null)
                      .map(salad =>
                          RecipeService.getRecipeRating(salad.id).then(response => {
                            salad.rating = response.data.data.rating
                          })
                      )
              ).then(() => {})
            })
            .catch(error => {
              console.error(error)
            })
      }

    },

    fetchIngredients() {
      IngredientService.getAllByPage(this.currentIngredientPage)
          .then(response => {
            this.ingredients = [...response.data.data.ingredients]
            this.totalIngredientPages = response.data.data.num_pages
          })
    },
    onSelectAllIngredientsChange(event) {
      this.selectedIngredients.value = event.checked ? this.ingredients.map((item) => item.id) : [];
      this.selectAllIngredients = event.checked;
    },
    onSelectIngredientsChange(event) {
      this.selectAllIngredients = event.value.length === this.ingredients.length;
    },

    fetchSaladTypes() {
      SaladTypeService.getAllByPage(this.currentSaladTypePage)
          .then(response => {
            this.saladTypes = response.data.data.salad_types
            this.totalSaladTypePages = response.data.data.num_pages
          })
    },
    onSelectAllSaladTypesChange(event) {
      this.selectedSaladTypes.value = event.checked ? this.saladTypes.map((item) => item.id) : [];
      this.selectAllSaladTypes = event.checked;
    },
    onSelectSaladTypeChange(event) {
      this.selectAllSaladTypes = event.value.length === this.saladTypes.length;
    },

    onPageChange(event) {
      this.currentPage = event.page + 1;
      this.fetchSalads();
    },
  }
}
</script>
