

export interface User {
	id: string
	createt_at: string
	updated_at: string
	email: string
	username: String
}

export interface UsersResponse {
	users: User[]
}


export interface Post {
	user_id: string
	id: string
	created_at: string
	updated_at: string
	image_url: string
}


export interface UserIDResponse {
	User: User
	Posts: Post[]
}



