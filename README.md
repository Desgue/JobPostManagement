# Job Post Management Backend Challenge
Searching Github repos for backend challenges I found this [repo](https://github.com/Michee27/backend-developer-test) and found the challenge quite interesting to tackle and pratice AWS integration in my APIs. In a way of keep praticing and developing my skills, I decided to tackle this in my own time, and this is my version presented in Go.



## The Challenge
The following is the challenge description given at the original repo with minor modification made by me.

1. Relational Database: Utilize a SQL database (PostgreSQL 16) with two tables (companies and jobs). The DDL script in the ddl folder of this repository initializes these tables. The companies table is pre-populated with fictitious records, which you should not modify. Focus on managing records in the jobs table. You don't need to worry about setting up the database, consider the database is already running in the cloud. Your code only needs to handle database connections. To test your solution, use your own database running locally or in the server of your choice.

2. Serverless Environment: Implement asynchronous, event-driven logic using AWS Lambda and AWS SQS for queue management.

3. Job Feed Repository: Integrate a job feed with AWS S3. This feed should periodically update a JSON file reflecting the latest job postings.

### API Endpoints
- GET /companies: List existing companies.

- GET /companies/:company_id: Fetch a specific company by ID.

- GET /feed: Serve a job feed in JSON format, containing published jobs (column status = 'published').

- POST /job: Create a job posting draft.

- PUT /job/:job_id: Edit a job posting draft (title, location, description).

- PUT /job/:job_id/publish: Publish a job posting draft.

- PUT /job/:job_id/archive: Archive an active job posting.

- DELETE /job/:job_id: Delete a job posting draft.

### Integration Features

- Use a caching mechanism to handle high traffic, fetching data from an S3 file updated periodically by an AWS Lambda function. The feed should return the job ID, title, description, company name and the date when the job was created. This endpoint should not query the database, the content must be fetched from S3.

- This endpoint receives a massive number of requests every minute, so the strategy here is to implement a simple cache mechanism that will fetch a previously stored JSON file containing the published jobs and serve the content in the API. You need to implement a serverless component using AWS Lambda, that will periodically query the published jobs and store the content on S3. The GET /feed endpoint should fetch the S3 file and serve the content. You don't need to worry about implementing the schedule, assume it is already created using AWS EventBridge. You only need to create the Lambda component.

## Extra Features (Optional)

- Job Moderation: using artificial intelligence, we need to moderate the job content before allowing it to be published, to check for potential harmful content. Every time a user requests a job publication (PUT /job/:job_id/publish), the API should reply with success to the user, but the job should not be immediately published.
It should be queued using AWS SQS, feeding the job to a Lambda component. Using OpenAI's free moderation API, create a Lambda component that will evaluate the job title and description, and test for hamrful content. If the content passes the evaluation, the component should change the job status to published, otherwise change to rejected and add the response from OpenAI API to the notes column.
