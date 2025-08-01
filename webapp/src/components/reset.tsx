
function Reset() {

	async function handleReset() {
		const response = await fetch("http://localhost:8080/admin/reset", {
			method: "POST"
		})


		if (response.ok) {
			console.log("Database reset")
			window.location.reload()
		}

	}


	return (
		<>
			<button onClick={handleReset}>Reset</button>
		</>
	)

}

export default Reset
