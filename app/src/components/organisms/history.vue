<script lang="ts" setup>
import { faCopy, faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import Button from "../atoms/button.vue";

const props = defineProps<{
  links: {
    id: number;
    original: string;
    shorted: string;
    visits: number;
  }[];
}>();
</script>

<template>
  <section class="history-container">
    <div v-for="link in $props.links" class="history-item">
      <div class="left">
        <span class="shorted">{{ link.shorted }}</span>
        <span class="original">{{ link.original }}</span>
      </div>

      <div class="right">
        <span class="visits">{{ link.visits }} accesses</span>
        <div class="actions">
          <Button size="small" @click="() => {}">
            <template #icon>
              <FontAwesomeIcon :icon="faCopy" />
            </template>
          </Button>
          <Button size="small" @click="() => {}">
            <template #icon>
              <FontAwesomeIcon :icon="faTrash" />
            </template>
          </Button>
        </div>
      </div>
    </div>
  </section>
</template>

<style lang="css" scoped>
.history-container,
.history-item {
  width: 100%;
}

.history-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  gap: 16px;
  padding: 16px 0;
  border-top: 1px solid var(--gray-200);
}

.left,
.right {
  overflow: hidden;
  flex: 1;
}

/* left */
.left {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.left .shorted,
.left .original {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.left .shorted {
  font-weight: 600;
  font-size: 14px;
  line-height: 18px;
  color: var(--blue-base);
}
.left .original {
  font-weight: 400;
  font-size: 12px;
  line-height: 16px;
  color: var(--gray-500);
}

/* right */
.right {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-end;
  gap: 20px;
}
.right .visits {
  font-weight: 400;
  font-size: 12px;
  line-height: 16px;
  color: var(--gray-500);
}

.actions {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 4px;
}
</style>
