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
const createPost_Body = z
  .object({ title: z.string().max(100), body: z.string() })
  .passthrough();
const updatePost_Body = z
  .object({ title: z.string().max(100), body: z.string() })
  .partial()
  .passthrough();
const Comment = z
  .object({
    id: z.number().int(),
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
  .object({ id: z.number().int(), username: z.string() })
  .passthrough();

export const schemas = {
  Post,
  createPost_Body,
  updatePost_Body,
  Comment,
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
    parameters: [
      {
        name: "limit",
        type: "Query",
        schema: z.number().int().gte(1).lte(100).optional().default(20),
      },
      {
        name: "offset",
        type: "Query",
        schema: z.number().int().gte(0).optional().default(0),
      },
    ],
    response: z.array(Post),
    errors: [
      {
        status: 401,
        description: `Unauthorized`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
  {
    method: "post",
    path: "/posts",
    alias: "createPost",
    description: `Create a new post by providing a title and body.`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        description: `Post details needed for creation`,
        type: "Body",
        schema: createPost_Body,
      },
    ],
    response: Post,
  },
  {
    method: "get",
    path: "/posts/:postId",
    alias: "getPost",
    description: `Retrieve a post by its ID.`,
    requestFormat: "json",
    parameters: [
      {
        name: "postId",
        type: "Path",
        schema: z.number().int(),
      },
    ],
    response: Post,
  },
  {
    method: "put",
    path: "/posts/:postId",
    alias: "updatePost",
    description: `Update an existing post by providing a title and body.`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        description: `Post details needed for update`,
        type: "Body",
        schema: updatePost_Body,
      },
      {
        name: "postId",
        type: "Path",
        schema: z.number().int(),
      },
    ],
    response: Post,
    errors: [
      {
        status: 400,
        description: `You are not the author of the post`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
  {
    method: "delete",
    path: "/posts/:postId",
    alias: "deletePost",
    description: `Delete an existing post by its ID (logical deletion by setting &#x60;deleted_at&#x60;).`,
    requestFormat: "json",
    parameters: [
      {
        name: "postId",
        type: "Path",
        schema: z.number().int(),
      },
    ],
    response: z.void(),
    errors: [
      {
        status: 400,
        description: `You are not the author of the post`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
  {
    method: "get",
    path: "/posts/:postId/comments",
    alias: "getComments",
    description: `Retrieve a list of comments for a post sorted by id in descending order.`,
    requestFormat: "json",
    parameters: [
      {
        name: "postId",
        type: "Path",
        schema: z.number().int(),
      },
      {
        name: "limit",
        type: "Query",
        schema: z.number().int().gte(1).lte(100).optional().default(20),
      },
      {
        name: "offset",
        type: "Query",
        schema: z.number().int().gte(0).optional().default(0),
      },
    ],
    response: z.array(Comment),
  },
  {
    method: "post",
    path: "/posts/:postId/comments",
    alias: "createComment",
    description: `Create a new comment for a post by providing a body.`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        description: `Comment details needed for creation`,
        type: "Body",
        schema: z.object({ body: z.string() }).passthrough(),
      },
    ],
    response: Comment,
  },
  {
    method: "put",
    path: "/posts/:postId/comments/:commentId",
    alias: "updateComment",
    description: `Update an existing comment by providing a body.`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        description: `Comment details needed for update`,
        type: "Body",
        schema: z.object({ body: z.string() }).partial().passthrough(),
      },
      {
        name: "postId",
        type: "Path",
        schema: z.number().int(),
      },
      {
        name: "commentId",
        type: "Path",
        schema: z.number().int(),
      },
    ],
    response: Comment,
    errors: [
      {
        status: 400,
        description: `You are not the author of the comment`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
  {
    method: "delete",
    path: "/posts/:postId/comments/:commentId",
    alias: "deleteComment",
    description: `Delete an existing comment by its ID (logical deletion by setting &#x60;deleted_at&#x60;).`,
    requestFormat: "json",
    parameters: [
      {
        name: "postId",
        type: "Path",
        schema: z.number().int(),
      },
      {
        name: "commentId",
        type: "Path",
        schema: z.number().int(),
      },
    ],
    response: z.void(),
    errors: [
      {
        status: 400,
        description: `You are not the author of the comment`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
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
        status: 401,
        description: `Unauthorized`,
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
  {
    method: "get",
    path: "/user",
    alias: "getSignedInUser",
    description: `get signed-in user.`,
    requestFormat: "json",
    response: UserResponse,
    errors: [
      {
        status: 401,
        description: `Unauthorized`,
        schema: z.object({ message: z.string() }).partial().passthrough(),
      },
    ],
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
