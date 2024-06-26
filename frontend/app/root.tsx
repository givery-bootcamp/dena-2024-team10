import type { LinksFunction, LoaderFunctionArgs } from "@remix-run/node";
import {
	Links,
	Meta,
	Outlet,
	Scripts,
	ScrollRestoration,
	redirect,
	useLoaderData,
	useRouteError,
} from "@remix-run/react";
import { AxiosError } from "axios";
import classNames from "classnames";
import stylesheet from "~/tailwind.css?url";
import apiClient from "./apiClient/apiClient";
import Header from "./components/header";

export const links: LinksFunction = () => [
	{ rel: "stylesheet", href: stylesheet },
];

export async function loader({
	request,
}: LoaderFunctionArgs): Promise<{ id?: number; username?: string }> {
	try {
		const user = await apiClient.getSignedInUser({
			headers: {
				Cookie: request.headers.get("Cookie"),
			},
		});
		return user;
	} catch (e) {
		if (e instanceof AxiosError) {
			if (e.response?.status === 401) {
				const ALLOW_UNAUTHORIZED_PATHS = ["/signin", "/signup"];
				const host = request.headers.get("Host");
				const path = host && request.url.split(host)[1];
				if (path && ALLOW_UNAUTHORIZED_PATHS.includes(path)) return {};
				throw redirect("/signin");
			}
		}
		return {};
	}
}

export function Layout({ children }: { children: React.ReactNode }) {
	const user = useLoaderData<typeof loader>();

	return (
		<html lang="en">
			<head>
				<meta charSet="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<Meta />
				<Links />
			</head>
			<body>
				<Header isSignedIn={!!user?.id} username={user?.username} />
				{children}
				<ScrollRestoration />
				<Scripts />
			</body>
		</html>
	);
}

export function ErrorBoundary() {
	const error = useRouteError();
	console.error(error);
	return (
		<html lang="ja">
			<head>
				<title>Oh no!</title>
				<Meta />
				<Links />
			</head>
			<body>
				<h1
					className={classNames(
						"text-3xl",
						"font-bold",
						"underline",
						"text-red-500",
						"text-center",
						"mt-8",
					)}
				>
					Oh no!
				</h1>
				<Scripts />
			</body>
		</html>
	);
}

export default function App() {
	return <Outlet />;
}
