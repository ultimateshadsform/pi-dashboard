<template>
  <div class="flex items-center justify-center relative" v-if="weather">
    <img
      class="w-32 h-32 object-cover absolute -left-[60%] mix-blend-difference select-none"
      :src="`http://openweathermap.org/img/wn/${weather.weather[0].icon}@4x.png`"
      :alt="weather.weather[0].description"
    />
    <div>
      <h3 class="text-2xl font-bold">Weather in {{ city }}</h3>
      <h4>
        Temperature:
        <span class="font-semibold">{{ weather.main.temp }}°C</span>
      </h4>
      <h4>
        Feels like:
        <span class="font-semibold">{{ weather.main.feels_like }}°C</span>
      </h4>
      <h4>
        Humidity:
        <span class="font-semibold">{{ weather.main.humidity }}%</span>
      </h4>
      <h4>
        Wind speed:
        <span class="font-semibold">{{ weather.wind.speed }} m/s</span>
      </h4>
      <h4>
        Weather:
        <span class="font-semibold">{{ weather.weather[0].description }}</span>
      </h4>
    </div>
  </div>
</template>

<script setup lang="ts">
import { WeatherData } from '@/types/weather';
import { useFetch } from '@vueuse/core';
import { onUnmounted, ref } from 'vue';

const weather = ref<WeatherData | null>(null);

const city = 'Skurup';
const units = 'metric';

const url = `https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${
  import.meta.env.VITE_WEATHER_API_KEY
}&units=${units}`;

const fetchWeather = async () => {
  const { data } = await useFetch<WeatherData>(url).get().json<WeatherData>();
  return data;
};

// Update weather every hour
const weatherInterval = setInterval(() => {
  fetchWeather().then((data) => {
    weather.value = data.value;
  });
}, 1000 * 60 * 60);

// Fetch weather on mounted
fetchWeather().then((data) => {
  weather.value = data.value;
});

onUnmounted(() => {
  clearInterval(weatherInterval);
});
</script>
