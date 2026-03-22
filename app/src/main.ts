import { createApp, h } from "vue";
import { createMemoryHistory, createRouter, RouterView } from "vue-router";
import Home from "./components/pages/home/Home.vue";
import "./reset.css";
import "./style.css";

const routes = [{ path: "/", component: Home }];

export const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

createApp({ render: () => h(RouterView) })
  .use(router)
  .mount("#app");
