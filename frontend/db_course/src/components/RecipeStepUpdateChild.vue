<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputText v-model="stepName" placeholder="Название шага" :class="{ 'p-invalid': !stepName }" />
      <InputText v-model="stepDescription" placeholder="Описание шага" :class="{ 'p-invalid': !stepDescription }" />
      <Button @click="updatingCancelled">
        отмена
      </Button>
      <Button @click="updateStep" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px">
        Обновить
      </Button>
    </InputGroup>
  </div>
</template>

<script>
import RecipeStepModel from "@/models/RecipeStepModel.js";
import RecipeStepService from "@/services/recipeStep.service.js";
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { ref } from 'vue';
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';

export default {
  name: 'UpdateRecipeStepChild',

  props: {
    step_id: String,
    step_name_prop: String,
    step_description_prop: String
  },

  components: {
    Button,
    InputGroup,
    InputGroupAddon,
    InputText,
    Toast,
  },
  setup() {
    const stepName = ref('');
    const stepDescription = ref('')

    return {
      stepName,
      stepDescription,
      toast: useToast(),
    };
  },
  created() {
    this.fillInfo();
  },
  methods: {
    fillInfo() {
      this.stepName = this.step_name_prop
      this.stepDescription = this.step_description_prop
    },

    updateStep() {
      const step = new RecipeStepModel(0, this.stepName, this.stepDescription)

      RecipeStepService.update(this.step_id, step)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Шаг рецепта успешно обновлен', life: 3000 });
            this.updatingSuccess(response.data.data.step)
          })
          .catch(error => {
            let errText = error.response.data.error
            if (error.response.data.error.includes("verifying url")) {
              errText = "не допускается использование ссылок."
            } else if (error.response.data.error.includes("verifying keywords")) {
              errText = "обнаружено нежелательное слово."
            }

            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при обновлении шага рецепта: ${errText}`, life: 3000 });
            this.updatingError(error.response.data.error)
          })
    },

    updatingSuccess(updatedStep) {
      this.$emit('updatingSuccess', updatedStep)
    },

    updatingCancelled() {
      this.$emit('updatingCancelled')
    },

    updatingError(errorText) {
      this.$emit('updatingError', errorText)
    }
  }
}

</script>