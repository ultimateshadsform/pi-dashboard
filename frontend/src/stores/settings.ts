import { useStorage } from '@vueuse/core';
import { defineStore } from 'pinia';

export const useSettingStore = defineStore('settings', {
  // arrow function recommended for full type inference
  state: () => {
    return {
      // all these properties will have their type inferred automatically
      quote: useStorage('quote', {
        enabled: true,
        apiKey: import.meta.env.VITE_NINJA_API_KEY,
      }),

      weather: useStorage('weather', {
        enabled: true,
        apiKey: import.meta.env.VITE_WEATHER_API_KEY,
      }),
    };
  },
});
