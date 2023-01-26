# Password validation

Rest API for password validation.

## Rules

* minSize: at least x number of characters
* minUpperCase: at least x number of upper case characters
* lowerCase: at least x number of lower case characters
* minDigit: at least x number of digits
* minSpecialChars: at least x number of special characters
* noRepeat: no repeated character in sequence


# Run Tests

```bash
go test ./... -cover -v
```

# Run docker

```
docker build -t password-validation .
docker run  -i -t -p 8080:8080 password-validation
```

# Usage

Endpoint url: http://localhost:8080/api/verify
POST method

Request:

```
{
	"password": "PswSord1423@#$@",
	"rules":[
		{"rule": "minSize","value": 2},
		{"rule": "minDigit","value": 4},
        {"rule": "minSpecialChars","value": 4},
		{"rule": "minUpperCase","value": 2},
		{"rule": "minLowerCase","value": 2},
		{"rule": "noRepeted","value": 0}
	]
}
```

Response:

```
{
	"verify": true,
	"noMatch": []
}
```

