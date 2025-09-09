import { NextResponse } from "next/server";
import { NextRequest } from "next/server";
import { getSession } from "./lib/session";


export async function middleware(request: NextRequest) {

	const token = await getSession()

	if (!token) {
		return NextResponse.redirect(new URL("/login", request.url))
	}

	const requestHeaders = new Headers(request.headers)
	requestHeaders.set('Authorization', token)

	const response = NextResponse.next({
		request: {
			headers: requestHeaders
		}
	})

	return response

}

export const config = {
	matcher: '/upload'
}
