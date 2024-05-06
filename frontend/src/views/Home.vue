<template>
  <div
    class="page flex items-center justify-evenly p-4"
    :class="{ 'gap-y-8': !enableWeather }"
  >
    <TransitionGroup name="fade" appear>
      <div class="flex flex-col justify-center items-center" key="time">
        <h1 class="font-semibold text-4xl flex items-center gap-x-4">
          <span
            >{{ time.hours }}<span class="blink">:</span
            >{{ time.minutes }}</span
          >
          <span class="text-xl">{{ time.date }}</span>
        </h1>
        <QuoteComponent v-if="enableQuote" key="quote" class="max-w-lg" />
      </div>
      <WeatherComponent v-if="enableWeather" key="weather" />
    </TransitionGroup>
  </div>
</template>

<style lang="scss" scoped>
.fade-move,
.fade-enter-active,
.fade-leave-active {
  transition: all 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.fade-leave-active {
  position: absolute;
}

.blink {
  animation: blink 2s infinite ease-in-out;

  @keyframes blink {
    0% {
      opacity: 0;
    }
    50% {
      opacity: 1;
    }
    100% {
      opacity: 0;
    }
  }
}
</style>

<script setup lang="ts">
import { DateTime } from 'luxon';
import { computed, onUnmounted, ref } from 'vue';
import { useSettingStore } from '@/stores/settings';
import QuoteComponent from '@/components/Home/QuoteComponent.vue';
import WeatherComponent from '@/components/Home/WeatherComponent.vue';

const settings = useSettingStore();
const enableWeather = computed(() => settings.weather.enabled);
const enableQuote = computed(() => settings.quote.enabled);

// Reactive time that changes every second
// Show current time in HH:mm format
// Show date: August 1, 2021
const time = ref({
  hours: DateTime.local().toFormat('HH'),
  minutes: DateTime.local().toFormat('mm'),
  date: DateTime.local().toFormat('MMMM d, yyyy'),
});

const timeInterval = setInterval(() => {
  time.value = {
    hours: DateTime.local().toFormat('HH'),
    minutes: DateTime.local().toFormat('mm'),
    date: DateTime.local().toFormat('MMMM d, yyyy'),
  };
}, 1000 * 10);

onUnmounted(() => {
  clearInterval(timeInterval);
});
</script>
