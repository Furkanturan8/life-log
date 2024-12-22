# Life Log

This project is a backend application built using the GO [Fiber](https://gofiber.io/) framework. It provides APIs for managing users, revenues, expenses, budgets and reports.

## 🚀 Features
- **User Management**: Create, get, and delete users.
- **Income Tracking**: Add, update, get, and delete income records.
- **Expense Tracking**: Add, update, get, and delete expense records.
- **Budget Management**: Create, update, get, and delete budgets.
- **Report Generation**: Generate financial reports.

## 🛠️ Technologies
- **GoLang**: Main programming language.
- **Fiber**: Web framework for API routing.
- **Bun**: ORM library for database interactions.
- **PostgresSQL**: Database used for storage.

## 📚 API Endpoints

### User Routes
- **GET	/api/users**  ->	Get all users
- **POST	/api/user**  ->	Create a new user
- **GET	/api/user/:id**  ->	Get user by ID
- **DELETE	/api/user/:id**  ->	Delete a user by ID

### Income Routes
- **GET	/api/incomes**  ->	Get all incomes
- **POST	/api/income**  ->	Add a new income
- **GET	/api/income/:id**  ->	Get income by ID
- **PUT	/api/income/:id**  ->	Update income by ID
- **DELETE	/api/income/:id**  ->	Delete income by ID

### Expense Routes
- **GET	/api/expenses**  ->	Get all expenses
- **POST	/api/expense**  ->	Add a new expense
- **GET	/api/expense/:id**  ->	Get expense by ID
- **PUT	/api/expense/:id**  ->	Update expense by ID
- **DELETE	/api/expense/:id**  ->	Delete expense by ID

### Budget Routes
- **GET	/api/budgets**	-> Get all budgets
- **POST	/api/budget**	 -> Create a new budget
- **GET	/api/budget/:id**  ->	Retrieve budget by ID
- **PUT	/api/budget/:id**	-> Update budget by ID
- **DELETE	/api/budget/:id**  ->	Delete budget by ID

### Report Routes
**POST	/api/report**	 -> Generate a financial report

[View Report PDF](https://github.com/Furkanturan8/life-log/blob/main/uploads/report-11-2024-NX.pdf)
