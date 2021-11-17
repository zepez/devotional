<template>
	<Header />
	<Copyright />
  <router-view />

	<Copyright />
</template>

<style>
#app {
  font-size: 1.2rem;
	background-color: #E4E7EB;
	color: #222222;
	min-height: 100vh;
}
</style>


<script lang="ts">
import Header from "./components/HeaderLayout.vue";
import Copyright from "./components/CopyRight.vue";

export default {
	components: {
		Copyright,
		Header
	}
};
</script>
