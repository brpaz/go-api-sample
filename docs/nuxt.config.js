import theme from '@nuxt/content-theme-docs'

const router = {
	router: {
		base: process.env.NODE_ENV === "production" ? "go-api-sample" : "/"
	}
}

export default theme();
