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
