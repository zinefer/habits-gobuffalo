<template>
    <b-navbar type="is-dark">
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '/' }">
                Habits
            </b-navbar-item>
        </template>

        <!--<template slot="start">
            <b-navbar-item href="/">
                Home
            </b-navbar-item>
        </template>-->

        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <a class="button is-primary">
                        <strong>Sign up</strong>
                    </a>
                    <a class="button is-light">
                        Log in
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script charset="utf-8">
export default {
  data() {
    return { };
  },
  created() { },
  watch: { },
  methods: { }
};
</script>