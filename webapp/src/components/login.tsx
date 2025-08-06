import { getAccessToken, getRefreshToken, saveRefreshToken, setAccessToken } from "../utils/tokens"


async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
	event.preventDefault()

	const formData = new FormData(event.target)

	const jsonData = JSON.stringify(Object.fromEntries(formData))

	try {
		const response = await fetch("http://localhost:8080/api/login", {
			method: "POST",
			body: jsonData
		})

		const data = await response.json()

		if (data) {
			setAccessToken(data.token);
			saveRefreshToken(data.refresh_token)

		}
	} catch (error) {
		console.log(error)
	}

	console.log(getAccessToken())
	console.log(getRefreshToken())




}


function Login() {
	return (
		<>
			<h2>Login</h2>
			<form onSubmit={handleSubmit}>
				<label>Email:</label>
				<input type="email" name="email" placeholder="johndoe@email.com" />
				<label>Password:</label>
				<input type="password" name="password" placeholder="password" />
				<button type="submit">Submit</button>
			</form>

		</>
	)
}


export default Login
