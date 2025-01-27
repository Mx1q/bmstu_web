<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputText v-model="stepName" placeholder="Название шага" :class="{ 'p-invalid': !stepName }" />
      <InputText v-model="stepDescription" placeholder="Описание шага" :class="{ 'p-invalid': !stepDescription }" />
      <Button @click="createRecipeStep" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px; width: 107px">
        Создать
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

import RecipeStepService from "@/services/recipeStep.service.js";
import RecipeStepModel from "@/models/RecipeStepModel.js";

export default {
  name: 'CreateRecipeStepChild',

  props: {
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
    const stepName = ref('Название шага')
    const stepDescription = ref('Описание шага')

    return {
      stepName,
      stepDescription,
      toast: useToast(),
    };
  },
  created() {
  },
  methods: {
    createRecipeStep() {
      const recipeStep = new RecipeStepModel(this.recipe_id, this.stepName, this.stepDescription, 1)

      RecipeStepService.create(recipeStep)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Шаг рецепта успешно создан', life: 3000 });
          })
          .catch(error => {
            let errText = error.response.data.error
            if (error.response.data.error.includes("verifying url")) {
              errText = "не допускается использование ссылок."
            } else if (error.response.data.error.includes("verifying keywords")) {
              errText = "обнаружено нежелательное слово."
            }

            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при создании шага рецепта: ${errText}`, life: 3000 });
          })
    },
  }
}

</script>