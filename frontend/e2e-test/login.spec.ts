import { test, expect } from "@playwright/test";

test("ログインができる", async ({ page }) => {
	await page.goto("http://localhost:3000/signin");
	await page.getByLabel("ユーザー名").click();
	await page.getByLabel("ユーザー名").fill("taro");
	await page.getByLabel("ユーザー名").press("Tab");
	await page.getByLabel("パスワード").fill("password");
	await page.getByRole("button", { name: "サインイン" }).click();
	await page.waitForURL("**/");
	await expect(page.getByRole("heading", { name: "投稿一覧" })).toBeVisible();
});
