Reading a Google Sheet in Go involves interacting with the Google Sheets API. Here's a step-by-step guide to set up and write a Go program for this task. Note that this process requires you to set up Google Cloud credentials and enable the Google Sheets API.

### Step 1: Setting Up Google Cloud Project and Credentials

1. **Create a Google Cloud Project:** Go to the [Google Cloud Console](https://console.cloud.google.com/) and create a new project.

2. **Enable Google Sheets API:** In the Cloud Console, navigate to the project you just created, search for the Google Sheets API, and enable it.

3. **Create Credentials:**
   - In the API & Services > Credentials area, create credentials.
   - Choose "Service account" and follow the steps to create a new service account.
   - Once the service account is created, create and download a JSON key for this account.

4. **Share Your Sheet:**
   - Open the Google Sheet you want to access.
   - Share it with the email address of the service account you created (found in the JSON key file).

### Step 2: Writing the Go Program

First, install the Google Sheets API Go client library:

```sh
go get google.golang.org/api/sheets/v4
go get golang.org/x/oauth2/google
```

Then, write your Go program:

```go
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"

    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

func main() {
    ctx := context.Background()

    b, err := ioutil.ReadFile("path/to/your/service-account-key.json") // Update the path to your JSON key file
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }

    // Configure the Google Sheets API client
    config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := config.Client(ctx)

    sheetService, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // Specify the Spreadsheet ID and Range
    spreadsheetId := "your_spreadsheet_id" // Replace with your Spreadsheet ID
    readRange := "Sheet1!A1:E" // Adjust the range accordingly

    // Read the data
    resp, err := sheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    if len(resp.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        for _, row := range resp.Values {
            // Print columns A and E, which correspond to indices 0 and 4.
            fmt.Printf("%s, %s\n", row[0], row[4])
        }
    }
}
```

Replace `"path/to/your/service-account-key.json"` with the path to your downloaded JSON key, `"your_spreadsheet_id"` with the ID of your Google Sheet, and adjust the `readRange` to match the range of cells you want to read.

### Running the Program

Once you've set everything up and replaced the placeholders with your specific details, you can run your Go program. It will authenticate using the service account and read the specified range from your Google Sheet.

This is a basic example to get you started. Depending on your needs, you might want to expand the functionality, handle different ranges, or process the data in more complex ways. The Google Sheets API offers a lot of flexibility for these tasks.