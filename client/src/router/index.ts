import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";
import Sign from "../views/Sign.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/sign",
    name: "Sign",
    component: Sign,
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
