<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputText v-model="saladName" placeholder="Название" :class="{ 'p-invalid': !saladName }" />
      <InputText v-model="saladDescription" placeholder="Описание" :class="{ 'p-invalid': !saladDescription }" />
      <Button @click="updateSalad" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px">
        Обновить
      </Button>
    </InputGroup>
  </div>
</template>

<script lang="ts">
import SaladModel from "@/models/SaladModel.ts";
import SaladService from "../services/salad.service.ts";
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { ref } from 'vue';
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';

export default {
  name: 'UpdateSaladChild',

  props: {
    salad_id: String,
  },

  components: {
    Button,
    InputGroup,
    InputGroupAddon,
    InputText,
    Toast
  },
  setup() {
    const saladName = ref('');
    const saladDescription = ref('')
    const toast = useToast()

    return {
      saladName,
      saladDescription,
      toast,
    };
  },
  created() {
    this.fillInfo();
  },
  methods: {
    fillInfo() {
      SaladService.getSaladDetails(this.salad_id)
          .then(response => {
            const salad = response.data.data.salad
            this.saladName = salad.name
            this.saladDescription = salad.description
          })
          .catch(error => {
            console.error(error)
          })
    },

    updateSalad() {
      const salad = new SaladModel(this.saladName, this.saladDescription)
      SaladService.updateSalad(this.salad_id, salad)
          .then(response => {
            this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Салат успешно обновлен', life: 3000 });
            this.updatingSuccess(response.data.data.salad)
          })
          .catch(error => {
            let errText = error.response.data.error
            if (error.response.data.error.includes("verifying url")) {
              errText = "не допускается использование ссылок."
            } else if (error.response.data.error.includes("verifying keywords")) {
              errText = "обнаружено нежелательное слово."
            }

            this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при обновлении салата: ${errText}`, life: 3000 });
          })
    },

    updatingSuccess(salad) {
      this.$emit('updatingSuccess', salad)
    },

    updatingError(errorText) {
      this.$emit('updatingError', errorText)
    }
  }
}

</script>