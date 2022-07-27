<script>
	import Modal from './Modal.svelte';
	
	async function get(){
		const response = await fetch("/todo")
		const data = await response.json()
		console.log(data)
	}

	get()
</script>

<main>

	
<Modal />
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h2 {
		color: #000000;
		text-transform: uppercase;
		font-size: 3em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}


</style>