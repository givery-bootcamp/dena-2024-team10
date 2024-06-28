import * as VitestConfig from "vitest/config";
import tsconfigPaths from "vite-tsconfig-paths";

export default VitestConfig.defineConfig({
	test: {
		globals: true,
		environment: "jsdom",
		include: ["test/*test.tsx"],
	},
	plugins: [tsconfigPaths()],
});
