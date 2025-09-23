package api

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func ServeHTML() string {
	htmlContent := `<!DOCTYPE html>
        <html lang='en'>
        <head>
            <meta charset='UTF-8'>
            <meta name='viewport' content='width=device-width, initial-scale=1.0'>
            <title>ShrinkIt</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                }
                .form-container {
                    text-align: center;
                    padding: 20px;
                    border: 1px solid #ccc;
                    border-radius: 5px;
                    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
                }
                input[type='text'] {
                    padding: 8px;
                    margin: 10px 0;
                    border-radius: 4px;
                    border: 1px solid #ddd;
                }
                button {
                    padding: 8px 16px;
                    border-radius: 4px;
                    background-color: #4CAF50;
                    color: white;
                    border: none;
                    cursor: pointer;
                }
                button:hover {
                    background-color: #45a049;
                }
            </style>
        </head>
        <body>
            <div class='form-container'>
                <h2>ShrinkIt</h2>
                <form>
                    <input type='text' id='inputField' placeholder='Enter URL...' required>
                    <br>
                    <button type='button' onclick='handleSubmit()'>Shrink</button>
                </form>
            </div>

            <script>
                function handleSubmit() {
                    const inputValue = document.getElementById('inputField').value;
                    if(inputValue) {
                        console.log(inputValue);
                    } else {
                        console.log("Incorrect input");
                    }
                }
            </script>
        </body>
        </html>`
	return htmlContent
}

func ShrinkItUIRequestHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Request received:", request)

	// Custom business logic here

	// Simple response
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       ServeHTML(),
		Headers:    map[string]string{"Content-Type": "text/html"},
	}, nil
}
