'use server'
import { cookies, headers } from "next/headers"
var AccessToken = ""



export async function createSession(accessToken: string, refreshToken: string, expiresAt: number) {

	AccessToken = accessToken

	const cookieStore = await cookies()


	cookieStore.set('refreshToken', refreshToken, {
		httpOnly: true,
		secure: true,
		expires: expiresAt,
		sameSite: 'lax',
		path: '/'
	})

}


export async function getSession() {
	console.log("Attempting to refresh token")

	if (AccessToken !== "") {
		return AccessToken
	}


	const refresh = (await cookies()).get('refreshToken')?.value
	console.log(refresh)

	const response = await fetch("http://localhost:8080/api/refresh", {
		headers: {
			"Authorization": "Bearer " + refresh
		}
	})

	const data = await response.json()
	console.log(data)

	AccessToken = data['token']


	return AccessToken
}

export async function login(formData: FormData) {
	const object = Object.fromEntries(formData.entries());
	try {
		const response = await fetch('http://localhost:8080/api/login', {
			method: 'POST',
			body: JSON.stringify(object),
		})
		const data = await response.json()
		if (!response.ok) {
			throw Error(`Failed to login: ${data['error']}`)
		}
		const expires = Date.parse(data['expires_at'])
		await createSession(data['token'], data['refresh_token'], expires)

		return true

	} catch (error) {

		console.log(error)

		return false

	}
}
