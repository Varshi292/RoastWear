const request = require('supertest');
const express = require('express');
const fs = require('fs');
const path = require('path');
const mediaRouter = require('../Pages/index');

const app = express();
app.use(express.json());
app.use('/', mediaRouter);

describe('Media Upload and Retrieval API', () => {
    test('Should upload an image successfully', async () => {
        const imagePath = 'C:/Users/bggri/Downloads/pngimg.com - doge_meme_PNG16.png';
        
        if (!fs.existsSync(imagePath)) {
            console.error("Test image file does not exist. Please provide a valid test image.");
            return;
        }
        
        const response = await request(app)
            .post('/upload_image') 
            .field('username', 'testUser')
            .attach('image', imagePath);
        
        expect(response.status).toBe(200);
        expect(response.body).toHaveProperty('message', 'Image uploaded and data saved.');
        expect(response.body).toHaveProperty('filepath');
    });

    test('Should return an error for invalid image type', async () => {
        const response = await request(app)
            .post('/upload_image')
            .field('username', 'testUser')
            .attach('image', 'C:/Users/bggri/Downloads/invalid.txt');
        
        expect(response.status).toBe(500);
        expect(response.body).toHaveProperty('error');
    });

    test('Should retrieve uploaded images for a valid username', async () => {
        const response = await request(app)
            .get('/get_user_media?username=testUser');
        
        expect(response.status).toBe(200);
        expect(Array.isArray(response.body.images)).toBe(true);
    });

    test('Should return 404 if no media found for user', async () => {
        const response = await request(app)
            .get('/get_user_media?username=nonExistingUser');
        
        expect(response.status).toBe(404);
        expect(response.body.message).toBe('No images found for this user.');
    });
});