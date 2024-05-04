<template>
  <div
    class="page flex flex-col text-purple-300 overflow-hidden"
    ref="swipeRef"
  >
    <Transition name="fade" appear>
      <img
        :src="wallpaper"
        alt="Shadow the Hedgehog"
        class="background select-none pointer-events-none"
        :key="wallpaper"
      />
    </Transition>
    <RouterView v-slot="{ Component }">
      <Transition
        mode="out-in"
        :name="direction === 'left' ? 'slide-right' : 'slide-left'"
      >
        <KeepAlive>
          <component :is="Component" />
        </KeepAlive>
      </Transition>
    </RouterView>
  </div>
</template>

<style lang="scss" scoped>
.background {
  @apply w-full h-full object-cover absolute -z-10;
  filter: contrast(0.5) brightness(0.5);
}

.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition: all 0.5s ease;
}

.slide-left-enter-from {
  transform: translateX(-100%);
}

.slide-left-leave-to {
  transform: translateX(100%);
}

.slide-right-enter-from {
  transform: translateX(100%);
}

.slide-right-leave-to {
  transform: translateX(-100%);
}

.fade-enter-active,
.fade-leave-active {
  transition: all 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

<script setup lang="ts">
import { useSwipe } from '@vueuse/core';
import { onUnmounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { getRandomWallpaper } from '@/ts/wallpaper';

const router = useRouter();

const swipeRef = ref<HTMLDivElement | null>(null);
const wallpaper = ref(getRandomWallpaper());

const pages = router.getRoutes();

const wallpaperTimeout = setInterval(() => {
  wallpaper.value = getRandomWallpaper();
}, 1000 * 60 * 20);

const enableSwipe = ref(true);
const { isSwiping, direction, lengthX } = useSwipe(swipeRef, {
  onSwipe() {
    if (isSwiping.value && enableSwipe.value) {
      if (direction.value === 'left' && Math.abs(lengthX.value) > 100) {
        enableSwipe.value = false;
        const currentPage = pages.findIndex(
          (page) => page.path === router.currentRoute.value.path
        );

        const nextPage = pages[currentPage + 1];

        if (nextPage) {
          router.push(nextPage.path);
        } else {
          router.push(pages[0].path);
        }
      } else if (direction.value === 'right' && Math.abs(lengthX.value) > 100) {
        enableSwipe.value = false;
        const currentPage = pages.findIndex(
          (page) => page.path === router.currentRoute.value.path
        );

        const prevPage = pages[currentPage - 1];

        if (prevPage) {
          router.push(prevPage.path);
        } else {
          router.push(pages[pages.length - 1].path);
        }
      }
    }
  },
  onSwipeEnd() {
    enableSwipe.value = true;
  },
});

onUnmounted(() => {
  clearInterval(wallpaperTimeout);
});
</script>
