import { useState, useEffect } from "react"

function Users() {
	interface User {
		id: string;
		created_at: string;
		updated_at: string;
		email: string;
		username: string;
	}

	const [users, setUsers] = useState(Array<User>)

	useEffect(() => {
		fetch("http://localhost:8080/api/users")
			.then(response => response.json())
			.then(json => {
				setUsers(json)
			})
			.catch(error => console.log(error))
	}, []);



	return (
		<>
			{users ? <pre>{users.map((user) => <div key={user.id}>Username: {user.username} Email: {user.email}</div>)}</pre> : 'Loading ...'}
		</>
	)

}

export default Users
