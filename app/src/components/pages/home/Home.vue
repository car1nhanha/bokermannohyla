<script setup lang="ts">
import { faDownload } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { computed, ref } from "vue";
import Button from "../../atoms/button.vue";
import Input from "../../molecules/Input.vue";
import History from "../../organisms/history.vue";

const links = [
  {
    id: 1,
    original: "devsite.portfolio.com.br/devname-123456",
    shorted: "brev.ly/Portfolio-Dev",
    visits: 123,
  },
  {
    id: 2,
    original: "devsite.portfolio.com.br/devname-123456",
    shorted: "brev.ly/Portfolio-Dev",
    visits: 123,
  },
];

const input = ref("");

const inputIsValid = computed(() => {
  return /^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$/.test(
    input.value,
  );
});
</script>

<template>
  <div class="content">
    <div class="header">
      <h1>thoropa</h1>
    </div>
    <div class="main">
      <div class="form">
        <h2 class="title">New link</h2>

        <div class="form-content">
          <Input label="Original link" placeholder="www.example.com" v-model="input" />
          <Input label="shortened link" placeholder="thoropa.com/" :disabled="true" />
        </div>

        <Button size="large" @click="() => {}" :disabled="!inputIsValid">Shorten</Button>
      </div>

      <div class="history">
        <div class="history-header">
          <h2 class="title">My links</h2>
          <Button size="small" @click="() => {}">
            <template #icon>
              <FontAwesomeIcon :icon="faDownload" />
            </template>
            Download CSV
          </Button>
        </div>

        <History :links="links" />
      </div>
    </div>
  </div>
</template>

<style src="./home.css" lang="css" scoped></style>
