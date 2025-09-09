import UploadFile from "@/ui/upload-file";
import { headers } from "next/headers";
import { redirect } from "next/navigation";


export default async function upload() {

	const header = await headers()
	const middlewareToken = header.get('Authorization')
	if (!middlewareToken) {
		redirect("/login")
	}


	return (
		<div>
			<UploadFile params={{ token: middlewareToken }} />
		</div>
	)
}
