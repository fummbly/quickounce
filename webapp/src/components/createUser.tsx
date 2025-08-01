
function CreateUser() {

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault()
		const formData = new FormData(event.target)

		const jsonData = JSON.stringify(Object.fromEntries(formData))


		const response = await fetch("http://localhost:8080/api/users", {
			method: "POST",
			body: jsonData
		})

		if (response.ok) {
			console.log("User Created")
		}

		window.location.reload()

	}

	return (
		<>
			<h2>Create User</h2>
			<form onSubmit={handleSubmit}>
				<label>Email: </label>
				<input type="email" name="email" />
				<label>Username: </label>
				<input type="username" name="username" />
				<label>Password: </label>
				<input type="password" name="password" />
				<button type="submit">Submit</button>
			</form>
		</>
	)

}

export default CreateUser
