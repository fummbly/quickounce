'use client'
import { FormEvent } from "react"
import { createSession } from "@/lib/session";

export default function LoginForm() {

	async function onSubmit(event: FormEvent<HTMLFormElement>) {
		event.preventDefault()

		const formData = new FormData(event.currentTarget)
		const object = Object.fromEntries(formData.entries());
		const response = await fetch('http://localhost:8080/api/login', {
			method: 'POST',
			body: JSON.stringify(object),
		})

		const data = await response.json()
		const expires = Date.parse(data['expires_at'])
		createSession(data['token'], data['refresh_token'], expires)
	}

	return (
		<form onSubmit={onSubmit}>
			<div>
				<label htmlFor="email">Email</label>
				<input id='email' name='email' type='email' placeholder='Email' />
			</div>
			<div>
				<label htmlFor='password'>Password</label>
				<input id='password' name='password' type='password' placeholder='Password' />
			</div>
			<button type='submit'>Login</button>
		</form>
	)

}
