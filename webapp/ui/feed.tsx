
import NotFound from './not-found'
import UserPreview from './user-preview'
import { getUsers } from '@/lib/users'


export default async function Feed() {

	const users = await getUsers()

	if (!users) {
		return NotFound()
	}

	return (
		<div className='flex flex-col w-full m-6'>
			{users?.map((user) => (
				< UserPreview key={user.id} params={{ user }} />
			))}
		</div>
	)
}
