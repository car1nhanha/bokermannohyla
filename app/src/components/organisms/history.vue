<script lang="ts" setup>
import { faCopy, faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useToast } from "vue-toast-notification";
import { ApiService } from "../../service/api";
import Button from "../atoms/button.vue";

const appURL = import.meta.env.VITE_APP_URL;
const toast = useToast();

const props = defineProps<{
  links: {
    Id: string;
    Original: string;
    Accesses: number;
  }[];
  reloadLinks: () => void;
}>();

const copyToClipboard = (text: string) => {
  navigator.clipboard.writeText(text);
  toast.success("Link copied to clipboard!");
};

const deleteLink = (id: string) => {
  const confirmed = confirm("Are you sure you want to delete this link?");
  if (!confirmed) return;

  ApiService.deleteLink(id)
    .then(() => {
      toast.success("Link deleted successfully!");
      props.reloadLinks();
    })
    .catch(() => {
      toast.error("Failed to delete the link. Please try again.");
    });
};
</script>

<template>
  <section class="history-container">
    <div v-for="link in $props.links" class="history-item">
      <div class="left">
        <span class="shorted">{{ appURL }}/{{ link.Id }}</span>
        <span class="original">{{ link.Original }}</span>
      </div>

      <div class="right">
        <span class="visits">{{ link.Accesses }} accesses</span>
        <div class="actions">
          <Button size="small" @click="() => copyToClipboard(`${appURL}/${link.Id}`)">
            <template #icon>
              <FontAwesomeIcon :icon="faCopy" />
            </template>
          </Button>
          <Button size="small" @click="() => deleteLink(link.Id)">
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

.history-container {
  max-height: 400px;
  overflow: scroll;
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
