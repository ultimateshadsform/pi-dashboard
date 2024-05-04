/// <reference types="vite/client" />

// WEATHER_API_KEY .env

interface ImportMetaEnv {
  readonly VITE_WEATHER_API_KEY: string;
  readonly VITE_NINJA_API_KEY: string;
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
