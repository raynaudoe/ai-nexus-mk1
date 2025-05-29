const { test, expect } = require('@playwright/test');

test('add, delete, and complete tasks', async ({ page }) => {
  await page.goto('http://localhost:3000'); // Replace with your app's URL

  // Add a task
  await page.locator('input[placeholder="Add a task..."]').fill('Buy groceries');
  await page.locator('button:has-text("Add")').click();
  await expect(page.locator('li:has-text("Buy groceries")')).toBeVisible();

  // Add another task
  await page.locator('input[placeholder="Add a task..."]').fill('Walk the dog');
  await page.locator('button:has-text("Add")').click();
  await expect(page.locator('li:has-text("Walk the dog")')).toBeVisible();

  // Complete a task
  await page.locator('li:has-text("Buy groceries") input[type="checkbox"] ').check();
  await expect(page.locator('li:has-text("Buy groceries")')).toHaveClass('completed');

  // Delete a task
  await page.locator('li:has-text("Walk the dog") button:has-text("Delete")').click();
  await expect(page.locator('li:has-text("Walk the dog")')).not.toBeVisible();
});