<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import MaterialSymbolsSettings from '~icons/material-symbols/settings';

// Add this interface definition
interface Image {
  id: string;
  author: string;
  width: number;
  height: number;
  url: string;
  download_url: string;
}

const router = useRouter()

const images = ref<Image[]>([])
const currentIndex = ref(0)
const weather = ref('Sunny, 25°C')
const joke = ref("Why don't scientists trust atoms? Because they make up everything!")

const visibleImages = computed((): Image[] => {
  const current = images.value[currentIndex.value]
  const next = images.value[(currentIndex.value + 1) % images.value.length]
  return current && next ? [current, next] : []
})

const fetchImages = async () => {
  try {
    const response = await fetch('https://picsum.photos/v2/list?page=1&limit=5')
    if (!response.ok) throw new Error('Failed to fetch images')
    images.value = await response.json()
  } catch (error) {
    console.error('Error fetching images:', error)
  }
}

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % images.value.length
}

let intervalId: NodeJS.Timeout

onMounted(async () => {
  await fetchImages()
  intervalId = setInterval(nextSlide, 5000) // Change slide every 5 seconds
})

onUnmounted(() => {
  clearInterval(intervalId)
})

watch(images, () => {
  if (images.value.length > 0 && currentIndex.value >= images.value.length) {
    currentIndex.value = 0
  }
})

const openSettings = () => {
  router.push('/settings')
}

</script>

<template>
  <div class="w-full h-full flex items-center">
    <div class="relative h-full w-full flex flex-col">
      <!-- Slideshow -->
      <div class="absolute inset-0 flex items-center justify-center w-full h-full">
        <transition-group name="fade" tag="div" class="relative w-full h-full">
          <img v-for="(image, index) in visibleImages" :key="image.id" :src="image.download_url"
            :alt="'Random image ' + image.id"
            class="absolute inset-0 w-full h-full object-cover transition-opacity duration-1000"
            :class="{ 'opacity-100': index === 0, 'opacity-0': index === 1 }" />
        </transition-group>
        <div v-if="!images.length" class="w-full h-full absolute flex items-center justify-center bg-neutral-900">
          <p class="text-white">Loading images...</p>
        </div>
      </div>

      <!-- Overlay content -->
      <div class="absolute inset-0 bg-black bg-opacity-40 p-4 sm:p-6 flex flex-col justify-between">
        <!-- Top bar with Weather and Joke -->
        <div class="flex flex-col sm:flex-row justify-between items-start text-white space-y-4 sm:space-y-0">
          <div class="weather w-full sm:w-1/2">
            <h2 class="text-xl sm:text-2xl font-bold mb-2">Weather</h2>
            <p class="text-sm sm:text-base">{{ weather }}</p>
          </div>
          <div class="joke w-full sm:w-1/2 sm:text-right">
            <h2 class="text-xl sm:text-2xl font-bold mb-2">Joke of the Day</h2>
            <p class="text-sm sm:text-base">{{ joke }}</p>
          </div>
        </div>

        <!-- Bottom bar with Settings button -->
        <div class="flex justify-end">
          <button @click="openSettings" class="text-white hover:text-gray-200 transition-colors duration-200"
            aria-label="Open settings">
            <MaterialSymbolsSettings class="w-6 h-6" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>



<style lang="scss" scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 1s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>