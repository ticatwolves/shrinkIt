# ShrinkIt — Serverless URL Shortener  

ShrinkIt is an **open-source, serverless URL shortener** built with **AWS SAM** and **Go (Golang)**. It allows you to shorten long URLs into tiny, shareable links while staying cost-efficient and fully under your control.  

## ✨ Features  
- 🔗 **Short URLs** — Generate simple, easy-to-share links.  
- ⚡ **Serverless** — Powered by AWS Lambda for high scalability and low cost.  
- 🛠 **Built with Go** — Fast, efficient, and highly portable.  
- 🧩 **Infrastructure as Code** — Managed using AWS SAM templates.  

## 🏗️ Architecture  
ShrinkIt is designed as a **fully serverless app**:  
- **API Gateway** — Provides REST endpoints for creating and retrieving short URLs.  
- **Lambda (Go)** — Business logic written in Go, handling shortening and redirection.  
- **DynamoDB** — Stores original URLs and short-code mappings.  
- **AWS SAM** — Deploys and manages the entire infrastructure.  

```
[ Client ] → [ API Gateway ] → [ Lambda (Go) ] → [ DynamoDB ]
```

## 🚀 Getting Started  

### Prerequisites  
- [AWS CLI](https://docs.aws.amazon.com/cli/) configured with your account  
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/) installed  
- [Go 1.23+](https://go.dev/dl/)  

### 1. Clone the Repository  
```bash
git clone https://github.com/ticatwolves/shrinkIt.git
cd shrinkIt
```

### 2. Build the Project  
```bash
make build
```

### 3. Deploy to AWS  
```bash
make deploy
```
Follow the prompts to configure stack name, region, and permissions.  

### 4. Usage  
- **Shorten a URL**  
  ```bash
  curl -X POST https://<api-id>.execute-api.<region>.amazonaws.com/prod     -H "Content-Type: application/json"     -d '{"url": "https://example.com/very/long/url"}'
  ```
  Response:
  ```json
  { "short_url": "https://<api-id>.execute-api.<region>.amazonaws.com/prod/shrink/abcd123" }
  ```

- **Redirect to original URL**  
  ```bash
  curl -i https://<api-id>.execute-api.<region>.amazonaws.com/prod/shrink/abcd123
  ```

## 🛠 Development  

### Run Locally with SAM  
```bash
make local-start
```
Test with:  
```bash
curl http://127.0.0.1:3000/shrink -d '{"url":"https://example.com"}'
```

## 📂 Project Structure  
```
.
├── cmd/             # Go Lambda handler
├── internal/        # Business logic
├── deploy/template.yaml    # AWS SAM template
├── go.mod
└── README.md
```

## 📜 License  
This project is licensed under the **MIT License**.  
