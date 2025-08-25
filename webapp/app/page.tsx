import Feed from "@/ui/feed";
import LoginForm from "@/ui/login-form";
import UploadFile from "@/ui/upload-file";

export default function Home() {
	return (
		<div className="font-sans flex">

			<Feed />
			<UploadFile />
			<LoginForm />

		</div>
	);
}
