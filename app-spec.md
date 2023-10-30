# Holodolor - Abdominal Pain Tracker: Application Specification

## Table of Contents
1. [Objective and Scope](#objective-and-scope)
2. [Features and Functionality](#features-and-functionality)
3. [Data Model](#data-model)

---

## 1. Objective and Scope

### Objective
To create a web application named "Holodolor" for tracking and analyzing various health-related factors potentially contributing to abdominal pain. The application aims to assist in the diagnostic process by offering a comprehensive health log to healthcare providers.

### Scope
- **Target User**: The application is intended for personal use, targeting the user experiencing abdominal pain.
- **Purpose**: Holodolor will serve as a centralized platform for logging events such as food and water intake, medication schedules, episodes of pain, bowel movement details, and body weight.

---

## 2. Features and Functionality

### Core Features

1. **User Authentication**
   - Login
   - Logout
   - Password reset

2. **Dashboard**
   - Timeline view with selectable time frames (day, week, month, year)
   - Icons representing different event types
   - Graphical representation of pain intensity and body weight over time

3. **Event Logging**
   - Log different types of events (food, water, medication, pain, bowel movement, weight)
   - Timestamp for each event

4. **Event Customization**
   - Create custom event types with name, color, and icon
   - Add custom additional properties to event types

5. **Export Data**
   - Export logs to PDF or CSV for sharing with healthcare providers

6. **Search and Filter**
   - Search logs by date, type, or specific terms
   - Filter logs by categories

7. **Notifications and Reminders**
   - Set reminders for medication and daily logging

8. **Settings**
   - User profile management
   - Notification preferences

---

## 3. Data Model

### Tables

1. **User**
   - UserID (Primary Key)
   - Username
   - Email
   - Password (hashed)
   - Created At
   - Updated At

2. **Events**
   - EventID (Primary Key)
   - UserID (Foreign Key)
   - EventTypeID (Foreign Key)
   - Timestamp
   - Created At
   - Updated At

3. **EventTypes**
   - EventTypeID (Primary Key)
   - UserID (Foreign Key)
   - Name
   - Color
   - Icon

4. **AdditionalProperties**
   - AdditionalPropertyID (Primary Key)
   - UserID (Foreign Key)
   - Name

5. **EventTypeAdditionalProperty**
   - UserID (Foreign Key)
   - EventTypeID (Foreign Key)
   - AdditionalPropertyID (Foreign Key)
   - Value

6. **EventAdditionalPropertyValue**
   - ValueID (Primary Key)
   - EventID (Foreign Key)
   - AdditionalPropertyID (Foreign Key)
   - Value

## 4. APIs and Integrations (Refined)

Given the latest requirements, there is no need for external APIs or third-party services. Notifications are not required, and data export will be handled internally, initially supporting only CSV format.

### Internal APIs

1. **Authentication API**
   - Endpoints for user registration, login, and password reset.
   - Uses JWT (JSON Web Tokens) for stateless authentication.

2. **Event Logging API**
   - CRUD (Create, Read, Update, Delete) operations for event logs.
   - Supports batch operations for multiple logs.

3. **Dashboard API**
   - Provides summarized data for the dashboard timeline and graphical representations.
   - Aggregates data by time frame (day, week, month, year).

4. **Event Customization API**
   - CRUD operations for custom event types and additional properties.
  
5. **Data Export API**
   - Generates CSV files of logs based on user-selected criteria.

6. **Search and Filter API**
   - Supports query parameters for searching and filtering logs based on event type, date, and other custom properties.

## 5. User Interface and Experience

### 5.1 Mockups

#### Description
- Mockups serve as a visual guide for the layout and interaction patterns within the application.

#### Action Items
- Create mockups for the following pages:
  1. Landing page
  2. Login/Registration page
  3. Dashboard
  4. Event Logging form
  5. Settings page

### 5.2 User Workflows

#### 5.2.1 Landing Page

- **Not Logged In**: 
  - Display the title "HOLODOLOR."
  - Below the title, provide two buttons: LOGIN and REGISTER.
  
- **Logged In**: 
  - Redirect the user to the Dashboard.

#### 5.2.2 Header Strip

- **Left Side**: Display "Holodolor."
- **Right Side**: Feature a cog wheel icon for Settings and a user icon for user actions like Logout.

#### 5.2.3 Dashboard

- The dashboard is represented as a vertical timeline.
  - **Top of the Screen**: Represents 00:00.
  - **Bottom of the Screen**: Represents 23:00.
  
- Time is depicted by two lines dividing the screen into three equal parts:
  - **Left Side**: For short events like food, drink, and medication.
  - **Right Side**: For pain events.
  
- **Short Events**:
  - Depicted by a circle with a specified color and an icon inside, determined by the event type.
  - Circles appear on the left side of the time line, alternating to the right side if overlapping.
  - Clicking a circle expands it into a rectangle showing event details.
  
- **Pain Events**:
  - Represented as a thick line with rounded edges, spanning from the pain start time to the end time.

#### 5.2.4 Floating Action Button

- Positioned at the bottom-right corner of the screen.
- Marked by a "+" sign.
- Clicking opens an interface to add a new event.

#### 5.2.5 Additional UI Elements

- **Error Messages**: Display clear and informative error messages for form validations.
- **Progress Indicators**: Use visual cues like spinners during tasks requiring loading time.

### Proposed Next Steps

1. **User Interface and Experience**: Create mockups and define user workflows.
2. **Tech Stack**: Choose appropriate technologies for implementation.
3. **Performance and Scalability**: Plan for future growth and performance optimization.
4. **Security Measures**: Outline authentication and data security protocols.
5. **Testing Strategies**: Define test cases for each feature.
6. **Deployment and Maintenance**: Plan deployment and maintenance schedules.
7. **Timeline and Milestones**: Set timelines for each development phase.
