// Shared Playwright helpers and utilities for all platforms

export async function waitForSelectorWithTimeout(page, selector, timeout = 10000) {
  await page.waitForSelector(selector, { timeout });
}
