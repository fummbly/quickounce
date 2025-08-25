import { FormEvent } from "react"

export async function login(formData: FormData) {
	'use server'
	console.log(formData)
}
