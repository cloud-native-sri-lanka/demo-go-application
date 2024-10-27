# Demo Go Application

This is a simple Go application designed to demonstrate basic Docker container behavior, including handling crashes and health checks. The app includes three endpoints: a home page, a health check, and a crash trigger.

## Endpoints

- **Home**: `GET /` - Returns a welcome message.
- **Health**: `GET /health` - Returns the status of the application.
- **Crash**: `GET /crash` - Triggers an intentional crash to simulate container failure.
