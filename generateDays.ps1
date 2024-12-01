# Base directory where the "Days" folder will be created
$baseDir = "Days"

# Create the base directory if it doesn't exist
if (!(Test-Path -Path $baseDir)) {
    New-Item -Path $baseDir -ItemType Directory
}

# Loop from Day2 to Day25
for ($i = 2; $i -le 25; $i++) {
    # Create the folder for the current day
    $dayDir = "$baseDir\Day$i"
    if (!(Test-Path -Path $dayDir)) {
        New-Item -Path $dayDir -ItemType Directory
    }

    # Create the Go file for the current day
    $filePath = "$dayDir\day$i.go"
    $content = @"
package day$i

import "fmt"

func Run() {
    fmt.Println("Running Day $i!")
}
"@
    Set-Content -Path $filePath -Value $content
}

Write-Host "Folders and Go files for Day2 to Day25 have been created."
