'use client'
import { FormEvent } from "react"
import { createSession, login } from "@/lib/session";
import { redirect } from "next/navigation";

export default function LoginForm() {

	async function onSubmit(event: FormEvent<HTMLFormElement>) {
		event.preventDefault()

		const formData = new FormData(event.currentTarget)

		const success = await login(formData)

		if (!success) {
			return "Something went wrong"
		}

		redirect("/")

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
