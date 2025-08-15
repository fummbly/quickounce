'use client'
import { ChangeEvent, useState } from "react"
import Image from "next/image"

export default function UploadFile() {

	const [backgroundImage, setBackgroundImage] = useState("")

	function onChange(e: ChangeEvent<HTMLInputElement>) {
		const file = e.target.files[0]


		if (file) {
			setBackgroundImage(URL.createObjectURL(file))
		}

	}

	return (
		<div>
			<div>
				<Image
					src={backgroundImage ? backgroundImage : null}
					width={300}
					height={300}
					alt="image preview" />
			</div>
			<div>
				<input name='filepicker' type="file" onChange={onChange} />
			</div>
		</div>
	)
}
