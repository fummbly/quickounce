import './App.css'
import Users from './components/users'
import CreateUser from './components/createUser'
import Reset from './components/reset'
import Login from './components/login'

function App() {

	return (
		<>
			<Reset />
			<CreateUser />
			<Login />
			<Users />
		</>
	)
}

export default App
