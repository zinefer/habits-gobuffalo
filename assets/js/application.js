require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

import Vue from "vue";
import VueRouter from "router";
import BootstrapVue from "bootstrap-vue";

import HomeComponent from "./components/home.vue";

$(() => {

  Vue.use(VueRouter);
  Vue.use(BootstrapVue);

  const routes = [
    {path: "/", component: HomeComponent}
  ];

  const router = new VueRouter({
    mode: "history",
    routes
  });

  const app = new Vue({
    router
  }).$mount("#app");

});
