<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputNumber v-model="nos" placeholder="Количество порций" :min="1" prefix="Количество порций: " />
      <InputNumber v-model="ttc" placeholder="Время приготовления (мин.)" :min="1" prefix="Время приготовления (мин.): " />
      <Button @click="updateRecipe" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px">
        Обновить
      </Button>
    </InputGroup>
  </div>
</template>

<script>
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputNumber from 'primevue/inputnumber';
import Button from 'primevue/button';
import { ref } from 'vue';
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';

import RecipeModel from "@/models/RecipeModel.js";
import RecipeService from "@/services/recipe.service.js";

export default {
  name: 'UpdateRecipeChild',

  props: {
    salad_id: String,
    recipe_id: String,
  },

  components: {
    Button,
    InputGroup,
    InputGroupAddon,
    InputNumber,
    Toast,
  },
  setup() {
    const nos = ref(0)
    const ttc = ref(0)

    return {
      nos,
      ttc,
      toast: useToast(),
    };
  },
  created() {
    this.fillInfo();
  },
  methods: {
    fillInfo() {
      RecipeService.getRecipeDetails(this.salad_id)
          .then(response => {
            const recipe = response.data.data.recipe
            this.nos = recipe.number_of_servings
            this.ttc = recipe.time_to_cook
          })
          .catch(error => {
            console.error(error)
          })
    },

    updateRecipe() {
      const recipe = new RecipeModel(this.nos, this.ttc)

      // FIXME: remove me
      console.log(`Bearer ${localStorage.getItem('user').replace(/"/g, '')}`)

      RecipeService.updateRecipe(this.recipe_id, recipe)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Рецепт успешно обновлен', life: 3000 });
            this.updatingSuccess(response.data.data.recipe)
          })
          .catch(error => {
            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при обновлении рецепта: ${error.response.data.error}`, life: 3000 });
            this.updatingError(error.response.data.error)
          })
    },

    updatingSuccess(recipe) {
      this.$emit('updatingSuccess', recipe)
    },

    updatingError(errorText) {
      this.$emit('updatingError', errorText)
    }
  }
}

</script>