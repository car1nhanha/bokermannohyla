import { createApp, h } from "vue";
import { createRouter, createWebHistory, RouterView } from "vue-router";
import ToastPlugin from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";
import Home from "./components/pages/home/Home.vue";
import Redirect from "./components/pages/redirect/redirect.vue";
import "./reset.css";
import "./style.css";

const routes = [
  { path: "/", component: Home },
  { path: "/:id", component: Redirect },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp({ render: () => h(RouterView) })
  .use(ToastPlugin)
  .use(router)
  .mount("#app");
