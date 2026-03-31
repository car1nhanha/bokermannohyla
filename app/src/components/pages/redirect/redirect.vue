<script lang="ts" setup>
import { ref, watchEffect } from "vue";
import { useRoute } from "vue-router";
import { ApiService } from "../../../service/api";
import TemplateCard from "../../templates/TemplateCard.vue";

const route = useRoute();
const id = route.params.id;

let success = ref(true);

watchEffect(() => {
  ApiService.getLink(id as string)
    .then((response) => {
      window.location.href = `${response.Original}`;
    })
    .catch(() => {
      success.value = false;
    });
});
</script>

<template>
  <TemplateCard title="Redirecting..." v-if="success">
    <template #image>
      <img src="/logo.svg" alt="Loading" />
    </template>
    <p class="description">
      The link will open automatically in a few moments.<br />
      Didn't get redirected? <RouterLink to="/">Access it here.</RouterLink>
    </p>
  </TemplateCard>

  <!-- 404 -->

  <TemplateCard title="Link not found" v-if="!success">
    <template #image>
      <img src="/404.svg" alt="Loading" />
    </template>
    <p class="description">
      The link you are trying to access does not exist, has been removed,<br />
      or is an invalid URL. Learn more at <RouterLink to="/">here.</RouterLink>
    </p>
  </TemplateCard>
</template>
