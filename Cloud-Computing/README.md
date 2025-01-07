# Gymlens Cloud-Computing 🏋️‍♂️☁️

## Backend Service for Gymlens Application 🚀

Welcome to the backend service of **Gymlens**, a platform that combines fitness and technology! Built with **Go Fiber**, containerized using **Docker**, and deployed on **Google Cloud Platform (GCP)**, this service powers the Gymlens app with seamless authentication and data storage using **Firebase** and **Cloud Firestore**.

---

## 🌟 Features
✨ **User Authentication & Management** with Firebase  
✨ RESTful API endpoints for Gymlens functionality  
✨ Scalable deployment with **Google Cloud Platform**  
✨ Secure data storage with **Firestore**  
✨ Dockerized for portability and quick deployment  

---

## 🛠️ Tools & Technologies
- 🌐 **[Google Cloud Platform (GCP)](https://cloud.google.com/)** for cloud services  
- ⚡ **[Go Fiber](https://gofiber.io/)** for a fast and lightweight web framework  
- 🔒 **[Firebase](https://firebase.google.com/)** for authentication and database management  
- 🐳 **[Docker](https://www.docker.com/)** for containerization and deployment  
- 🍨 **[Postman](https://www.postman.com/)** for API documentation and testing  

---

## 📃 Setup Firebase

To interact with Firebase and **Cloud Firestore** from your local environment, you need to configure the **Firebase Admin SDK**.

1. Follow the official guide to set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable:  
   - [Firebase Admin SDK Setup](https://firebase.google.com/docs/admin/setup#initialize_the_sdk_in_non-google_environments)

---

## 🔒 Handling Credentials Safely

The application requires a service account JSON file for Firebase (`GOOGLE_APPLICATION_CREDENTIALS`). This file should **not be included in the repository** for security reasons. Follow these steps to manage it securely:

1. Download the service account JSON file from the Firebase Console.
2. Store it locally in a secure directory.
3. Set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the file's path:
   ```bash
   export GOOGLE_APPLICATION_CREDENTIALS=/path/to/your/credentials.json
   ```
4. Ensure the `.gitignore` file excludes the file:
   ```
   /path/to/your/credentials.json
   ```

By following these steps, you ensure the security of your sensitive data without compromising functionality.

---

## 🚧 Installation Guide

### 1. Clone the repository:
```bash
git clone https://github.com/Gymlens/Cloud-Computing.git
cd Cloud-Computing
```

### 2. Install Dependencies:
```bash
go mod tidy
```

### 3. Set Up Environment Variables:
Create a `.env` file at the root of your project with the following content (replace placeholders with your actual credentials):

```env
PORT=8080
GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json
FIREBASE_API_KEY=your-firebase-api-key
```

### 4. Run the Application:
Use `make` to run the application:
```bash
make run
```

### 5. Verify the Application:
Open a browser and visit [http://localhost:8080/api/ping](http://localhost:8080/api/ping) to check if the server is running properly.

---

## 🎨 Project Structure

Here’s an overview of the directory structure in this project:

```bash
Cloud-Computing
├── api
│   └── user.go                # API-related user routes
├── bin
│   └── Gymlens                # Compiled binary file
├── cmd
│   └── app
│       └── main.go            # Main application entry point
├── config
│   └── config.go             # Configuration setup
├── db
│   └── db.go                 # Database connection logic
├── internal
│   └── server
│       ├── controller
│       │   ├── auth.go       # Authentication controller
│       │   └── user.go       # User-related controller
│       ├── middleware
│       │   ├── auth_middleware.go  # Middleware for auth validation
│       │   └── config.go          # Config-based middleware
│       ├── router
│       │   └── router.go          # Router setup for API
│       └── server.go           # Server setup and initialization
├── models
│   └── user.go               # User model definition
├── pkg
│   └── auth
│       └── auth.go           # Auth logic and helpers
├── .env                      # Environment variables configuration
├── Dockerfile                # Docker configuration for deployment
├── go.mod                    # Go module dependencies
├── go.sum                    # Go sum for module validation
└── Makefile                  # Makefile for automation
```

---

## 📃 API Documentation

The API documentation is available through **Postman**. You can view and test all available endpoints by visiting the following link:
- [Gymlens API Documentation on Postman](https://documenter.getpostman.com/view/40111497/2sAYBYeVPL)

---

## 🚧 Deployment

This application is containerized with **Docker** for easy deployment. To deploy the app in a Docker container:

1. **Build the Docker Image:**
```bash
docker build -t gymlens-cloud-computing .
```

2. **Run the Docker Container:**
```bash
docker run -p 8080:8080 --env-file .env gymlens-cloud-computing
```

3. **Verify the Deployment:**
Visit [http://localhost:8080/api/ping](http://localhost:8080/api/ping) to verify that the deployment was successful.

---

## 📚 Contribution

We welcome contributions! To contribute to the project, please follow these steps:

1. **Fork the repository**.
2. Create a new branch for your feature:
   ```bash
   git checkout -b feature-name
   ```
3. Make your changes and commit:
   ```bash
   git commit -m 'Add new feature'
   ```
4. Push your changes to the forked repository:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request to merge your feature.

---

## 📖 License

This project is licensed under the **MIT License**. See the [LICENSE](https://github.com/rayyandzaki/Gymlens/blob/main/LICENSE) file for more details.

---

## 🌐 Cloud Architecture

Below is the architecture diagram showing how different components of the Gymlens backend interact:

<div align="center">
  <img src="https://raw.githubusercontent.com/Gymlens/.github/main/profile/assets/project-architecture.png" alt="Cloud Architecture" style="width: 100%;">
</div>