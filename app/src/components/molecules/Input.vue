<script lang="ts" setup>
const emit = defineEmits<{
  "update:modelValue": [value: string];
}>();

const props = defineProps({
  label: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
    required: true,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  modelValue: {
    type: String,
    default: "",
  },
});

const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit("update:modelValue", target.value);
};
</script>

<template>
  <div class="input-container">
    <label class="label">{{ props.label }}</label>
    <input
      class="input"
      type="text"
      :placeholder="props.placeholder"
      :disabled="props.disabled"
      :value="props.modelValue"
      @input="handleInput"
    />
  </div>
</template>

<style lang="css" scoped>
.input-container {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
  width: 100%;
}

.label {
  font-weight: 400;
  font-size: 10px;
  line-height: 14px;
  text-transform: uppercase;
  color: #4d505c;
}

.label:has(+ .input:focus) {
  color: var(--blue-base);
}

.input {
  all: unset;
  box-sizing: border-box;
  padding: 16px;
  border: 1px solid var(--gray-300);
  border-radius: 8px;
  width: 100%;

  font-weight: 400;
  font-size: 14px;
  line-height: 18px;
  color: var(--gray-600);
}

.input::placeholder {
  color: var(--gray-400);
}

.input:focus {
  outline: 1px solid var(--blue-base);
}
</style>
