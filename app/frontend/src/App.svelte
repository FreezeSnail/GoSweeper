<script>
	import { onMount } from "svelte"
	//import { apiData, board} from "./store.js";

	onMount(async () => {
		try {
			let resp = await fetch("http://localhost:8090/1/Map", {
				headers: {
				'Content-Type': 'application/json'
			}})
			console.log(resp)
			let json = await resp.text()
			console.log(json)
		} catch(err) {
			console.error("its fug", err)
		}
})

</script>

<main>
	<h1>Hello!</h1>


</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>