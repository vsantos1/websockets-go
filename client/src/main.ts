import { createApp } from "vue";
import "./assets/css/reset.css";
import "./assets/css/main.css";
import router from "./router";
import App from "./App.vue";

createApp(App).use(router).mount("#app");
