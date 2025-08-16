# HitList Web App

## Overview
A full-stack web app that integrates with the **Spotify Web API**.  
Users can search for songs, submit daily entries, and view a randomly selected **Hit of the Day**.  
Authenticated users can stream the featured track directly on their active Spotify device.

## Tech Stack
- **Go** (os, time, net/http, encoding/json, html/template)
- **PostgreSQL**
- HTML/CSS, JavaScript

## Features
- Search for songs using Spotify's Web API  
- Submit daily entries (resets every 24 hours)  
- Randomly select and display the "Hit of the Day"  
- Playback directly on a logged-in Spotify device  

## Setup
1. Clone the repo  
2. Install dependencies  
3. Run the app:  
   ```bash
   go run main.go
