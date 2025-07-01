# Create Product Microservice

This microservice allows adding new products to the ToyShop platform. It is part of the product management domain and stores product data in MongoDB. Images can be uploaded and are stored in a local folder.

## Technologies Used

- Go (Golang)
- Gin (web framework)
- MongoDB
- MIME/multipart for image uploads
- Docker
- GitHub Actions

## Getting Started

### Prerequisites

- Go >= 1.18
- MongoDB
- Git

### Installation

```bash
git clone https://github.com/andrespaida/create_product.git
cd create_product
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory with the following content:

```env
PORT=4002
MONGO_URI=mongodb://your_mongo_host:27017
DB_NAME=toyshop_db
COLLECTION_NAME=products
```

### Running the Service

```bash
go run main.go
```

The service will be running at `http://localhost:4002`.

## Available Endpoint

### POST `/products`

Creates a new product entry. This endpoint accepts `multipart/form-data` including an image.

#### Form Data Fields:

- `name` (string)
- `description` (string)
- `price` (float)
- `stock` (int)
- `category` (string)
- `image` (file)

#### Example Response:

```json
{
  "message": "Product created successfully",
  "product_id": "60f5cbb2a2e3f0a001d4a7d9"
}
```

Images are stored in the `uploads/` folder and accessible via:

```
http://<host>:4002/uploads/<filename>
```

## Docker

To build and run the service using Docker:

```bash
docker build -t create-product .
docker run -p 4002:4002 --env-file .env create-product
```

## GitHub Actions Deployment

This project includes a GitHub Actions workflow for automatic deployment to an EC2 instance. Configure the following secrets in your GitHub repository:

- `EC2_HOST`
- `EC2_USERNAME`
- `EC2_KEY`
- `EC2_PORT` (optional)

## License

This project is licensed under the MIT License.