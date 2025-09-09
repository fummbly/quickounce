

export async function createPost(token: string, formData: FormData) {

	try {

		const response = await fetch("http://localhost:8080/api/posts", {
			method: 'POST',
			headers: {
				'Authorization': 'Bearer ' + token
			},
			body: formData,
		})

		if (!response.ok) {
			throw Error("Failed to upload photo")
		}

		return true
	} catch (e) {
		return false
	}


}
