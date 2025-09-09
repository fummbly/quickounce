'use client'
import { ChangeEvent, FormEvent, useState } from "react"
import Image from "next/image"
import { createPost } from "@/lib/posts"
import { redirect } from "next/navigation"

interface UploadParams {
	params: {
		token: string
	}

}

export default function UploadFile({ params }: UploadParams) {
	const { token } = params
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

		const success = await createPost(token, formData)

		if (!success) {
			return "Something went wrong"
		}

		redirect("/")


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
