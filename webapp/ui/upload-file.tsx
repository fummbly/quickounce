'use client'
import { ChangeEvent, FormEvent, useState } from "react"
import Image from "next/image"
import { getSession } from "@/lib/session"

export default function UploadFile() {

	const [backgroundImage, setBackgroundImage] = useState("")

	function onChange(e: ChangeEvent<HTMLInputElement>) {
		const file = e.target.files[0]

		if (file) {
			setBackgroundImage(URL.createObjectURL(file))
		}

	}

	async function onSubmit(event: FormEvent<HTMLFormElement>) {
		event.preventDefault()

		const formData = new FormData(event.currentTarget)

		const token = await getSession()
		console.log(token)

		try {
			const response = await fetch("http://localhost:8080/api/posts", {
				method: 'POST',
				headers: {
					"Authorization": "Bearer " + token,
				},
				body: formData,
			})

			if (response.ok) {
				console.log("Photo Successfully uploaded")
			}
		} catch (e) {
			console.log(e)
		}
	}

	return (
		<div>
			<div className="justify-center border-2 bg-blue-300">
				<Image
					src={backgroundImage ? backgroundImage : null}
					width={300}
					height={300}
					alt="image preview" />
			</div>
			<div>
				<form onSubmit={onSubmit}>
					<input name='photo' type="file" onChange={onChange} />
					<input name='submit' type="submit" />
				</form>
			</div>
		</div>
	)
}
