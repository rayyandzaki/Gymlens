# GymLens Deploy-ML-Model 🚀🤖

Welcome to the **GymLens Deploy-ML-Model** repository! This repo contains the deployment setup for the GymLens machine learning model. It integrates a pre-trained **DenseNet121** model to provide smart predictions via a Python-based API.

---

## 🌟 Features
- 🔄 RESTful API for seamless machine learning inference
- 🤖 Pre-trained **DenseNet121** model for image classification
- 🐳 Containerized using **Docker** for easy deployment
- ☁️ Deployable on **Google Cloud Platform (GCP)** using **Cloud Build** for continuous integration and delivery

---

## 🛠 Tools and Technologies
- **Python** (Flask framework)
- **TensorFlow/Keras** for model inference
- **Google Cloud Platform (GCP)** for cloud services
- **Docker** for containerization
- **Cloud Build** for CI/CD automation

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

## 🚀 Installation
1. **Clone the repository:**
   ```bash
   git clone https://github.com/Gymlens/Deploy-ML-Model.git
   ```

2. **Install dependencies:**
   ```bash
   pip install -r requirements.txt
   ```

3. **Run the application locally:**
   ```bash
   python app.py
   ```

4. **Test the application:**
   Navigate to `http://localhost:5000` to access the API.

---

## 🌿 Environment Variables
The following environment variables are required to run the API:

- `GOOGLE_APPLICATION_CREDENTIALS`: Path to the GCP service account JSON file.
- `MODEL_PATH`: Path to the pre-trained model file (`gym_model_densenet121_9958.h5`).

Add these variables to your environment or a `.env` file:
```env
GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json
MODEL_PATH=gym_model_densenet121_9958.h5
```

---

## 🗂 Project Structure
```
Deploy-ML-Model
├── app.py
├── cloudbuild.yaml
├── Dockerfile
├── gym_model_densenet121_9958.h5
├── requirements.txt
├── service_account_token_generator.py
├── .gitignore
└── vscode
```

---

## 📃 API Documentation

The API documentation is available through **Postman**. You can view and test all available endpoints by visiting the following link:
- [Gymlens API Documentation on Postman](https://documenter.getpostman.com/view/40111497/2sAYBYeVPL)

---

## 📦 Deployment
This application is containerized with Docker and deployable on GCP using Cloud Build.

### Docker Deployment
1. **Build the Docker image:**
   ```bash
   docker build -t gym-lens-ml-api .
   ```

2. **Run the Docker container:**
   ```bash
   docker run -p 5000:5000 --env-file .env gym-lens-ml-api
   ```

3. **Verify the deployment:**
   Visit `http://localhost:5000` to check if the application is running.

### GCP Deployment
1. **Ensure Cloud Build is enabled:**
   - Go to the GCP Console > APIs & Services > Library.
   - Enable the **Cloud Build API**.

2. **Trigger a build:**
   ```bash
   gcloud builds submit --config cloudbuild.yaml
   ```

3. **Deploy to Cloud Run or Compute Engine as needed.**

---

## 📡 API Endpoints
- `POST /predict`: Accepts input data and returns model predictions.

Example usage:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"data": "input_data_here"}' http://localhost:5000/predict
```

---

## 🤝 Contribution
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

---

## 📖 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for more details.

---