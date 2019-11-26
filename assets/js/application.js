require("expose-loader?$!expose-loader?jQuery!jquery");

import Vue from "vue";
import VueRouter from "router";
import Buefy from 'buefy'

import HomePage from "./pages/home.vue";

$(() => {

  Vue.use(VueRouter);
  Vue.use(Buefy);

  const routes = [
    {path: "/", component: HomePage}
  ];

  const router = new VueRouter({
    mode: "history",
    routes
  });

  const app = new Vue({
    router
  }).$mount("#app");

});
