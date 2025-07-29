
async function test() {
	const url = "/api/users"
	try {
		const response = await fetch(url);
		if (!response.ok) {
			throw new Error(`Response Status: ${response.status}`);
		}

		const json = await response.json();
		console.log(json);
		var userString = ""
		Object.keys(json).forEach(key => {
			userString += `${json[key]["username"]}\n`
		})

		document.getElementById("users").innerText = userString
	} catch (error) {
		console.log(error.message);
	}

}


async function createUser() {

	username = document.getElementById("username").value
	email = document.getElementById("email").value
	password = document.getElementById("password").value

	const url = "/api/users"
	try {
		const response = await fetch(url, {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				username: username,
				email: email,
				password: password
			})
		});
		if (!response.ok) {
			console.log(`Response not ok: ${response.status}`)
		}

		const json = await response.json();
		console.log(json);
	} catch (error) {
		console.log(error)
	}


}
