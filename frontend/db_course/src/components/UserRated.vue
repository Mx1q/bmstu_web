<template>
  <DataView :value="salads" paginator :rows="rows" :totalRecords="totalPages*rows" @page="onPageChange">
    <template #empty>
      <div class="d-flex justify-content-center align-items-center center">
        <div class="flex">
        <span class="font-medium text-primary text-xl">
          Вы еще не оценили ни одного салата
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

import Button from 'primevue/button';
import DataView from 'primevue/dataview';

import { ref } from 'vue';

export default {
  name: 'UserRatedPage',
  components: {
    Button,
    DataView,
  },
  data() {
    return {
      rows: 30,
      totalPages: 0,
      salads: [],
      currentPage: 1,
    }
  },
  created() {
    this.fetchSalads()
  },
  methods: {
    fetchSalads() {
      let status = 4
      SaladService.getUserRated(this.currentPage)
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
    },

    onPageChange(event) {
      this.currentPage = event.page + 1;
      this.fetchSalads();
    },
  }
}
</script>
