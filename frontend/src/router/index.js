import { createWebHistory, createRouter } from "vue-router";
import Home from "@/views/Home.vue";
import Files from "@/views/Files.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/files",
    name: "Files",
    component: Files,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  linkActiveClass: "active", // active class for non-exact links.
  linkExactActiveClass: "active" // active class for *exact* links.
});

export default router;