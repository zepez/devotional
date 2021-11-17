<template>
	<div class="container">
		<h1>{{ devotional.name }}</h1>
		<div v-html="devotional.html"></div>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";
import api from "../config/axios";

const getById = async (id:string) => {
	const response = await api.get(`/devotional/${id}`)
		.then(res => res.data)
		.catch(e => console.error(e));
	
	return response;
};


export default defineComponent({
	data() {
		return { devotional: { name: "", html: "" } };
	},

	async created() {
		this.devotional = await getById(this.$route.params.id.toString());
	},
});
</script>
