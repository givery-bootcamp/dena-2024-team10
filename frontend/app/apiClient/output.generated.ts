import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const Post = z
  .object({
    id: z.number().int(),
    title: z.string(),
    body: z.string(),
    user_id: z.number().int(),
    username: z.string(),
    created_at: z.string().datetime({ offset: true }),
    updated_at: z.string().datetime({ offset: true }),
  })
  .passthrough();
const signupUser_Body = z
  .object({ username: z.string(), password: z.string() })
  .passthrough();
const UserResponse = z
  .object({
    user: z
      .object({ id: z.number().int(), username: z.string() })
      .passthrough(),
  })
  .partial()
  .passthrough();

export const schemas = {
  Post,
  signupUser_Body,
  UserResponse,
};

const endpoints = makeApi([
  {
    method: "get",
    path: "/posts",
    alias: "getPosts",
    description: `Retrieve a list of posts sorted by id in descending order.`,
    requestFormat: "json",
    response: z.array(Post),
  },
  {
    method: "post",
    path: "/signin",
    alias: "signinUser",
    description: `Authenticate a user by providing a username and password.`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        description: `User credentials needed for signing in`,
        type: "Body",
        schema: signupUser_Body,
      },
    ],
    response: UserResponse,
    errors: [
      {
        status: 400,
        description: `Sign-in failed`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
  {
    method: "post",
    path: "/signout",
    alias: "signoutUser",
    description: `sign out a user.`,
    requestFormat: "json",
    response: z.void(),
    errors: [
      {
        status: 400,
        description: `Sign-out failed`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
  {
    method: "post",
    path: "/signup",
    alias: "signupUser",
    description: `Create a new user by providing a username and password.`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        description: `User credentials needed for signing up`,
        type: "Body",
        schema: signupUser_Body,
      },
    ],
    response: UserResponse,
    errors: [
      {
        status: 400,
        description: `Sign-up failed`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
