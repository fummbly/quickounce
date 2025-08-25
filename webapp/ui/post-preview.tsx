
import { Post } from "@/actions/apiCalls"
import PostThumbnail from "./post-thumbnail"


interface PostPreviewParams {
	params: {
		posts: Post[]
	}
}

export default async function PostPreview({ params }: PostPreviewParams) {
	const { posts } = params


	return (
		<div className="flex gap-2">
			{posts.map((post) => (
				<PostThumbnail key={post.id} params={{ post }} />
			))}
		</div>
	)
}
