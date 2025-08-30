'use client'
import { createUser } from "@/lib/users";
import { FormEvent } from "react";
import NotFound from "./not-found";
import { redirect } from "next/navigation";

export default function CreateUserForm() {


	async function onSubmit(event: FormEvent<HTMLFormElement>) {
		event.preventDefault()
		const formData = new FormData(event.currentTarget)
		const user = await createUser(formData)
		if (!user) {
			return "Something went wrong"
		}
		redirect("/login")
	}


	return (
		<form onSubmit={onSubmit}>
			<label htmlFor="username">Username</label>
			<input name="username" id="username" type="text" placeholder="Username" />
			<label htmlFor="email">Email</label>
			<input name="email" id="email" type="email" placeholder="Email" />
			<label htmlFor="password">Password</label>
			<input name="password" id="password" type="password" placeholder="Password" />
			<button type="submit">Submit</button>

		</form>
	)
}
