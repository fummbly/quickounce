
import NotFound from './not-found'
import UserPreview from './user-preview'
import { getUsers } from '@/lib/users'


export default async function Feed() {

	const users = await getUsers()

	if (!users) {
		return NotFound()
	}

	return (
		<div className='bg-gray-800 flex'>
			<h2>Users</h2>
			<div className='flex flex-col w-full m-6 border-2 border-gray-400'>
				{users?.map((user) => (
					< UserPreview key={user.id} params={{ user }} />
				))}
			</div>
		</div>
	)
}
