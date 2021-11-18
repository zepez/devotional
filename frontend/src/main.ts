import { createApp } from "vue";
import App from "./App.vue";
import router from "./router/index";
import { createMetaManager } from "vue-meta";
import "./styles/index.css";

createApp(App).use(router).use(createMetaManager()).mount("#app");
