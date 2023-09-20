# goWeather

## Overview

The Weather CLI is a command-line interface (CLI) application that allows you to retrieve real-time weather details for a given city using the weatherapi.com API. With just a city name as input, you can quickly access information such as temperature, weather conditions, and more.

## Prerequisites

Before running the Weather CLI, ensure you have the following:

- Go installed on your system. You can download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)
- An API key from weatherapi.com. You can obtain one by signing up at [https://www.weatherapi.com/](https://www.weatherapi.com/).

## Installation

1. Clone this repository to your local machine:

   bash
   git clone https://github.com/adhilsalim/goWeather.git
   

2. Navigate to the project directory:

   bash
   cd weather-cli
   

3. Open the `goweather.go` file and replace `YOUR API KEY HERE` with your actual weatherapi.com API key.

4. Build the CLI application:

   bash
   go build goweather.go
   

## Usage

To retrieve weather information for a city, run the CLI with the following command:

bash
./goweather <cityname>


Replace `<cityname>` with the name of the city for which you want to fetch weather details.

*Example:*

bash
./goweather kottayam


The CLI will then make a request to the weatherapi.com API and display the weather information for the specified city.

## Contributing

Contributions and improvements to the Weather CLI are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: `git checkout -b feature/new-feature`.
3. Make your changes and commit them: `git commit -m "Add new feature"`.
4. Push your changes to your fork: `git push origin feature/new-feature`.
5. Create a pull request on the main repository.
