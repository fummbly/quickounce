import { User, UserIDResponse } from "@/actions/apiCalls";
import PostPreview from "./post-preview";



interface UserPreviewProps {
	params: {
		user: User
	};
}

export default async function UserPreview({ params }: UserPreviewProps) {

	const { user } = params

	const response = await fetch(`http://localhost:8080/api/users/${user.username}`)
	const data: UserIDResponse = await response.json()

	console.log(data)



	const posts = data.Posts


	return (
		<div className="flex items-center justify-between w-full p-5">
			<div>
				<h3>{user.username}</h3>
				<h6 className="text-gray-600">{user.email}</h6>

				<PostPreview params={{ posts }} />
			</div>
		</div>
	)
}
