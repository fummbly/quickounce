import type { NextConfig } from "next";

const nextConfig: NextConfig = {
	/* config options here */
};

module.exports = {
	images: {
		remotePatterns: [new URL('http://localhost:8080/uploads/**')],
	},
}

export default nextConfig;
