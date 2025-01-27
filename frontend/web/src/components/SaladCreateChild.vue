<template>
  <div class="card flex flex-column md:flex-row gap-3">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-pencil"></i>
      </InputGroupAddon>
      <InputText v-model="saladName" placeholder="Название" :class="{ 'p-invalid': !saladName }" />
      <InputText v-model="saladDescription" placeholder="Описание" :class="{ 'p-invalid': !saladDescription }" />
      <Button @click="createSalad" style="border-top-right-radius: 4px; border-bottom-right-radius: 4px">
        Создать
      </Button>
    </InputGroup>
  </div>
</template>

<script lang="ts">
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { ref } from 'vue';

import SaladModel from "@/models/SaladModel.ts";
import SaladService from "../services/salad.service.ts";

export default {
  name: 'SaladCreateChild',
  components: {
    Button,
    InputGroup,
    InputGroupAddon,
    InputText,
  },
  setup() {
    const saladName = ref('');
    const saladDescription = ref('')

    return {
      saladName,
      saladDescription,
    };
  },
  created() {
  },
  methods: {
    createSalad() {
      const salad = new SaladModel(this.saladName, this.saladDescription)

      SaladService.createSalad(salad)
          .then(response => {
            this.creatingSuccess()
          })
          .catch(error => {
            this.creatingError(error.response.data.error)
          })
    },

    creatingSuccess() {
      this.$emit('creatingSuccess')
    },

    creatingError(errorText) {
      this.$emit('creatingError', errorText)
    }
  }
}

</script>