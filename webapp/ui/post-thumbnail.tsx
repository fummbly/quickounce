import Image from "next/image"
import { Post } from "@/actions/apiCalls"

interface PostThumbnailParams {
	params: {
		post: Post
	}
}


export default function PostThumbnail({ params }: PostThumbnailParams) {

	const { post } = params

	return (
		<div className="w-25 h-25 bg-cyan-600" >
			<Image src={`http://localhost:8080/uploads/${post.image_url}.png`} alt="user post" width={100} height={100} />
			<h1>{post.image_url} </h1>
		</div>
	)

}
