export function saveRefreshToken(token: string) {
	document.cookie = "refreshToken=" + token + ";HttpOnly"
}


export function getRefreshToken() {
	let decodedCookie = decodeURIComponent(document.cookie)
	let ca = decodedCookie.split(';');
	const name = "refreshToken="
	for (let i = 0; i < ca.length; i++) {
		let c = ca[i]
		while (c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
	}
	return "";
}


let access_token = "";


export function setAccessToken(token: string) {
	access_token = token;
}

export function getAccessToken() {
	return access_token;
}


