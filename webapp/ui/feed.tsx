
import { User } from '@/actions/apiCalls'
import UserPreview from './user-preview'


export default async function Feed() {

	const response = await fetch('http://localhost:8080/api/users')
	const users: User[] = await response.json()


	return (
		<div className='flex flex-col w-full m-6'>
			{users.map((user) => (
				< UserPreview key={user.id} params={{ user }} />
			))}
		</div>
	)
}
