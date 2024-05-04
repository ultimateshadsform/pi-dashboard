<template>
  <div v-if="quote">
    <h4 class="text-lg text-balance">
      {{ quote?.quote }}
    </h4>
    <p class="font-semibold">
      — {{ quote?.author }} |
      <span class="capitalize">{{ quote?.category }}</span>
    </p>
  </div>
</template>

<style lang="scss" scoped></style>

<script setup lang="ts">
import { useSettingStore } from '@/stores/settings';
import { Quote } from '@/types/ninja';
import { useFetch } from '@vueuse/core';
import { onUnmounted, ref } from 'vue';
const appSettings = useSettingStore();

const quoteApi = 'https://api.api-ninjas.com/v1/quotes?category=';

const quote = ref<Quote | null>(null);
const categories = ref<string[]>([
  'age',
  'alone',
  'amazing',
  'anger',
  'architecture',
  'art',
  'attitude',
  'beauty',
  'best',
  'birthday',
  'business',
  'car',
  'change',
  'communication',
  'computers',
  'cool',
  'courage',
  'dad',
  'dating',
  'death',
  'design',
  'dreams',
  'education',
  'environmental',
  'equality',
  'experience',
  'failure',
  'faith',
  'family',
  'famous',
  'fear',
  'fitness',
  'food',
  'forgiveness',
  'freedom',
  'friendship',
  'funny',
  'future',
  'god',
  'good',
  'government',
  'graduation',
  'great',
  'happiness',
  'health',
  'history',
  'home',
  'hope',
  'humor',
  'imagination',
  'inspirational',
  'intelligence',
  'jealousy',
  'knowledge',
  'leadership',
  'learning',
  'legal',
  'life',
  'love',
  'marriage',
  'medical',
  'men',
  'mom',
  'money',
  'morning',
  'movies',
  'success',
]);

const fetchQuote = async () => {
  const randomCategory =
    categories.value[Math.floor(Math.random() * categories.value.length)];
  const { data } = await useFetch<Quote[]>(quoteApi + randomCategory, {
    headers: {
      'X-Api-Key': appSettings.quote.apiKey,
    },
  })
    .get()
    .json<Quote[]>();
  return data;
};

const quoteInterval = setInterval(() => {
  fetchQuote().then((data) => {
    if (data.value) {
      quote.value = data.value[0];
    }
  });
  // Fetch new quote 45 minutes
}, 1000 * 60 * 45);

fetchQuote().then((data) => {
  if (data.value) {
    quote.value = data.value[0];
  }
});

onUnmounted(() => {
  clearInterval(quoteInterval);
});
</script>
