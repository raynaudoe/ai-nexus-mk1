const { test, expect } = require('@playwright/test');

test.describe('Gallery Functionality', () => {
  test('should display images in the gallery', async ({ page }) => {
    await page.goto('/');
    const galleryImages = await page.locator('.gallery img');
    expect(await galleryImages.count()).toBeGreaterThan(0);
  });

  test('should have a responsive layout', async ({ page }) => {
    await page.goto('/');
    await page.setViewportSize({ width: 600, height: 800 });
    const gallery = await page.locator('.gallery');
    expect(await gallery.boundingBox()).toBeDefined();

    await page.setViewportSize({ width: 1200, height: 800 });
    expect(await gallery.boundingBox()).toBeDefined();
  });
});
