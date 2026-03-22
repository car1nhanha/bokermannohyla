<script lang="ts" setup>
defineOptions({
  inheritAttrs: false,
});

const props = withDefaults(
  defineProps<{
    size?: "small" | "large";
  }>(),
  {
    size: "small",
  },
);
</script>

<template>
  <button class="button" :class="props.size" v-bind="$attrs">
    <slot name="icon"></slot>
    <span v-if="$slots.default"><slot></slot></span>
  </button>
</template>

<style lang="css" scoped>
.button {
  all: unset;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  width: 100%;
  padding: 15px;
  gap: 12px;
  background: var(--blue-base);
  border-radius: 8px;
  cursor: pointer;

  box-sizing: border-box;

  font-weight: 600;
  font-size: 14px;
  line-height: 18px;
  color: #ffffff;
  transition: 0.3s;
}
.button:hover:not(:disabled) {
  background: var(--blue-dark);
}

.button:disabled {
  cursor: not-allowed;
  opacity: 0.5;
}

.button.small {
  max-width: min-content;
  background-color: var(--gray-200);
  color: var(--gray-500);
  text-wrap: nowrap;
  padding: 8px;
  gap: 6px;
  border-radius: 4px;
  font-size: 12px;
  line-height: 16px;
  border: 1px dashed transparent;
}

.button.small:hover:not(:disabled) {
  border: 1px dashed #9747ff;
  background-color: var(--gray-200);
}

.button.small:disabled {
  opacity: 0.5;
}
</style>
