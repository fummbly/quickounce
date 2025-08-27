import { User } from "@/actions/apiCalls"


export async function getUsers() {
	try {
		const response = await fetch('http://localhost:8080/api/users')
		const data = await response.json()
		if (!response.ok) {
			throw Error(`Failed to get users: ${data["error"]}`)
		}

		const users: User[] = data

		return users

	} catch (error) {

		console.log(error)
	}
}
