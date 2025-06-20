# IITH Canteen Pooling App

A platform for IITH (Indian Institute of Technology Hyderabad) community members to request food deliveries from campus canteens, with incentives for those who fulfill these requests.

## Overview

The IITH Canteen Pooling App is designed to solve the problem of long queues and wait times at campus canteens. It allows users to:

1. Post requests for food items from campus canteens
2. Accept requests from others and deliver their food
3. Earn commission for delivering food to requesters

This creates a win-win situation where busy students/staff can save time, while those who deliver earn extra money.

## Features

- **Request Creation**: Users can create food delivery requests specifying:
  - Canteen location
  - Food items
  - Delivery location
  - Time constraints
  - Commission offered

- **Request Acceptance**: Other users can browse and accept pending requests

- **Delivery Tracking**: Track the status of your request or delivery

- **Commission System**: Automatic calculation and tracking of commissions earned

- **User Profiles**: Build reputation through successful deliveries and fair requests

## How It Works

1. **Requester**:
   - Creates a request for food items
   - Sets a commission amount they're willing to pay
   - Waits for someone to accept their request

2. **Deliverer**:
   - Browses available requests
   - Accepts a request they can fulfill
   - Picks up the food from the canteen
   - Delivers it to the requester
   - Receives the commission

3. **Commission**:
   - Deliverers earn commission for each successful delivery
   - Commission rates are set by requesters based on factors like distance, time, and order size

## Technical Details

- **Backend**: Go (Golang) with Gin web framework
- **API**: RESTful API design
- **Authentication**: Secure user authentication system
- **Database**: [To be implemented]
- **Frontend**: [To be implemented]

## Installation

### Prerequisites
- Go 1.24.2 or higher
- Git

### Setup Instructions
1. Clone the repository
   ```
   git clone https://github.com/yourusername/LambdaCMD.git
   cd LambdaCMD
   ```

2. Install dependencies
   ```
   cd backend
   go mod download
   ```

3. Run the server
   ```
   go run main.go
   ```

4. The server will start on `http://localhost:8080`

## Development Status

This project is currently in early development as part of the Lambda Hackathon. The basic server infrastructure is set up, but the core functionality is still being implemented.

## Future Enhancements

- Mobile application for easier access
- Integration with campus payment systems
- Real-time notifications
- Rating system for users
- Analytics for popular food items and peak times

## Contributing

This project is part of the Lambda Hackathon. Contributions are welcome!

## License

[Include license information here]
