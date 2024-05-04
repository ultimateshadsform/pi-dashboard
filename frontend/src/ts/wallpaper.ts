const wallpapers = [
  'https://assetsio.gnwcdn.com/SONIC_2024Campaign_KeyArt_CMYK_FINAL.png?width=1200&height=1200&fit=bounds&quality=100&format=jpg&auto=webp',
  'https://mrwallpaper.com/images/hd/sonic-shadow-the-hedgehog-knq50hxparr7sf7b.jpg',
  'https://w.wallhaven.cc/full/g7/wallhaven-g7ojze.jpg',
  'https://wallpaperswide.com/download/sonic_frontiers_shadow_dlc_video_game-wallpaper-3440x1440.jpg',
  'https://wallpaperswide.com/download/sonic_the_hedgehog_badass-wallpaper-2560x1440.jpg',
  'https://image.api.playstation.com/vulcan/ap/rnd/202401/3018/64b9dd4395343aee9557174b66677b8f2e703042f70a520f.jpg',
  'https://w.forfun.com/fetch/c2/c29c5a9cf73295253b5fa8c5556d72ae.jpeg',
];

export const getRandomWallpaper = () => {
  return wallpapers[Math.floor(Math.random() * wallpapers.length)];
};
