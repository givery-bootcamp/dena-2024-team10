import { type ActionFunctionArgs, json, type MetaFunction } from "@remix-run/node";
import { Form, useLoaderData, useNavigate, useNavigation } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  const data = useLoaderData<typeof loader>()
  const submitting = useNavigation()

  return (
    <h1 className="text-3xl font-bold underline">
      {data.count}
      <Form method="post">
        <input type="number" name="count"/>
        <input type="submit" disabled={submitting.state !== "idle"} className="px-8 py-4 bg-red-200 rounded-full hover:bg-red-500 hover:text-white disabled:opacity-20 transition-colors" />
      </Form>
    </h1> 
  ); 
}



// server side
let counter = 0

const delay = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))

export function loader() {
  return json({count: counter})
}

export async function action(args: ActionFunctionArgs) {
  const formData = await args.request.formData()
  counter += Number.parseInt(formData.get("count") as string)
  await delay(1000) 
  return null
}