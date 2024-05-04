<template>
  <div class="page p-4 flex flex-col items-center justify-center gap-y-5">
    <h1 class="text-4xl font-bold">Settings</h1>
    <div class="flex flex-col gap-y-2">
      <div>
        <div class="flex align-items-center">
          <Checkbox
            :binary="true"
            v-model="enableWeather"
            inputId="weather"
            value="Weather"
          />
          <label for="weather" class="ml-2"> Enable weather</label>
        </div>
        <div class="flex align-items-center">
          <Checkbox
            :binary="true"
            v-model="enableQuote"
            inputId="quote"
            value="Quote"
          />
          <label for="quote" class="ml-2"> Enable quote</label>
        </div>
      </div>
      <Button class="w-full" label="Save" @click="saveSettings" />
    </div>
  </div>
</template>

<style lang="scss" scoped></style>

<script setup lang="ts">
import Button from 'primevue/button';
import Checkbox from 'primevue/checkbox';

import { useSettingStore } from '@/stores/settings';
import { ref } from 'vue';

const settings = useSettingStore();

const enableWeather = ref(settings.weather.enabled);
const enableQuote = ref(settings.quote.enabled);

const saveSettings = () => {
  settings.$patch({
    weather: { enabled: enableWeather.value },
    quote: { enabled: enableQuote.value },
  });
};
</script>
