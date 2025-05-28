# NOT FINISHED YET, PLEASE DO NOT USE UNTIL RELEASED

# GolarEdge

Hello **Go**pher, GolarEdge is a comprehensive Go API client library for the SolarEdge Monitoring API. GolarEdge provides easy access to all available endpoints, allowing developers to seamlessly integrate SolarEdge monitoring data into their Go applications.

## Features
 
- **Complete Endpoint Coverage:** All endpoints of the SolarEdge Monitoring API are currently implemented and accessible.
- **Idiomatic Go:** Designed with Go best practices in mind, offering a clean and intuitive API.
- **Easy to Use:** Simple setup and straightforward function calls for quick data retrieval.

## Installation
To use GolarEdge in your Go project, run the following command:
go get github.com/Adrigorithm/GolarEdge

## Usage
```go
// Soon
```

### API Key
It is recommended to store your SolarEdge API key as an environment variable (e.g., SOLAREDGE_API_KEY) and retrieve it in your application. Never expose your token in plain text anywhere except for testing in development (and even then rather not).

## Documentation
For detailed information on all available methods and data structures, please refer to the GoDoc Reference.

## Future Plans
I have exciting plans to enhance GolarEdge further:
 
- [] **Rate Limiting:** Implement robust rate limiting to ensure compliance with SolarEdge API usage policies and prevent exceeding request limits.,
- [] **Caching:** Introduce intelligent caching mechanisms to reduce redundant API calls and improve performance.,
- [] **Modbus Support:** Add support for Modbus integration with SolarEdge devices, enabling direct communication and data retrieval from inverters and other equipment.,

## Contributing
We welcome contributions to GolarEdge! If you have suggestions, bug reports, or want to contribute code, please feel free to open an issue or submit a pull request.

## License
This library is distributed under the MIT License. See the LICENSE file for more information.

## Acknowledgements
- Google (and external contributors) for creating Go and its packages
- SolarEdge Monitoring API Documentation
