# GymLens 🏋️‍♂️

This project is a personal remake of GymLens, a mobile application that categorizes gym equipment and provides detailed usage instructions. The goal is to help beginners better understand gym tools and perform exercises safely, promoting healthier lifestyles and encouraging regular physical activity in Indonesia.

## Key Points 🔗

- **Problem**: Low awareness and understanding of gym equipment usage among Indonesians, contributing to limited physical activity and increasing health risks.
- **Proposed Solution**: A mobile application that categorizes gym equipment and provides detailed descriptions, proper usage guidance, and safety tips.
- **Method**: Application of image classification techniques using deep learning (DenseNet-121 model).
- **Results**: Achieved high classification accuracy, promoting effective and safe workouts through personalized recommendations.

## Raw Datasets 📃

- [Gym Equipments Classification](https://www.kaggle.com/datasets/aadarshvelu/gym-equipements-classification/data)  
- [Gym Equipment Image Dataset](https://www.kaggle.com/datasets/rifqilukmansyah381/gym-equipment-image)

## Process Overview ✅

1. **Developing Endpoints**:
   - Created endpoints for user login and authentication using `auth.go` and `user.go`.
   - Integrated secure middleware for authentication and configuration validation.
   - Can store user information such as name, profile photo, and bio, allowing efficient user data management

2. **Deploying Machine Learning Model**:
   - Packaged the trained DenseNet-121 model (`gym_model_densenet121_9958.h5`) for deployment.
   - Used `app.py` to expose model inference as a REST API.

3. **Dockerizing the Application**:
   - Wrote `Dockerfile` for both the machine learning model and cloud backend services to ensure consistent deployment environments.

4. **Artifact Registry**:
   - Uploaded Docker images to Google Artifact Registry for versioned storage and efficient cloud deployment.

5. **Deploying to Cloud Run**:
   - Configured and deployed services to Google Cloud Run for scalable and serverless execution.
   - Successfully deployed APIs to support mobile application development.

6. **API Integration for Mobile Development**:
   - Exposed APIs for mobile developers to fetch gym equipment classifications and usage guidelines.

7. **Traffic Monitoring**:
   - Set up monitoring tools to track API usage and system performance.
   - Ensured high availability and optimized resource allocation.

## File Structure 🗂

```
GymLens/  
│  
├── Cloud-Computing/                          # Cloud deployment scripts and backend  
│   ├── api/                                 # API-related user routes  
│   ├── bin/                                 # Compiled binary file  
│   ├── cmd/                                 # Application entry point  
│   │   └── app/                              
│   │       └── main.go                      # Main application entry point  
│   ├── config/                              # Configuration setup  
│   ├── db/                                  # Database connection logic  
│   ├── internal/                            # Internal server logic  
│   │   ├── server/                          
│   │   │   ├── controller/                  # Controllers for API  
│   │   │   │   ├── auth.go                  # Authentication controller  
│   │   │   │   └── user.go                  # User-related controller  
│   │   │   ├── middleware/                  
│   │   │   │   ├── auth_middleware.go       # Middleware for auth validation  
│   │   │   │   └── config.go                # Config-based middleware  
│   │   │   ├── router/                      
│   │   │   │   └── router.go                # Router setup for API  
│   │   │   └── server.go                    # Server setup and initialization  
│   ├── models/                              # User model definition  
│   ├── pkg/                                 
│   │   └── auth/                            
│   │       └── auth.go                      # Auth logic and helpers  
│   ├── .env                                 # Environment variables configuration  
│   ├── Dockerfile                           # Docker configuration for deployment  
│   ├── go.mod                               # Go module dependencies  
│   ├── go.sum                               # Go sum for module validation  
│   └── Makefile                             # Makefile for automation  
│  
├── Deploy-ML-Model/                        # Machine learning model deployment  
│   ├── app.py                              # Main application script  
│   ├── cloudbuild.yaml                     # Cloud Build configuration  
│   ├── Dockerfile                          # Docker container configuration  
│   ├── gym_model_densenet121_9958.h5       # Trained model file  
│   ├── requirements.txt                    # Python dependencies  
│   ├── service_account_token_generator.py  # Script to generate service account tokens  
│   ├── .gitignore                          # Git ignore file  
│   └── vscode/                             # VS Code configuration  
└── README.md                               # Project documentation
```

## Further Implementation 🚀

The developed model can be deployed as a cloud-based service to classify gym equipment in real-time. Users can upload equipment images, and the app will provide detailed descriptions and usage guidelines. By integrating the model with a mobile application, users will receive actionable fitness guidance.

## Future Work 🌐

To enhance GymLens, additional features can be integrated, such as:

- **User personalization**: Tailoring recommendations based on fitness goals and experience levels.
- **Augmented reality (AR)**: Providing virtual demonstrations of equipment usage.
- **Multilingual support**: Expanding accessibility to non-English-speaking users.

By addressing these enhancements, GymLens can offer a more comprehensive solution for promoting fitness awareness and adoption in Indonesia.